package chain

import "fmt"

type IHandle interface {
	Filter(content string) bool
}

type SensitiveWordFilterChain struct {
	handler []IHandle
}

func (s *SensitiveWordFilterChain) addFilter(filter IHandle) {
	s.handler = append(s.handler, filter)
}

func (s SensitiveWordFilterChain) Filter(content string) bool {
	for _, handle := range s.handler {
		if handle.Filter(content) {
			return true
		}
	}

	return false
}

type filter1 struct{}

func (f filter1) Filter(content string) bool {
	// todo
	fmt.Println("过滤1", content)
	return false
}

type filter2 struct{}

func (f filter2) Filter(content string) bool {
	// todo
	fmt.Println("过滤2")
	return false
}
