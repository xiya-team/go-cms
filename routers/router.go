package routers

import (
	"github.com/astaxie/beego"
	"go-cms/controllers/commons"
	"go-cms/controllers/sys"
	"go-cms/controllers/wx"
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
	beego.Router("/api/common/page_not_found", &commons.CommonController{}, "*:PageNotFound")

	//用户相关
	beego.Router("/api/user/login", &sys.UserController{}, "post:Login")
	beego.Router("/api/user/create", &sys.UserController{}, "post:Create")
	beego.Router("/api/user/info", &sys.UserController{}, "get:UserInfo") // 获取用户消息
	beego.Router("/api/user/list", &sys.UserController{}, "post:Index") // 获取用户列表
	beego.Router("/api/user/check_token", &sys.UserController{}, "post:CheckToken")
	beego.Router("/api/user/logout", &sys.UserController{}, "post:Logout")

	//验证码校验
	beego.Router("/api/captcha/check", &sys.CaptchaController{}, "post:Hander")

	//参数设置
	beego.Router("/api/configs/index", &sys.ConfigsController{}, "*:Index")
	beego.Router("/api/configs/create", &sys.ConfigsController{}, "post:Create")
	beego.Router("/api/configs/update", &sys.ConfigsController{}, "put:Update")
	beego.Router("/api/configs/delete", &sys.ConfigsController{}, "delete:Delete")

	//岗位管理
	beego.Router("/api/post/index", &sys.PostController{}, "*:Index")
	beego.Router("/api/post/create", &sys.PostController{}, "post:Create")
	beego.Router("/api/post/update", &sys.PostController{}, "put:Update")
	beego.Router("/api/post/delete", &sys.PostController{}, "delete:Delete")

	//菜单管理
	beego.Router("/api/menu/index", &sys.MenuController{}, "*:Index")
	beego.Router("/api/menu/create", &sys.MenuController{}, "post:Create")
	beego.Router("/api/menu/update", &sys.MenuController{}, "put:Update")
	beego.Router("/api/menu/delete", &sys.MenuController{}, "delete:Delete")

	beego.Router("/api/menu/menus", &sys.MenuController{}, "*:Menus")
	beego.Router("/api/menu/find_menus", &sys.MenuController{}, "post:FindMenus")

	//字典管理
	beego.Router("/api/dict/index", &sys.DictTypeController{}, "*:Index")
	beego.Router("/api/dict/create", &sys.DictTypeController{}, "post:Create")
	beego.Router("/api/dict/update", &sys.DictTypeController{}, "put:Update")
	beego.Router("/api/dict/delete", &sys.DictTypeController{}, "delete:Delete")
	//字典数据管理
	beego.Router("/api/dictData/index", &sys.DictDataController{}, "*:Index")
	beego.Router("/api/dictData/create", &sys.DictDataController{}, "post:Create")
	beego.Router("/api/dictData/update", &sys.DictDataController{}, "put:Update")
	beego.Router("/api/dictData/delete", &sys.DictDataController{}, "delete:Delete")

	//部门管理
	beego.Router("/api/dept/findall", &sys.DeptController{}, "*:FindAll")
	beego.Router("/api/dept/create", &sys.DeptController{}, "post:Create")
	beego.Router("/api/dept/update", &sys.DeptController{}, "put:Update")
	beego.Router("/api/dept/delete", &sys.DeptController{}, "delete:Delete")


	//角色管理
	beego.Router("/api/role/index", &sys.RoleController{}, "*:Index")
	beego.Router("/api/role/create", &sys.RoleController{}, "post:Create")
	beego.Router("/api/role/update", &sys.RoleController{}, "put:Update")
	beego.Router("/api/role/delete", &sys.RoleController{}, "delete:Delete")

	//上传图片
	beego.Router("/api/upload/image", &commons.UploadController{}, "post:Image")

	//微信
	//beego.Router("/api/wechat/connect", &wx.WxConnectController{})
	beego.Router("/api/wechat/connect", &wx.WxConfigController{},"*:Get")
}
