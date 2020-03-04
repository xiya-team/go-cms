package middlewares

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/xiya-team/helpers"
	"go-cms/common"
	"go-cms/pkg/e"
	"go-cms/services"
)

func AuthMiddlewares() func(ctx *context.Context){
	var authMiddlewares = func(ctx *context.Context){
		user_id := common.UserId
		if !helpers.Empty(user_id){
			ctx.Output.Header("Content-Type", "application/json")

			userService := services.NewUserService()
			user_data := userService.FindByUserId(99999)
			logs.Debug(user_data.UserName)
			if helpers.Empty(user_data){

			}

			resBody, err := json.Marshal(OutResponse(e.ERROR, nil, "Method Not Allow"))
			if err!=nil{
				ctx.Output.Body(resBody)
			}
		}
	}
	return authMiddlewares
}


