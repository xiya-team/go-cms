package common

import (
	"github.com/astaxie/beego/context"
)
//包循环调用？再开一个包
var Fc *context.Context
var UserId int