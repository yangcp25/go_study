package proxy

import (
	"fmt"
	"time"
)

type userInterface interface {
	Login()
}
type user struct{}

func (u *user) Login() {
	//todo
	fmt.Println("---")
	time.Sleep(3 * time.Second)
}

type userProxy struct {
	user *user
}

func newUserProxy(user *user) *userProxy {
	return &userProxy{
		user: user,
	}
}

func (receiver *userProxy) Login() {
	start := time.Now()
	receiver.user.Login()
	duration := time.Now().Sub(start)
	fmt.Println(duration)
}
