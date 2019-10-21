package wx

import (
	"fmt"
	"github.com/astaxie/beego"
	"go-cms/controllers"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
)

type WxConfigController struct {
	controllers.BaseController
}

func (c *WxConfigController) Get() {
	//配置微信参数
	config := &wechat.Config{
		AppID         : beego.AppConfig.String("wechat::AppID"),
		AppSecret     : beego.AppConfig.String("wechat::AppSecret"),
		Token         : beego.AppConfig.String("wechat::Token"),
		EncodingAESKey: beego.AppConfig.String("wechat::EncodingAESKey"),
	}
	wc := wechat.NewWechat(config)

	// 传入request和responseWriter
	server := wc.GetServer(c.Ctx.Request, c.Ctx.ResponseWriter)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}
