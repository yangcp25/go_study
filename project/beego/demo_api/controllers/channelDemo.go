package controllers

import (
	"demo_api/services/mq"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
	"time"
)

type ChannelController struct {
	beego.Controller
}

func (ctrl *ChannelController) Test() {
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
