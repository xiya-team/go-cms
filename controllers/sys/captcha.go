package sys

import (
	"encoding/json"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/tidwall/gjson"
	"go-cms/controllers"
	"go-cms/pkg/e"
)

type CaptchaController struct {
	controllers.BaseController
}

type Captcha struct {
	Ticket  string      `json:"ticket" form:"ticket"`
	Randstr   string      `json:"randstr"  form:"randstr"`
}

func (c *CaptchaController) Check()  {
	var captcha Captcha
	data := c.Ctx.Input.RequestBody
	//json数据封装到user对象中
	err := json.Unmarshal(data, &captcha)
	
	Ticket := captcha.Ticket
	Randstr := captcha.Randstr
	UserIp := c.Ctx.Input.IP()
	
	req := httplib.Get("https://captcha.tencentcloudapi.com/")
	req.Param("Action","DescribeCaptchaResult")
	req.Param("CaptchaType","9")
	req.Param("Ticket",Ticket)
	req.Param("UserIp",UserIp)
	req.Param("Version","2019-07-22")
	req.Param("Randstr",Randstr)
	req.Param("CaptchaAppId","1251180753")
	req.Param("AppSecretKey","zlqfnkcniyxxZvJQV2I2Xona69vQFpAE")
	str, err := req.String()
	
	if err != nil {
		c.JsonResult(e.ERROR, "error")
	}
	
	value := gjson.Get(str, "retcode")
	if value.Int() == 0 {
		c.JsonResult(e.SUCCESS, "success")
	}else {
		c.JsonResult(e.ERROR, str)
	}
}

func (c *CaptchaController) Hander(){
	var captcha Captcha
	data := c.Ctx.Input.RequestBody
	//json数据封装到user对象中
	err := json.Unmarshal(data, &captcha)
	
	Ticket := captcha.Ticket
	Randstr := captcha.Randstr
	UserIp := c.Ctx.Input.IP()
	
	req := httplib.Get("https://ssl.captcha.qq.com/ticket/verify")
	req.Param("Ticket",Ticket)
	req.Param("UserIP",UserIp)
	req.Param("Randstr",Randstr)
	req.Param("aid","2076088864")
	req.Param("AppSecretKey","06bEYSvZpRbeo6n_bMR0G_g**")
	str, err := req.String()
	
	if err != nil {
		c.JsonResult(e.ERROR, "error")
	}
	
	value := gjson.Get(str, "response")
	if value.Int() == 1 {
		c.JsonResult(e.SUCCESS, "success")
	}else {
		logs.Error(str)
		c.JsonResult(e.ERROR, "验证失败!")
	}
}