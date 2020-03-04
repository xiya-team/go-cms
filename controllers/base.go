package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/xiya-team/helpers"
	"go-cms/common"
	"go-cms/models"
	"go-cms/pkg/d"
	"go-cms/pkg/e"
	"time"
)

type BaseController struct {
	beego.Controller
	StartTime        int64
	HandlerSeconds   float64
}

func init() {
	// 中间件注册
}

func (c *BaseController) Prepare() {
	// 启动时间
	c.StartTime = time.Now().UnixNano()

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

func (c *BaseController) Finish() {
	handlerSecond := float64(time.Now().UnixNano()-c.StartTime) / float64(1e9)
	c.HandlerSeconds = handlerSecond
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

	//记录操作日志
	is_log_record, _ := beego.AppConfig.Bool("is_log_record")
	if  is_log_record {
		switch code{
		case 0:
			c.InfoLog(msg)
		case 1:
			c.ErrorLog(msg)
		default:
			c.RecordLog(msg,code)
		}
	}

	c.ServeJSON()
	c.StopRun()
}


//获取当前url
func (c *BaseController) CurrentUrl() string {
	return helpers.Strtolower(c.Ctx.Request.URL.String())
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

// insert action log
func (c *BaseController) RecordLog(message string, level int) {
	userAgent := c.Ctx.Request.UserAgent()
	referer := c.Ctx.Request.Referer()
	getParams := c.Ctx.Request.URL.String()
	path := c.Ctx.Request.URL.Path
	postParamsMap := map[string][]string(c.Ctx.Request.PostForm)

	var postParams []byte
	if helpers.Empty(postParamsMap) {
		postParams = c.Ctx.Input.RequestBody
	}else {
		postParams, _ = json.Marshal(postParamsMap)
	}

	var model = models.LogInfo{
		Level:      level,
		Path:       path,
		Get:        getParams,
		Post:       string(postParams),
		Message:    message,
		Ip:         c.Ctx.Input.IP(),
		UserAgent:  userAgent,
		Referer:    referer,
		CreatedBy:  common.UserId,
		Method:  	c.Ctx.Request.Method,
		UpdatedBy:  0,
		Status:     0,
		Username:   common.UserName,
		CreateTime: time.Now(),
	}

	model.Create()
}

func (c *BaseController) ErrorLog(message string) {
	c.RecordLog(message, models.Log_Level_Error)
}

func (c *BaseController) WarningLog(message string) {
	c.RecordLog(message, models.Log_Level_Warning)
}

func (c *BaseController) InfoLog(message string) {
	c.RecordLog(message, models.Log_Level_Info)
}

func (c *BaseController) DebugLog(message string) {
	c.RecordLog(message, models.Log_Level_Debug)
}