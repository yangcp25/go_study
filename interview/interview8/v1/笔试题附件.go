package main

import (
	"bufio"
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
type Conn struct {
	conn     net.Conn
	reader   *bufio.Reader
	sendLock sync.Mutex
	recvLock sync.Mutex
}

// streamWriter 实现了 io.WriteCloser 接口，用于发送数据。
type streamWriter struct {
	conn   *Conn
	closed bool
}

// Write 将数据 p 封装成一个数据帧并发送。
func (sw *streamWriter) Write(p []byte) (n int, err error) {
	if sw.closed {
		return 0, fmt.Errorf("writer is closed")
	}
	if len(p) == 0 {
		return 0, nil
	}

	// 锁在 Send 方法中已经获取，这里直接使用
	if err := writeFrame(sw.conn.conn, p); err != nil {
		return 0, err
	}
	return len(p), nil
}

// Close 发送一个结束帧，并释放写锁。
func (sw *streamWriter) Close() error {
	if sw.closed {
		return nil
	}
	sw.closed = true
	// 发送长度为0的帧作为结束标志
	err := writeFrame(sw.conn.conn, nil)
	sw.conn.sendLock.Unlock() // 释放锁
	return err
}

// Send 传入一个 key 表示发送者将要传输的数据对应的标识；
// 返回 writer 可供发送者分多次写入大量该 key 对应的数据；
// 当发送者已将该 key 对应的所有数据写入后，调用 writer.Close 告知接收者：该 key 的数据已经完全写入；
func (conn *Conn) Send(key string) (writer io.WriteCloser, err error) {
	conn.sendLock.Lock() // 获取锁，直到 writer.Close() 才释放

	// 发送 Key Frame
	if err := writeFrame(conn.conn, []byte(key)); err != nil {
		conn.sendLock.Unlock()
		return nil, err
	}

	return &streamWriter{conn: conn}, nil
}

// streamReader 实现了 io.Reader 接口，用于接收数据。
type streamReader struct {
	conn             *Conn
	remainingInChunk int // 当前数据块中剩余的字节数
}

// Read 从数据帧中读取数据。
func (sr *streamReader) Read(p []byte) (n int, err error) {
	// 加锁以保护对底层 reader 的并发访问
	sr.conn.recvLock.Lock()
	defer sr.conn.recvLock.Unlock()

	// 如果当前数据块已读完，则读取下一个数据块的头部
	if sr.remainingInChunk == 0 {
		// 读取下一个数据块的长度
		var length uint32
		if err := binary.Read(sr.conn.reader, binary.BigEndian, &length); err != nil {
			return 0, err // 可能返回 io.EOF
		}

		// 如果长度为0，说明这个 key 的数据流结束了
		if length == 0 {
			return 0, io.EOF
		}
		sr.remainingInChunk = int(length)
	}

	// 确定本次读取的字节数
	bytesToRead := len(p)
	if bytesToRead > sr.remainingInChunk {
		bytesToRead = sr.remainingInChunk
	}

	// 从底层连接读取数据
	n, err = io.ReadFull(sr.conn.reader, p[:bytesToRead])
	if err != nil {
		// 如果发生非预期错误，也可能需要将 remainingInChunk 清零
		return n, err
	}

	sr.remainingInChunk -= n
	return n, nil
}

// Receive 返回一个 key 表示接收者将要接收到的数据对应的标识；
// 返回的 reader 可供接收者多次读取该 key 对应的数据；
// 当 reader 返回 io.EOF 错误时，表示接收者已经完整接收该 key 对应的数据；
func (conn *Conn) Receive() (key string, reader io.Reader, err error) {
	conn.recvLock.Lock()
	defer conn.recvLock.Unlock()

	// 读取 Key Frame
	keyBytes, err := readFramePayload(conn.reader)
	if err != nil {
		return "", nil, err // 如果连接关闭，这里会正确返回 io.EOF
	}

	return string(keyBytes), &streamReader{conn: conn}, nil
}

// Close 关闭你实现的连接对象及其底层的 TCP 连接
func (conn *Conn) Close() {
	conn.conn.Close()
}

// NewConn 从一个 TCP 连接得到一个你实现的连接对象
func NewConn(conn net.Conn) *Conn {
	return &Conn{
		conn:   conn,
		reader: bufio.NewReader(conn),
	}
}

// 辅助函数：写入一个数据帧（长度 + 数据）
func writeFrame(w io.Writer, data []byte) error {
	length := uint32(len(data))
	if err := binary.Write(w, binary.BigEndian, length); err != nil {
		return err
	}
	if length > 0 {
		_, err := w.Write(data)
		return err
	}
	return nil
}

// 辅助函数：读取一个数据帧的载荷
func readFramePayload(r io.Reader) ([]byte, error) {
	var length uint32
	if err := binary.Read(r, binary.BigEndian, &length); err != nil {
		return nil, err
	}
	if length == 0 {
		return []byte{}, nil
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
