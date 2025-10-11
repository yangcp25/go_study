package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"net"
	"sync"
)
import (
	"bufio"
	"encoding/binary"
)

const (
	chunkSize      = 32 * 1024 // 每个 chunk 的大小（32KB）
	perKeyChanBuf  = 64        // 每个 key 的 channel 缓冲
	incomingBuf    = 32        // incoming channel 缓冲
	maxFrameLength = 1 << 30   // 单帧最大长度保护（1GB）
)

// chunk 是从 pool 拿到的缓冲片段
type chunk struct {
	b []byte
	n int // 有效字节数
}

type perKey struct {
	pr   *io.PipeReader
	pw   *io.PipeWriter
	ch   chan chunk
	done chan struct{}
}

type incomingStream struct {
	key    string
	reader io.Reader
}

type Conn struct {
	conn   net.Conn
	reader *bufio.Reader

	incoming chan incomingStream

	mu     sync.Mutex
	stream map[string]*perKey

	sendMu sync.Mutex

	pool sync.Pool

	closeOnce sync.Once
	closed    chan struct{}
}

// NewConn 包装 net.Conn 并启动 readLoop
func NewConn(c net.Conn) *Conn {
	cc := &Conn{
		conn:     c,
		reader:   bufio.NewReader(c),
		incoming: make(chan incomingStream, incomingBuf),
		stream:   make(map[string]*perKey),
		closed:   make(chan struct{}),
	}
	cc.pool.New = func() interface{} {
		buf := make([]byte, chunkSize)
		return buf
	}
	go cc.readLoop()
	return cc
}

// Close 关闭连接（幂等）
func (c *Conn) Close() {
	c.closeOnce.Do(func() {
		close(c.closed)
		c.conn.Close()
		// readLoop 会做后续清理并关闭 incoming 通道
	})
}

// Send 返回 writer；调用方必须在完成该 key 的写入后调用 writer.Close() 来发送 End Frame 并释放 sendMu
func (c *Conn) Send(key string) (io.WriteCloser, error) {
	// 串行化逻辑流写：持有 sendMu 直到 Close() 释放
	c.sendMu.Lock()
	// 先写 Key Frame
	if err := writeLenPrefixed(c.conn, []byte(key)); err != nil {
		c.sendMu.Unlock()
		return nil, err
	}
	sw := &streamWriter{c: c, once: sync.Once{}}
	return sw, nil
}

type streamWriter struct {
	c      *Conn
	closed bool
	once   sync.Once
}

// Write 将 p 拆成若干 chunk，每个 chunk 写成一个 data frame [len][data]
func (w *streamWriter) Write(p []byte) (int, error) {
	if w.closed {
		return 0, io.ErrClosedPipe
	}
	total := 0
	for len(p) > 0 {
		to := len(p)
		if to > chunkSize {
			to = chunkSize
		}
		if err := writeLenPrefixed(w.c.conn, p[:to]); err != nil {
			return total, err
		}
		total += to
		p = p[to:]
	}
	return total, nil
}

// Close 发送 End Frame(len==0) 并释放 send 锁
func (w *streamWriter) Close() error {
	var err error
	w.once.Do(func() {
		err = writeLenPrefixed(w.c.conn, nil)
		w.closed = true
		w.c.sendMu.Unlock()
	})
	return err
}

// Receive 返回下一个到达的 key 与对应的 reader；连接关闭且没有更多流时返回 io.EOF
func (c *Conn) Receive() (string, io.Reader, error) {
	select {
	case s, ok := <-c.incoming:
		if !ok {
			return "", nil, io.EOF
		}
		return s.key, s.reader, nil
	case <-c.closed:
		// 尝试再取一次以防 race
		select {
		case s, ok := <-c.incoming:
			if !ok {
				return "", nil, io.EOF
			}
			return s.key, s.reader, nil
		default:
			return "", nil, io.EOF
		}
	}
}

// ---------------- readLoop 与辅助 ----------------

