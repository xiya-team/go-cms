package middlewares

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/syyongx/php2go"
	"go-cms/pkg/d"
	"go-cms/pkg/e"
	"go-cms/pkg/util"
)

var supportMethod = [6]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}

//配置不需要登录的url
var urlMapping = []string{"/api/user/login","/api/user/create"}

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
		
		allow := false
		current_url := ctx.Request.URL.RequestURI()
		for _, baseurl := range urlMapping{
			if php2go.Strtolower(baseurl) == php2go.Strtolower(current_url) {
				allow = true
				break
			}
		}
		
		//判断是否需要登录
		if allow == false{
			token := ctx.Input.Header(beego.AppConfig.String("jwt::token_name"))
			b, _ := util.CheckToken(token)
			if(b == false){
				Data := make(map[interface{}]interface{})
				Data["json"] = d.LayuiJson(e.ERROR, "需要登录", "", "")
				ctx.Output.JSON(Data["json"], false, false)
				panic(errors.New("user stop run"))
				php2go.Exit(0)
				return
			}
			
			//
			//_, ok := ctx.Input.Session("uid").(string)
			//ok2 := strings.Contains(ctx.Request.RequestURI, "/login")
			//if !ok && !ok2{
			//	ctx.Redirect(302, "/login/index")
			//}
		}
		
	}
	return restfulHandler
}