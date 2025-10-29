package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"net"
	"sync"
)

// Conn 是你需要实现的一种连接类型，它支持下面描述的若干接口；
// 为了实现这些接口，你需要设计一个基于 TCP 的简单协议；
//
// 协议设计：
// 采用 Length-Value (LV) 格式进行分帧。
// 一个完整的 key-value 数据流由以下帧序列组成：
// 1. Key Frame: [4-byte length of key][key content]
// 2. Data Frames: [4-byte length of data chunk][data chunk content] ... (可以有多个)
// 3. End Frame: [4-byte length = 0]
// 所有长度均为 BigEndian 编码的 uint32。
// 定义协议中的帧类型
const (
	FrameTypeData byte = 0x01
	FrameTypeEnd  byte = 0x02
)

// incomingStream 用于在 readLoop 和 Receive 之间传递新发现的数据流
type incomingStream struct {
	key    string
	reader io.Reader
}

type Conn struct {
	netConn net.Conn

	// receive 相关
	incoming chan incomingStream       // 用于接收新的数据流
	readers  map[string]*io.PipeWriter // 存储每个 key 对应的 pipe writer

	// send 相关
	writeMutex sync.Mutex // 保证并发写 net.Conn 是安全的
}

// readLoop 是 Conn 的核心，负责读取和解析所有来自对端的数据
func (c *Conn) readLoop() {
	defer func() {
		// readLoop 退出时，关闭所有正在进行的 reader，并关闭 incoming channel
		for _, writer := range c.readers {
			writer.Close() // 让所有阻塞的 Read 调用返回 EOF
		}
		close(c.incoming)
	}()

	for {
		// 1. 读取帧类型
		var frameType byte
		if err := binary.Read(c.netConn, binary.BigEndian, &frameType); err != nil {
			return // 连接断开或出错
		}

		// 2. 读取 Key
		var keyLen uint16
		if err := binary.Read(c.netConn, binary.BigEndian, &keyLen); err != nil {
			return
		}
		keyBuf := make([]byte, keyLen)
		if _, err := io.ReadFull(c.netConn, keyBuf); err != nil {
			return
		}
		key := string(keyBuf)

		// 3. 根据帧类型处理
		switch frameType {
		case FrameTypeData:
			// 读取数据
			var dataLen uint32
			if err := binary.Read(c.netConn, binary.BigEndian, &dataLen); err != nil {
				return
			}

			// 查找或创建这个 key 对应的 pipe writer
			writer, ok := c.readers[key]
			if !ok {
				// 这是一个新的 key，创建 pipe
				pr, pw := io.Pipe()
				c.readers[key] = pw
				// 将新的 reader 通过 channel 发送给 Receive 调用
				c.incoming <- incomingStream{key: key, reader: pr}
				writer = pw
			}

			// 将数据拷贝到 pipe writer，这样另一端的 pipe reader 就能读到
			if _, err := io.CopyN(writer, c.netConn, int64(dataLen)); err != nil {
				return
			}

		case FrameTypeEnd:
			// 对端告知这个 key 的数据传输完毕
			if writer, ok := c.readers[key]; ok {
				writer.Close()         // 关闭 pipe writer，reader 会收到 EOF
				delete(c.readers, key) // 清理资源
			}
		}
	}
}

// NewConn 从一个 TCP 连接得到一个你实现的连接对象
func NewConn(conn net.Conn) *Conn {
	c := &Conn{
		netConn:  conn,
		incoming: make(chan incomingStream), // 无缓冲 channel
		readers:  make(map[string]*io.PipeWriter),
	}
	// 关键：启动后台读取循环
	go c.readLoop()
	return c
}

// connWriter 是 Send 方法返回的 writer 实现
type connWriter struct {
	conn *Conn
	key  string
}

func (cw *connWriter) Write(p []byte) (n int, err error) {
	// 加锁保证并发安全
	cw.conn.writeMutex.Lock()
	defer cw.conn.writeMutex.Unlock()

	// 写入帧头
	if err := binary.Write(cw.conn.netConn, binary.BigEndian, FrameTypeData); err != nil {
		return 0, err
	}
	// 写key
	keyBytes := []byte(cw.key)
	if err := binary.Write(cw.conn.netConn, binary.BigEndian, uint16(len(keyBytes))); err != nil {
		return 0, err
	}
	if _, err := cw.conn.netConn.Write(keyBytes); err != nil {
		return 0, err
	}
	// 写 data len
	if err := binary.Write(cw.conn.netConn, binary.BigEndian, uint32(len(p))); err != nil {
		return 0, err
	}

	// 写入数据
	return cw.conn.netConn.Write(p)
}

func (cw *connWriter) Close() error {
	cw.conn.writeMutex.Lock()
	defer cw.conn.writeMutex.Unlock()

	// 发送结束帧
	if err := binary.Write(cw.conn.netConn, binary.BigEndian, FrameTypeEnd); err != nil {
		return err
	}
	// 当前key
	keyBytes := []byte(cw.key)
	if err := binary.Write(cw.conn.netConn, binary.BigEndian, uint16(len(keyBytes))); err != nil {
		return err
	}
	_, err := cw.conn.netConn.Write(keyBytes)
	return err
}

func (c *Conn) Send(key string) (writer io.WriteCloser, err error) {
	// Send 的实现很简单，就是返回一个包装了 Conn 和 key 的 writer
	return &connWriter{conn: c, key: key}, nil
}

func (c *Conn) Receive() (key string, reader io.Reader, err error) {
	// 从 channel 接收一个数据流，如果 channel 中没有，则阻塞在这里
	stream, ok := <-c.incoming
	if !ok {
		// channel 已关闭，意味着连接已断开
		return "", nil, io.EOF
	}
	return stream.key, stream.reader, nil
}

func (c *Conn) Close() {
	c.netConn.Close()
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
		// 注意: 这里测试代码没有调用 writer.Close(), 在真实场景中应该调用
		// 但由于马上调用 conn.Close(), 底层连接关闭，对端 ReadAll 会正常结束
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
	if err != nil && err != io.EOF {
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
	// 增加 buf 大小以提高性能
	var buf = make([]byte, 32*1024) //调用者读取时的 buf 大小不是固定的，你的实现中不可假定 buf 为固定值
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

	multiWriter := io.MultiWriter(writer, hash)

	for i := 0; i < dataSize/bufSize; i++ {
		_, err := rand.Read(buf)
		if err != nil {
			panic(err)
		}
		n, err := multiWriter.Write(buf)
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
		mapKeyToChecksum = make(map[string]string) // 使用 make 初始化 map
		lock             sync.Mutex
	)
	ln := startServer(func(conn *Conn) {
		// 服务端等待客户端进行传输
		key, reader, err := conn.Receive()
		if err != nil {
			if err == io.EOF { // 客户端可能先关闭连接，这是正常情况
				return
			}
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
	fmt.Println("Running testCase0...")
	testCase0()
	fmt.Println("testCase0 passed.")
	fmt.Println("Running testCase1...")
	testCase1()
	fmt.Println("testCase1 passed.")
}
