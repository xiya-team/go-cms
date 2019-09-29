package e

import "github.com/astaxie/beego"

var MsgFlags = map[string]string{
	"landing successfully":           "登陆成功",
	"landing failed":                 "登陆失败",
}

func T(msg string)(str string) {
	if beego.AppConfig.String("i18n") == "ch" {
		return MsgFlags[msg]
	}
	if beego.AppConfig.String("i18n") == "us" {
		return msg
	}
	return
}
