package common

import "github.com/astaxie/beego/context"

//包循环调用？再开一个包
var UserId int = 0
var UserName string = "admin"
var Ctx *context.Context