package adapter

import "fmt"

type IRun interface {
	CrateServer(str string, cpu string)
}

type aliyun struct{}

func (receiver aliyun) RunT(str string, cpu string) {
	fmt.Println(str, "run in ", cpu, " cpu")
}

type AliyunAdapter struct {
	aliyun aliyun
}

func (a AliyunAdapter) CrateServer(str string, cpu string) {
	a.aliyun.RunT(str, cpu)
}

type tenxunyun struct{}

func (receiver tenxunyun) RunT(str string, cpu string) {
	fmt.Println(str, "run in ", cpu, " cpu")
}

type tenxunyunAdapter struct {
	tenxunyun tenxunyun
}

func (a tenxunyunAdapter) CrateServer(str string, cpu string) {
	a.tenxunyun.RunT(str, cpu)
}
