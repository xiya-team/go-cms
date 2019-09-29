package handlers

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

var supportMethod = [6]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}

// 支持伪造restful风格的http请求
// _method = "DELETE" 即将http的POST请求改为DELETE请求
func RestfulHandler() func(ctx *context.Context) {
	var restfulHandler = func(ctx *context.Context) {
		// 获取隐藏请求
		requestMethod := ctx.Input.Query("_method")
		
		if requestMethod ==  ""{
			// 正常请求
			requestMethod = ctx.Input.Method()
			logs.Debug("requestMethod:",requestMethod)
		}
		
		// 判断当前请求是否在允许请求内
		flag := false
		for _, method := range supportMethod{
			if method == requestMethod {
				flag = true
				break
			}
		}
		
		// 方法请求
		if flag == false {
			ctx.ResponseWriter.WriteHeader(405)
			ctx.Output.Body([]byte("Method Not Allow"))
			return
		}
		
		// 伪造请求方式
		if requestMethod != "" && ctx.Input.IsPost() {
			ctx.Request.Method = requestMethod
		}
	}
	return restfulHandler
}