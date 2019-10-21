package wx

import (
	"fmt"
	"go-cms/controllers"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
)

//const (
//	wxAppId         = "wx1626a5379b07dfc6" //your appId
//	wxAppSecret     = "f382d85a5831e74668ae30c68c47a7bf" //your appSecret
//	wxOriId         = "gh_b1e73d822fc3" //原始ID
//	wxToken         = "Dswq1322s1dfsf31s2af321231rew" //token
//	wxEncodedAESKey = "WWWJopR6VoDa7h4QfBGJsXmjMzUNq1aDrI2kQELHCyW"
//)

type WxConfigController struct {
	controllers.BaseController
}

func (c *WxConnectController) hello() {
	//配置微信参数
	config := &wechat.Config{
		AppID:          "your app id",
		AppSecret:      "your app secret",
		Token:          "your token",
		EncodingAESKey: "your encoding aes key",
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
