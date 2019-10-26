package commons

import (
	"go-cms/controllers"
	"go-cms/pkg/e"
)

// 全局验证码结构体
func init()  {
	// 验证码功能
}

type CommonController struct {
	controllers.BaseController
}

func (c *CommonController) PageNotFound() {
	c.JsonResult(e.ERROR, "404")
}