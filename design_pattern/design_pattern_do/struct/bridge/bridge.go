package bridge

import "fmt"

type sendInterface interface {
	Send(msg string)
}

type emailArray struct {
	email []string
}

func (receiver emailArray) Send(msg string) {
	// todo
	for _, v := range receiver.email {
		fmt.Println("给", v, ":", msg)
	}
}

func NewEmailArray(email []string) *emailArray {
	return &emailArray{
		email: email,
	}
}

type notifySend interface {
	notifyTo(msg string)
}
type notify struct {
	sender sendInterface
}

func newNotify(sender sendInterface) *notify {
	return &notify{sender: sender}
}

func (receiver notify) notifyTo(msg string) {
	receiver.sender.Send(msg)
}
