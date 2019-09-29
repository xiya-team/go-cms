package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
	"go-cms/fc"
	"go-cms/handlers"
	_ "go-cms/routers"
	"html/template"
	"net/http"
	
	"github.com/astaxie/beego"
)

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/404.html")
	data := make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}

func init() {
	//gii
	if b, err := beego.AppConfig.Bool("gii"); b {
		if err == nil {
			fc.Run() //开启gii
		}
		return
	}
	
	//1.添加解决跨域请求问题
	//2.文件下载文件夹，具体看beego官方文档
	// 支持表单伪造PUT,DELETE,PATCH,OPTIONS请求
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	
	beego.InsertFilter("*", beego.BeforeRouter, handlers.RestfulHandler())
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	
	//log
	logs.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	//输出文件名，行号
	logs.EnableFuncCallDepth(true)
	//异步log
	logs.Async(1e3)
	//404
	beego.ErrorHandler("404", page_not_found)
	
	beego.Run()
}