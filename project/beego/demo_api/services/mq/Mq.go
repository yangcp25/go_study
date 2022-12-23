package mq

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
)

func Connect() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	return conn, err
}

func Publish(exchange string, queueName string, body string) error {
	// 建立连接
	conn, err := Connect()

	if err != nil {
		return err
	}

	defer conn.Close()

	// 创建通道

	channel, err := conn.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	// 创建队列
	q, err := channel.QueueDeclare(
		queueName,
		true, // 持久化
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// 发送消息
	channel.Publish(exchange, q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(body),
	})

	return err
}

type CallBack func(msg string)

func Consumer(exchange string, queueName string, callback CallBack) {
	// 建立连接
	conn, err := Connect()
	if err != nil {
		return
	}
	defer conn.Close()

	// 创建通道
	channel, err := conn.Channel()
	if err != nil {
		return
	}
	defer channel.Close()

	// 创建队列
	q, err := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)

	msgs, err := channel.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		return
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			s := BytesToString(&(d.Body))
			callback(*s)
			d.Ack(false)
		}
	}()
	fmt.Printf("wait for msg")
	<-forever
}

func BytesToString(i *[]byte) *string {
	str := bytes.NewBuffer(*i)
	r := str.String()
	return &r
}

// PublishEx 订阅模式
func PublishEx(exchange string, types string, routingKey string, body string) error {
	// 建立连接
	conn, err := Connect()

	if err != nil {
		return err
	}

	defer conn.Close()

	// 创建通道
	channel, err := conn.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	// 创建交换机
	err = channel.ExchangeDeclare(
		exchange,
		types, // 持久化
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// 发送消息
	channel.Publish(exchange, routingKey, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(body),
	})

	return err
}

func ConsumerEx(exchange string, types string, routingKey string, callback CallBack) {
	// 建立连接
	conn, err := Connect()
	if err != nil {
		return
	}
	defer conn.Close()

	// 创建通道
	channel, err := conn.Channel()
	if err != nil {
		return
	}
	defer channel.Close()

	// 创建
	err = channel.ExchangeDeclare(
		exchange,
		types,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}

	// 创建队列
	q, err := channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)

	if err != nil {
		return
	}

	err = channel.QueueBind(q.Name, routingKey, exchange, false, nil)
	if err != nil {
		return
	}

	msgs, err := channel.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		return
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			s := BytesToString(&(d.Body))
			callback(*s)
			d.Ack(false)
		}
	}()
	fmt.Printf("wait for msg")
	<-forever
}

// ConsumerDlx 死信队列
func ConsumerDlx(exchangeA string, queueA string, exchangeB string, queueB string, ttl int, callback CallBack) {
	con, err := Connect()
	defer con.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	channel, err := con.Channel()
	defer channel.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = channel.ExchangeDeclare(
		exchangeA,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	queueAInfo, err := channel.QueueDeclare(queueA, true, false, false, false, amqp.Table{
		"x-message-ttl":          ttl,
		"x-dead-letter-exchange": exchangeB,
		//"x-dead-letter-queue" : "",
		//"x-dead-letter-routing-key":"",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	err = channel.QueueBind(queueAInfo.Name, "", exchangeA, false, nil)
	if err != nil {
		return
	}

	err = channel.ExchangeDeclare(exchangeB, "fanout", true, false, false, false, nil)
	if err != nil {
		return
	}

	queueBInfo, err := channel.QueueDeclare(queueB, true, false, false, false, nil)

	err = channel.QueueBind(queueBInfo.Name, "", exchangeB, false, nil)
	if err != nil {
		return
	}

	msg, err := channel.Consume(queueBInfo.Name, "", false, false, false, false, nil)
	if err != nil {
		return
	}

	forever := make(chan bool)
	go func() {
		for d := range msg {
			s := BytesToString(&(d.Body))
			callback(*s)
			err := d.Ack(false)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}()
	<-forever
}

// 死信生产端
func PublishDlx(exchangeA string, body string) error {
	conn, _ := Connect()
	defer conn.Close()
	channel, _ := conn.Channel()
	defer channel.Close()

	// 创建交换机
	err := channel.ExchangeDeclare(
		exchangeA,
		"fanout", // 持久化
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = channel.Publish(exchangeA, "", false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(body),
	})
	if err != nil {
		return err
	}
	return nil
}
