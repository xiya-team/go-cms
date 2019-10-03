package middlewares

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

//1.添加解决跨域请求问题
//2.文件下载文件夹，具体看beego官方文档
// 支持表单伪造PUT,DELETE,PATCH,OPTIONS请求
func CorsHandler() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
}
