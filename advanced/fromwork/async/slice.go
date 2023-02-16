package async

import "sync"

type safeSlice struct {
	data  []interface{}
	mutex sync.Mutex
}
