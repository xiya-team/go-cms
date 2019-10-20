package wechat


import (
	"fmt"
	"go-cms/controllers"
	"go-cms/pkg/e"
	"log"
	"strconv"
	"time"
	
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/jssdk"
	"gopkg.in/chanxuehong/wechat.v2/mp/menu"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/callback/request"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/callback/response"
	mpoauth2 "gopkg.in/chanxuehong/wechat.v2/mp/oauth2"
	"gopkg.in/chanxuehong/wechat.v2/oauth2"
	"gopkg.in/chanxuehong/wechat.v2/util"
)

const (
	wxAppId         = "wx1626a5379b07dfc6" //your appId
	wxAppSecret     = "f382d85a5831e74668ae30c68c47a7bf" //your appSecret
	wxOriId         = "gh_b1e73d822fc3" //原始ID
	wxToken         = "Dswq1322s1dfsf31s2af321231rew" //token
	wxEncodedAESKey = "WWWJopR6VoDa7h4QfBGJsXmjMzUNq1aDrI2kQELHCyW"
)

var (
	// 下面两个变量不一定非要作为全局变量, 根据自己的场景来选择.
	msgHandler     core.Handler
	msgServer      *core.Server
	oauth2Endpoint oauth2.Endpoint = mpoauth2.NewEndpoint(wxAppId, wxAppSecret)
)

type WxSignature struct {
	AppID     string `json:"appId"`
	Noncestr  string `json:"noncestr"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
	Url       string `json:"url"`
}

type WechatController struct {
	controllers.BaseController
}

func init() {
	mux := core.NewServeMux()
	mux.DefaultMsgHandleFunc(defaultMsgHandler)
	mux.DefaultEventHandleFunc(defaultEventHandler)
	mux.MsgHandleFunc(request.MsgTypeText, textMsgHandler)
	mux.EventHandleFunc(menu.EventTypeClick, menuClickEventHandler)

	msgHandler = mux
	msgServer = core.NewServer(wxOriId, wxAppId, wxToken, wxEncodedAESKey, msgHandler, nil)
}

func textMsgHandler(ctx *core.Context) {
	log.Printf("收到文本消息:\n%s\n", ctx.MsgPlaintext)

	msg := request.GetText(ctx.MixedMsg)
	resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, msg.Content)
	ctx.RawResponse(resp) // 明文回复
	//ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func defaultMsgHandler(ctx *core.Context) {
	log.Printf("收到消息:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func menuClickEventHandler(ctx *core.Context) {
	log.Printf("收到菜单 click 事件:\n%s\n", ctx.MsgPlaintext)

	event := menu.GetClickEvent(ctx.MixedMsg)
	resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, "收到 click 类型的事件")
	ctx.RawResponse(resp) // 明文回复
	//ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func defaultEventHandler(ctx *core.Context) {
	log.Printf("收到事件:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

// wxCallbackHandler 是处理回调请求的 http handler.
func (w *WechatController) WxCallbackHandler() {
	//log.Printf("回调处理:\n%s\n", w.Ctx.Request)
	msgServer.ServeHTTP(w.Ctx.ResponseWriter, w.Ctx.Request, nil)
}

// 通过code获取用户openid及用户基本信息
// @router /get_userinfo [post]
func (w *WechatController) GetUserInfo() {
	code := w.GetString("code")

	oauth2Client := oauth2.Client{
		Endpoint: oauth2Endpoint,
	}
	token, err := oauth2Client.ExchangeToken(code)
	if err != nil {
		log.Println(err)
		return
	}

	userinfo, err := mpoauth2.GetUserInfo(token.AccessToken, token.OpenId, "", nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("userinfo: %+v\r\n", userinfo)
	
	w.JsonResult(e.SUCCESS, "成功",userinfo)
}

// Desc: 自定义分享jsApiticket配置参数
// @router /get_sign [get]
func (w *WechatController) GetSign() {
	var (
		wxSignature       WxSignature
		accessTokenServer core.AccessTokenServer = core.NewDefaultAccessTokenServer(wxAppId, wxAppSecret, nil)
		wechatClient      *core.Client           = core.NewClient(accessTokenServer, nil)
	)

	var ticketServer = jssdk.NewDefaultTicketServer(wechatClient)

	//fmt.Println(base.GetCallbackIP(wechatClient))

	jsapiTicket, err := ticketServer.Ticket()
	if err != nil {
		fmt.Println(err)
	}

	nonceStr := util.NonceStr()
	timestamp := strconv.FormatInt(time.Now().UTC().Unix(), 10)
	url := "http://127.0.0.1/share.html"

	signature := jssdk.WXConfigSign(jsapiTicket, nonceStr, timestamp, url)

	wxSignature.AppID = wxAppId
	wxSignature.Noncestr = nonceStr
	wxSignature.Timestamp = timestamp
	wxSignature.Signature = signature
	wxSignature.Url = url
	
	w.JsonResult(e.SUCCESS, "成功",wxSignature)
}


//func (w *WechatController) Signature() {
//
//	//微信接入验证 这是首次对接微信 填写url后 微信服务器会发一个请求过来
//	//c.Ctx.Request.URL-------------wx_connect?signature=038d75ed5485b9881a01b3b93e85f9fff28ea739&echostr=5756456183388806654&timestamp=1476173150&nonce=1093541731
//
//	//开发者提交信息(包括URL、Token)后，微信服务器将发送Http Get请求到填写的URL上，
//	//GET请求携带四个参数：signature、timestamp、nonce和echostr。公众号服务程序应该按如下要求进行接入验证
//	timestamp, nonce,signatureIn := w.GetString("timestamp"), w.GetString("nonce"),w.GetString("signature")
//	signatureGen := makeSignature(timestamp, nonce)
//
//	//将加密后获得的字符串与signature对比，如果一致，说明该请求来源于微信
//	if signatureGen != signatureIn {
//		fmt.Printf("signatureGen != signatureIn signatureGen=%s,signatureIn=%s\n", signatureGen, signatureIn)
//		w.Ctx.WriteString("")
//	} else {
//		//如果请求来自于微信，则原样返回echostr参数内容 以上完成后，接入验证就会生效，开发者配置提交就会成功。
//		echostr := w.GetString("echostr")
//		w.Ctx.WriteString(echostr)
//	}
//}
//
//func makeSignature(timestamp, nonce string) string {
//
//	//1. 将 plat_token、timestamp、nonce三个参数进行字典序排序
//	sl := []string{wxToken, timestamp, nonce}
//	sort.Strings(sl)
//	//2. 将三个参数字符串拼接成一个字符串进行sha1加密
//	s := sha1.New()
//	io.WriteString(s, strings.Join(sl, ""))
//
//	return fmt.Sprintf("%x", s.Sum(nil))
//}