func (c *Conn) readLoop() {
	// 退出时清理：把每个活跃 stream 的 pw 以错误或 EOF 关闭，并关闭 incoming
	defer func() {
		c.mu.Lock()
		for _, pk := range c.stream {
			// 确保关闭 pipe writer（传递 EOF）
			pk.pw.CloseWithError(io.EOF)
			close(pk.ch)
		}
		c.mu.Unlock()
		close(c.incoming)
		c.conn.Close()
	}()

	for {
		// 1) 读 Key Frame (len + key)
		keyBytes, err := readLenPrefixed(c.reader)
		if err != nil {
			c.propagateErrorToStreams(err)
			return
		}
		key := string(keyBytes)

		// 2) 确保 perKey 存在并且 dispenser 正在运行
		pk := c.getOrCreatePerKey(key)

		// 3) 发布 incoming（如果缓冲满则异步发送，避免阻塞 readLoop）
		select {
		case c.incoming <- incomingStream{key: key, reader: pk.pr}:
		default:
			// 异步发送，避免阻塞解析 loop；若连接关闭则放弃
			go func(k string, r io.Reader) {
				select {
				case c.incoming <- incomingStream{key: k, reader: r}:
				case <-c.closed:
				}
			}(key, pk.pr)
		}

		// 4) 连续读 Data Frames，直到 length==0 表示 End Frame
		for {
			var length uint32
			if err := binary.Read(c.reader, binary.BigEndian, &length); err != nil {
				c.propagateErrorToStreams(err)
				return
			}
			if length == 0 {
				// 该 key 的传输结束，关闭对应 channel，让 dispenser 关闭 pw
				c.mu.Lock()
				cur, ok := c.stream[key]
				if ok {
					close(cur.ch)
				}
				c.mu.Unlock()
				break
			}
			if length > maxFrameLength {
				c.propagateErrorToStreams(fmt.Errorf("frame too large: %d", length))
				return
			}

			// 把 length 字节分块读取并发送到 per-key channel
			remaining := int(length)
			for remaining > 0 {
				toRead := remaining
				if toRead > chunkSize {
					toRead = chunkSize
				}
				buf := c.pool.Get().([]byte) // 从 pool 取缓冲
				// 读取确切 toRead 字节到 buf[:toRead]
				if _, err := io.ReadFull(c.reader, buf[:toRead]); err != nil {
					// 读失败，归还 buf 并传播错误
					c.pool.Put(buf)
					c.propagateErrorToStreams(err)
					return
				}
				// 发送 chunk（如果 per-key ch 满了，这里会阻塞，产生回压，只影响该 key）
				select {
				case pk.ch <- chunk{b: buf, n: toRead}:
				case <-c.closed:
					c.pool.Put(buf)
					return
				}
				remaining -= toRead
			}
		}
	}
}

// propagateErrorToStreams 会把 err 传播给所有活跃 stream（通过 CloseWithError），并关闭它们的 ch
func (c *Conn) propagateErrorToStreams(err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, pk := range c.stream {
		if pk.pw != nil {
			pk.pw.CloseWithError(err)
		}
		// 关闭 ch（若未关闭）
		select {
		case <-pk.done:
			// already done
		default:
			// safe close if not already closed
			select {
			case <-pk.done:
			default:
				close(pk.ch)
			}
		}
	}
}

// getOrCreatePerKey：若不存在就创建 perKey（包含 pr/pw、channel），并启动 dispenser goroutine
func (c *Conn) getOrCreatePerKey(key string) *perKey {
	c.mu.Lock()
	pk, ok := c.stream[key]
	if ok {
		c.mu.Unlock()
		return pk
	}
	// create pipe and channel
	pr, pw := io.Pipe()
	ch := make(chan chunk, perKeyChanBuf)
	pk = &perKey{
		pr:   pr,
		pw:   pw,
		ch:   ch,
		done: make(chan struct{}),
	}
	c.stream[key] = pk
	c.mu.Unlock()

	// 开启 dispenser：把 ch 中的 chunk 顺序写入 pw（写可能阻塞，仅影响该 key）
	go func(k string, pk *perKey) {
		defer func() {
			// 退出时关闭 pipe writer（若未已关闭），移除 map 条目
			pk.pw.Close()
			c.mu.Lock()
			delete(c.stream, k)
			c.mu.Unlock()
			close(pk.done)
		}()

		for chnk := range pk.ch {
			// 写入 pipe writer（可能阻塞）
			_, err := pk.pw.Write(chnk.b[:chnk.n])
			// 立刻把 buffer 放回 pool（即使写阻塞也会放回）
			c.pool.Put(chnk.b)
			if err != nil {
				// 出错则用 CloseWithError 结束
				pk.pw.CloseWithError(err)
				return
			}
		}
		// channel 关闭 -> 正常结束，defer 会关闭 pw
	}(key, pk)

	return pk
}

