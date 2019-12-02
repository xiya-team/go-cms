package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"go-cms/generate"
	"go-cms/middlewares"
	"go-cms/pkg/e"
	_ "go-cms/routers"
	"net/http"
	"runtime"
)

var Ctx *context.Context

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	data :=middlewares.Response{Code:e.ERROR,Msg:"404"}
	result,_:=json.Marshal(data)

	//ctx.Redirect(302, "/api/common/page_not_found")
	//t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/404.html")
	//data := make(map[string]interface{})
	//data["content"] = "page not found"
	//t.Execute(rw, data)
	rw.Write(result)
}

func init() {
	// 中间件注册
	middlewares.CorsHandler()
	beego.InsertFilter("*", beego.BeforeRouter, middlewares.RestfulHandler())
	//beego.InsertFilter("*", beego.BeforeRouter, middlewares.AuthMiddlewares())
}

func main() {
	//指定使用多核，核心数为CPU的实际核心数量
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	//gii
	if b, err := beego.AppConfig.Bool("gii"); b {
		fmt.Print("gii已经开启,代码生成中......")
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