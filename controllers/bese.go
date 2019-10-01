package controllers

import (
	"github.com/astaxie/beego"
	"github.com/syyongx/php2go"
	"go-cms/common"
	"go-cms/models"
	"go-cms/pkg/d"
	"go-cms/pkg/e"
	"go-cms/pkg/util"
)

type BaseController struct {
	beego.Controller
	ADMIN_TPL string
}


func (c *BaseController) Prepare() {
	c.ADMIN_TPL = "admin/"

	common.Fc = c.Ctx
	if user := c.GetSession("loginUser"); user != nil {
		common.UserId = user.(*models.User).Id
	}

	/*	controller, action := c.GetControllerAndAction()
		if controller!="UserController" && c.GetSession("loginUser") == nil{
			c.History("未登录","/login")
		}

		if controller == "UserController" && action == "Login" && c.GetSession("loginUser") != nil {
			c.History("已登录", "/admin")
		}*/
}

func (c *BaseController) History(msg string, url string) {
	if url == "" {
		c.Ctx.WriteString("<script>alert('" + msg + "');window.history.go(-1);</script>")
		c.StopRun()
	} else {
		c.Redirect(url, 302)
	}
}

func (c *BaseController) JsonResult(code int, msg string, data ...interface{}) {
	if len(data) > 1 {
		c.Data["json"] = d.LayuiJson(code, msg, data[0], data[1])
	} else if len(data) > 0 {
		c.Data["json"] = d.LayuiJson(code, msg, data[0], 0)
	} else {
		c.Data["json"] = d.LayuiJson(code, msg, 0, 0)
	}
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) CheckToken(){
	
	token := c.Ctx.Input.Header("Authorization")
	
	b, _ := util.CheckToken(token)
	
	if !b {
		c.JsonResult(e.ERROR,"验证失败!")
	}
	
	c.JsonResult(e.SUCCESS,"success")
}

//获取当前url
func (c *BaseController) CurrentUrl() string {
	return php2go.Strtolower(c.Ctx.Request.URL.String())
}