// writeLenPrefixed: 写 uint32(len) + payload（len==0 表示 end）
func writeLenPrefixed(w io.Writer, data []byte) error {
	var hb [4]byte
	binary.BigEndian.PutUint32(hb[:], uint32(len(data)))
	if _, err := w.Write(hb[:]); err != nil {
		return err
	}
	if len(data) > 0 {
		_, err := w.Write(data)
		return err
	}
	return nil
}

// readLenPrefixed: 读 uint32(len) + payload
func readLenPrefixed(r io.Reader) ([]byte, error) {
	var hb [4]byte
	if _, err := io.ReadFull(r, hb[:]); err != nil {
		return nil, err
	}
	length := binary.BigEndian.Uint32(hb[:])
	if length == 0 {
		return []byte{}, nil
	}
	if length > maxFrameLength {
		return nil, fmt.Errorf("length too large: %d", length)
	}
	buf := make([]byte, length)
	if _, err := io.ReadFull(r, buf); err != nil {
		return nil, err
	}
	return buf, nil
}

//////////////////////////////////////////////
///////// 接下来的代码为测试代码，请勿修改 /////////
//////////////////////////////////////////////

// 连接到测试服务器，获得一个你实现的连接对象
func dial(serverAddr string) *Conn {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		panic(err)
	}
	return NewConn(conn)
}

// 启动测试服务器
func startServer(handle func(*Conn)) net.Listener {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println("[WARNING] ln.Accept", err)
				return
			}
			go handle(NewConn(conn))
		}
	}()
	return ln
}

// 简单断言
func assertEqual[T comparable](actual T, expected T) {
	if actual != expected {
		panic(fmt.Sprintf("actual:%v expected:%v\n", actual, expected))
	}
}

// 简单 case：单连接，双向传输少量数据
func testCase0() {
	const (
		key  = "Bible"
		data = `Then I heard the voice of the Lord saying, “Whom shall I send? And who will go for us?”
And I said, “Here am I. Send me!”
Isaiah 6:8`
	)
	ln := startServer(func(conn *Conn) {
		// 服务端等待客户端进行传输
		_key, reader, err := conn.Receive()
		if err != nil {
			panic(err)
		}
		assertEqual(_key, key)
		dataB, err := io.ReadAll(reader)
		if err != nil {
			panic(err)
		}
		assertEqual(string(dataB), data)

		// 服务端向客户端进行传输
		writer, err := conn.Send(key)
		if err != nil {
			panic(err)
		}
		n, err := writer.Write([]byte(data))
		if err != nil {
			panic(err)
		}
		if n != len(data) {
			panic(n)
		}
		conn.Close()
	})
	//goland:noinspection GoUnhandledErrorResult
	defer ln.Close()

	conn := dial(ln.Addr().String())
	// 客户端向服务端传输
	writer, err := conn.Send(key)
	if err != nil {
		panic(err)
	}
	n, err := writer.Write([]byte(data))
	if n != len(data) {
		panic(n)
	}
	err = writer.Close()
	if err != nil {
		panic(err)
	}
	// 客户端等待服务端传输
	_key, reader, err := conn.Receive()
	if err != nil {
		panic(err)
	}
	assertEqual(_key, key)
	dataB, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	assertEqual(string(dataB), data)
	conn.Close()
}

