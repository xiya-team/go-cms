package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go-cms/generate"
	"go-cms/middlewares"
	_ "go-cms/routers"
	"html/template"
	"net/http"
)

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/404.html")
	data := make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}

func init() {
	// 中间件注册
	middlewares.CorsHandler()
	
	beego.InsertFilter("*", beego.BeforeRouter, middlewares.RestfulHandler())
}

func main() {
	
	//gii
	if b, err := beego.AppConfig.Bool("gii"); b {
		if err == nil {
			generate.Run() //开启gii
		}
		return
	}
	
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