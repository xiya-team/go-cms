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

	//用户相关
	beego.Router("/api/user/login", &sys.UserController{}, "post:Login")
	beego.Router("/api/user/create", &sys.UserController{}, "post:Create")
	beego.Router("/api/user/info", &sys.UserController{}, "get:UserInfo") // 获取用户消息
	beego.Router("/api/user/list", &sys.UserController{}, "get:UserList") // 获取用户列表
	beego.Router("/api/user/check_token", &sys.UserController{}, "post:CheckToken")
	beego.Router("/api/user/logout", &sys.UserController{}, "post:Logout")
	
	//验证码校验
	beego.Router("/api/captcha/check", &sys.CaptchaController{}, "post:Hander")
	
	//参数设置
	beego.Router("/api/configs/index", &sys.ConfigsController{}, "post:Index")
	beego.Router("/api/configs/create", &sys.ConfigsController{}, "post:Create")
	beego.Router("/api/configs/update", &sys.ConfigsController{}, "post:Update")
	beego.Router("/api/configs/delete", &sys.ConfigsController{}, "post:Delete")
	
	//岗位管理
	beego.Router("/api/post/index", &sys.PostController{}, "post:Index")
	beego.Router("/api/post/create", &sys.PostController{}, "post:Create")
	beego.Router("/api/post/update", &sys.PostController{}, "post:Update")
	beego.Router("/api/post/delete", &sys.PostController{}, "post:Delete")
	
	//菜单管理
	beego.Router("/api/menu/index", &sys.MenuController{}, "post:Index")
	beego.Router("/api/menu/create", &sys.MenuController{}, "post:Create")
	beego.Router("/api/menu/update", &sys.MenuController{}, "post:Update")
	beego.Router("/api/menu/delete", &sys.MenuController{}, "post:Delete")
	
	//字典管理
	beego.Router("/api/dict/index", &sys.DictTypeController{}, "post:Index")
	beego.Router("/api/dict/create", &sys.DictTypeController{}, "post:Create")
	beego.Router("/api/dict/update", &sys.DictTypeController{}, "post:Update")
	beego.Router("/api/dict/delete", &sys.DictTypeController{}, "post:Delete")
	//字典数据管理
	beego.Router("/api/dictData/index", &sys.DictDataController{}, "post:Index")
	beego.Router("/api/dictData/create", &sys.DictDataController{}, "post:Create")
	beego.Router("/api/dictData/update", &sys.DictDataController{}, "post:Update")
	beego.Router("/api/dictData/delete", &sys.DictDataController{}, "post:Delete")
}
