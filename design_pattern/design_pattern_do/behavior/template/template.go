package template

import "fmt"

type ISms interface {
	SendMsg(phone string, content string) error
}

type smg struct {
	ISms
}

func (receiver *smg) Valid(content string) error {
	if len(content) > 64 {
		return fmt.Errorf("超过限制")
	}
	return nil
}

func (receiver *smg) Send(phone string, content string) error {
	if err := receiver.Valid(content); err != nil {
		return err
	}

	return receiver.SendMsg(phone, content)
}

type TelecomSms struct {
	*smg
}

func NewTelecomSms() *TelecomSms {
	tel := &TelecomSms{}
	// 这里有点绕，是因为 go 没有继承，用嵌套结构体的方法进行模拟
	// 这里将子类作为接口嵌入父类，就可以让父类的模板方法 Send 调用到子类的函数
	// 实际使用中，我们并不会这么写，都是采用组合+接口的方式完成类似的功能
	tel.smg = &smg{ISms: tel}
	return tel
}

func (receiver *TelecomSms) SendMsg(phone string, content string) error {
	fmt.Println("走了电信接口")
	return nil
}