// 生成一个随机 key
func newRandomKey() string {
	buf := make([]byte, 8)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(buf)
}

// 读取随机数据，并返回随机数据的校验和：用于验证数据是否完整传输
func readRandomData(reader io.Reader, hash hash.Hash) (checksum string) {
	hash.Reset()
	var buf = make([]byte, 23<<20) //调用者读取时的 buf 大小不是固定的，你的实现中不可假定 buf 为固定值
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		_, err = hash.Write(buf[:n])
		if err != nil {
			panic(err)
		}
	}
	checksum = hex.EncodeToString(hash.Sum(nil))
	return checksum
}

// 写入随机数据，并返回随机数据的校验和：用于验证数据是否完整传输
func writeRandomData(writer io.Writer, hash hash.Hash) (checksum string) {
	hash.Reset()
	const (
		dataSize = 500 << 20 //一个 key 对应 500MB 随机二进制数据，dataSize 也可以是其他值，你的实现中不可假定 dataSize 为固定值
		bufSize  = 1 << 20   //调用者写入时的 buf 大小不是固定的，你的实现中不可假定 buf 为固定值
	)
	var (
		buf  = make([]byte, bufSize)
		size = 0
	)
	for i := 0; i < dataSize/bufSize; i++ {
		_, err := rand.Read(buf)
		if err != nil {
			panic(err)
		}
		_, err = hash.Write(buf)
		if err != nil {
			panic(err)
		}
		n, err := writer.Write(buf)
		if err != nil {
			panic(err)
		}
		size += n
	}
	if size != dataSize {
		panic(size)
	}
	checksum = hex.EncodeToString(hash.Sum(nil))
	return checksum
}

// 复杂 case：多连接，双向传输，大量数据，多个不同的 key
func testCase1() {
	var (
		mapKeyToChecksum = map[string]string{}
		lock             sync.Mutex
	)
	ln := startServer(func(conn *Conn) {
		// 服务端等待客户端进行传输
		key, reader, err := conn.Receive()
		if err != nil {
			panic(err)
		}
		var (
			h         = sha256.New()
			_checksum = readRandomData(reader, h)
		)
		lock.Lock()
		checksum, keyExist := mapKeyToChecksum[key]
		lock.Unlock()
		if !keyExist {
			panic(fmt.Sprintln(key, "not exist"))
		}
		assertEqual(_checksum, checksum)

		// 服务端向客户端连续进行 2 次传输
		for _, key := range []string{newRandomKey(), newRandomKey()} {
			writer, err := conn.Send(key)
			if err != nil {
				panic(err)
			}
			checksum := writeRandomData(writer, h)
			lock.Lock()
			mapKeyToChecksum[key] = checksum
			lock.Unlock()
			err = writer.Close() //表明该 key 的所有数据已传输完毕
			if err != nil {
				panic(err)
			}
		}
		conn.Close()
	})
	//goland:noinspection GoUnhandledErrorResult
	defer ln.Close()

	conn := dial(ln.Addr().String())
	// 客户端向服务端传输
	var (
		key = newRandomKey()
		h   = sha256.New()
	)
	writer, err := conn.Send(key)
	if err != nil {
		panic(err)
	}
	checksum := writeRandomData(writer, h)
	lock.Lock()
	mapKeyToChecksum[key] = checksum
	lock.Unlock()
	err = writer.Close()
	if err != nil {
		panic(err)
	}

	// 客户端等待服务端的多次传输
	keyCount := 0
	for {
		key, reader, err := conn.Receive()
		if err == io.EOF {
			// 服务端所有的数据均传输完毕，关闭连接
			break
		}
		if err != nil {
			panic(err)
		}
		_checksum := readRandomData(reader, h)
		lock.Lock()
		checksum, keyExist := mapKeyToChecksum[key]
		lock.Unlock()
		if !keyExist {
			panic(fmt.Sprintln(key, "not exist"))
		}
		assertEqual(_checksum, checksum)
		keyCount++
	}
	assertEqual(keyCount, 2)
	conn.Close()
}

func main() {
	testCase0()
	testCase1()
}
