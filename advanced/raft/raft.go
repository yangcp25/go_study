package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/hashicorp/raft"
	bolt "github.com/hashicorp/raft-boltdb"
)

// 简单的 FSM：提交即打印
type FSM struct{}

func (f *FSM) Apply(l *raft.Log) interface{} {
	fmt.Printf("Apply: %s\n", string(l.Data))
	return nil
}
func (f *FSM) Snapshot() (raft.FSMSnapshot, error) { return &snapshot{}, nil }
func (f *FSM) Restore(io.ReadCloser) error         { return nil }

type snapshot struct{}

func (s *snapshot) Persist(sink raft.SnapshotSink) error { return sink.Close() }
func (s *snapshot) Release()                             {}

func main() {
	// 1) 配置
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(os.Args[1]) // 节点 ID 来自第一个参数

	// 2) 网络传输：TCP
	addr, err := net.ResolveTCPAddr("tcp", os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	transport, err := raft.NewTCPTransport(os.Args[2], addr, 3, 10*time.Second, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

	// 3) 日志与快照存储
	store, err := bolt.NewBoltStore(fmt.Sprintf("raft-%s.db", os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}
	snapshotStore := raft.NewInmemSnapshotStore()

	// 4) 创建 Raft 实例
	r, err := raft.NewRaft(config, &FSM{}, store, store, snapshotStore, transport)
	if err != nil {
		log.Fatal(err)
	}

	// 5) Bootstrap 第一个节点
	if os.Args[1] == "node1" {
		cfg := raft.Configuration{
			Servers: []raft.Server{
				{ID: "node1", Address: transport.LocalAddr()},
				{ID: "node2", Address: raft.ServerAddress("127.0.0.1:12002")},
				{ID: "node3", Address: raft.ServerAddress("127.0.0.1:12003")},
			},
		}
		r.BootstrapCluster(cfg)
	}

	// 6) 简单命令行提交
	if config.LocalID == "node1" {
		go func() {
			for {
				time.Sleep(3 * time.Second)
				f := r.Apply([]byte("hello raft"), 5*time.Second)
				if err := f.Error(); err != nil {
					log.Println("apply error:", err)
				}
			}
		}()
	}

	select {} // 阻塞
}
