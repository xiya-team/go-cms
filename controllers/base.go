package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/syyongx/php2go"
	"go-cms/middlewares"
	"go-cms/pkg/d"
	"go-cms/pkg/e"
	"strings"
)

type BaseController struct {
	beego.Controller
}

//配置不需要登录的url
var urlMapping = []string{"user::login","captcha::check","wechat::connect"}

func (c *BaseController) Prepare() {

	current_url := c.Ctx.Request.URL.RequestURI()
	controllerName, actionName := getControllerAndAction(current_url)
	is_pass := php2go.InArray(php2go.Strtolower(controllerName+"::"+actionName),urlMapping)

	if(is_pass){
		beego.InsertFilter("*", beego.BeforeRouter, middlewares.RestfulHandler())
	}
	
	//if user := c.GetSession("loginUser"); user != nil {
	//	UserId = user.(*models.User).Id
	//}

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
	
	switch len(data) {
	case 4:
		c.Data["json"] = d.LayuiJson(code, msg, data[0], data[1],data[2],data[3])
	case 3:
		c.Data["json"] = d.LayuiJson(code, msg, data[0], data[1],data[2],false)
	case 2:
		c.Data["json"] = d.LayuiJson(code, msg, data[0], data[1],false,false)
	case 1:
		c.Data["json"] = d.LayuiJson(code, msg, data[0], false,false,false)
	default:
		c.Data["json"] = d.LayuiJson(code, msg, false, false,false,false)
	}
	c.ServeJSON()
	c.StopRun()
}


//获取当前url
func (c *BaseController) CurrentUrl() string {
	return php2go.Strtolower(c.Ctx.Request.URL.String())
}

// 自动化的表单验证器
func (c *BaseController) ValidatorAuto(frontendData interface{}) {
	
	defaultMessage := map[string]string{
		"Required":     "不能为空",
		"Min":          "不能小于%d",
		"Max":          "不能大于%d",
		"Range":        "取值必须在%d到%d之间",
		"MinSize":      "长度不能小于%d",
		"MaxSize":      "长度不能大于%d",
		"Length":       "长度必须等于%d",
		"Alpha":        "必须是字母",
		"Numeric":      "必须是数字",
		"AlphaNumeric": "必须是字母或者数字",
		"Match":        "必须出现 %s 关键字",
		"NoMatch":      "不能出现 %s 关键字",
		"AlphaDash":    "必须是字母，数组或者横线(-)",
		"Email":        "不合法的邮箱地址",
		"IP":           "不合法的IP",
		"Base64":       "不合法的Base64编码格式",
		"Mobile":       "不合法的手机号",
		"Tel":          "不合法的电话号码",
		"Phone":        "不合法的手机号",
		"ZipCode":      "不合法的邮编",
	}
	validation.SetDefaultMessage(defaultMessage)
	
	validate := validation.Validation{}
	
	isValid, err := validate.Valid(frontendData)
	if err != nil {
		c.JsonResult(e.ERROR,"数据有问题!")
	}
	
	if !isValid {
		for _, err := range validate.Errors {
			c.JsonResult(e.ERROR, err.Message)
			//c.JsonResult(e.ERROR, err.Key+":"+err.Message)
		}
	}
}

// 重定向
func (c *BaseController) RedirectTo(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}

func getControllerAndAction(url string)  (controllerName, actionName string){
	newStr := strings.ReplaceAll(strings.TrimLeft(url,"/api"),"/","|")

	tmp :=strings.Split(newStr, "|")
	var tow = ""
	if len(tmp) >2 {
		tow = tmp[1]
	}
	return tmp[0],tow
}