package mian

import (
	"context"
	"fmt"
	"net/http"
)

//const (
//	KEY = "trace_id"
//)
//
//func NewRequestID() string {
//	return strings.Replace("test", "-", "", -1)
//}
//
//func NewContextWithTraceID() context.Context {
//	ctx := context.WithValue(context.Background(), KEY, NewRequestID())
//	return ctx
//}
//
//func PrintLog(ctx context.Context, message string) {
//	fmt.Printf("%s|info|trace_id=%s|%s", time.Now().Format("2006-01-02 15:04:05"), GetContextValue(ctx, KEY), message)
//}
//
//func GetContextValue(ctx context.Context, k string) string {
//	v, ok := ctx.Value(k).(string)
//	if !ok {
//		return ""
//	}
//	return v
//}
//
//func ProcessEnter(ctx context.Context) {
//	PrintLog(ctx, "Golang梦工厂")
//}
//
////func main() {
////	ProcessEnter(NewContextWithTraceID())
////}

const requestIDKey int = 0

func WithRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			// 从 header 中提取 request-id
			reqID := req.Header.Get("X-Request-ID")
			// 创建 valueCtx。使用自定义的类型，不容易冲突
			ctx := context.WithValue(
				req.Context(), requestIDKey, reqID)

			// 创建新的请求
			req = req.WithContext(ctx)

			// 调用 HTTP 处理函数
			next.ServeHTTP(rw, req)
		},
	)
}

// 获取 request-id
func GetRequestID(ctx context.Context) string {
	s := ctx.Value(requestIDKey).(string)
	return s
}

func Handle(rw http.ResponseWriter, req *http.Request) {
	// 拿到 reqId，后面可以记录日志等等
	reqID := GetRequestID(req.Context())
	fmt.Println(reqID)
}

func main() {
	handler := WithRequestID(http.HandlerFunc(Handle))
	http.ListenAndServe("/", handler)
}
