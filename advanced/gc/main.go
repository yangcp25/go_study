package main

import (
	"fmt"
	"net/http"
	"sync"
)

package main

import (
"fmt"
"net/http"
_ "net/http/pprof"
"sync"
)

// 使用 sync.Pool 复用需要的 buf
var bufPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 10<<20) // 10 MB 大小的缓冲区
	},
}

func main() {
	// 开启 pprof 调试服务器
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	// 注册 /example2 路由
	http.HandleFunc("/example2", func(w http.ResponseWriter, r *http.Request) {
		// 从 Pool 中获取一个缓冲区
		b := bufPool.Get().([]byte)

		// 清空缓冲区
		for idx := range b {
			b[idx] = 0
		}

		// 模拟处理
		fmt.Fprintf(w, "done, %v", r.URL.Path[1:])

		// 将缓冲区放回 Pool 中以供后续复用
		bufPool.Put(b)
	})

	// 启动服务器
	fmt.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
