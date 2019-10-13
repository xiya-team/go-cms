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

	beego.Router("/api/user/login", &sys.UserController{}, "post:Login")
	beego.Router("/api/user/create", &sys.UserController{}, "post:Create")
	beego.Router("/api/user/info", &sys.UserController{}, "get:UserInfo") // 获取用户消息
	beego.Router("/api/user/list", &sys.UserController{}, "get:UserList") // 获取用户列表
	beego.Router("/api/user/check_token", &sys.UserController{}, "post:CheckToken")
	beego.Router("/api/user/logout", &sys.UserController{}, "get:Logout")
	
	beego.Router("/api/captcha/check", &sys.CaptchaController{}, "post:Hander")
	
	//参数设置
	beego.Router("/api/configs/index", &sys.ConfigsController{}, "post:Index")
	beego.Router("/api/configs/create", &sys.ConfigsController{}, "post:Create")
	beego.Router("/api/configs/update", &sys.ConfigsController{}, "post:Update")
	beego.Router("/api/configs/delete", &sys.ConfigsController{}, "post:Delete")
}
