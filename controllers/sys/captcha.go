package sys

import (
	"gitea/modules/httplib"
	"github.com/tidwall/gjson"
	"go-cms/controllers"
	"go-cms/pkg/e"
)

type CaptchaController struct {
	controllers.BaseController
}

func (c *CaptchaController) Check()  {
	Ticket := c.GetString("ticket")
	UserIp := c.Ctx.Input.IP()
	Randstr := c.GetString("randstr")
	
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
	if value.String() == "0" {
		c.JsonResult(e.SUCCESS, "success")
	}else {
		c.JsonResult(e.ERROR, str)
	}
}