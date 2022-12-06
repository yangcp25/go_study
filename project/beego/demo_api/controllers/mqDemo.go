package controllers

import (
	"demo_api/services/mq"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
	"time"
)

type MqDemoController struct {
	beego.Controller
}

func (ctrl *MqDemoController) Test() {
	go func() {
		count := 0
		for {
			mq.Publish("", "ycp_mq_test", "haha_"+strconv.Itoa(count))
			count++
			time.Sleep(1 * time.Second)
		}
	}()
	ctrl.Ctx.WriteString("test")
}

// 订阅模式
func (ctrl *MqDemoController) TestEx() {
	go func() {
		count := 0
		for {
			mq.PublishEx("ycp_test.demo.fanout", "fanout", "", "fanout_"+strconv.Itoa(count))
			count++
			time.Sleep(1 * time.Second)
		}
	}()
	ctrl.Ctx.WriteString("testEx")
}

// 路由模式
func (ctrl *MqDemoController) TestDirect() {
	go func() {
		count := 0
		for {
			if count%2 == 0 {
				mq.PublishEx("ycp_test.demo.direct", "direct", "two", "fanout_two_"+strconv.Itoa(count))
			} else {
				mq.PublishEx("ycp_test.demo.direct", "direct", "one", "fanout_one_"+strconv.Itoa(count))
			}
			count++
			time.Sleep(1 * time.Second)
		}
	}()
	ctrl.Ctx.WriteString("testDirect")
}

// 主题模式
func (ctrl *MqDemoController) TestTopic() {
	go func() {
		count := 0
		for {
			if count%2 == 0 {
				mq.PublishEx("ycp_test.demo.topic", "topic", "a.test.two", "a_topic_two_"+strconv.Itoa(count))
			} else {
				mq.PublishEx("ycp_test.demo.topic", "topic", "a.test.one", "a_topic_one_"+strconv.Itoa(count))
			}
			count++
			time.Sleep(1 * time.Second)
		}
	}()
	ctrl.Ctx.WriteString("TestTopic")
}

func (ctrl *MqDemoController) TestTopicTwo() {
	go func() {
		count := 0
		for {
			if count%2 == 0 {
				mq.PublishEx("ycp_test.demo.topic", "topic", "test.two", "topic_two_"+strconv.Itoa(count))
			} else {
				mq.PublishEx("ycp_test.demo.topic", "topic", "test.one", "topic_one_"+strconv.Itoa(count))
			}
			count++
			time.Sleep(1 * time.Second)
		}
	}()
	ctrl.Ctx.WriteString("TestTopicTwo")
}
