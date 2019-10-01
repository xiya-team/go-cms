package routers

import (
	"github.com/astaxie/beego"
	"go-cms/controllers/sys"
)

func init() {
	//ns := beego.NewNamespace("/api",
	//	beego.NSNamespace("/user",
	//		beego.NSInclude(
	//			&sys.UserController{},
	//		),
	//	),
	//)
	//beego.AddNamespace(ns)
	
	beego.Router("/api/user/login", &sys.UserController{}, "*:Login")
	//beego.Router("/api/user/check", &sys.UserController{}, "*:DoCheck")
}