package util

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"go-cms/models"
	"strings"
	"time"
)

func CreateToken(user models.User) string {
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["id"]=user.Id
	claims["user_name"]=user.UserName
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	//token.Claims=claims
	tokenString,_ :=token.SignedString([]byte(beego.AppConfig.String("jwt::secrets")))
	
	return tokenString
}

func CheckToken(tokenString string) (b bool, t *jwt.Token) {
	kv := strings.Split(tokenString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		fmt.Println("AuthString invalid:", tokenString)
		return false, nil
	}
	
	token, err := jwt.Parse(kv[1], func(*jwt.Token) (interface{}, error) {
		return []byte(beego.AppConfig.String("jwt::secrets")), nil
	})
	
	fmt.Println(err)
	
	if err != nil {
		switch err.(type) {
		
		case *jwt.ValidationError: // something was wrong during the validation
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				//ctx.Output.SetStatus(401)
				//resBody, err := json.Marshal(controllers.OutResponse(401, nil, "登录已过期，请重新登录"))
				//ctx.Output.Body(resBody)
				//if err != nil {
				//	panic(err)
				//}
				return false,nil
			default:
				//ctx.Output.SetStatus(401)
				//resBytes, err := json.Marshal(controllers.OutResponse(401, nil, "非法请求，请重新登录"))
				//ctx.Output.Body(resBytes)
				//if err != nil {
				//	panic(err)
				//}
				return false,nil
			}
		default: // something else went wrong
			//ctx.Output.SetStatus(401)
			//resBytes, err := json.Marshal(controllers.OutResponse(401, nil, "非法请求，请重新登录"))
			//ctx.Output.Body(resBytes)
			//if err != nil {
			//	panic(err)
			//}
			return false,nil
		}
		
		fmt.Println("转换为jwt claims失败.", err)
		return false, nil
	}
	
	if !token.Valid {
		// but may still be invalid
		//ctx.Output.SetStatus(401)
		//resBytes, err := json.Marshal(controllers.OutResponse(401, nil, "非法请求，请重新登录"))
		//ctx.Output.Body(resBytes)
		//if err != nil {
		//	panic(err)
		//}
		return false, nil
	}
	//GetUserNameByToken(kv[1])
	return true, nil
}

func GetUserIdByToken(tokenString string)  int{
	token,_ :=jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _,ok :=token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,fmt.Errorf("Unexpected signing method")
		}
		return []byte(beego.AppConfig.String("jwt::secrets")),nil
	})
	claims,_:=token.Claims.(jwt.MapClaims)
	id := claims["id"].(int)
	return id
}

func GetUserNameByToken(tokenString string)  string{
	token,_ :=jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _,ok :=token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,fmt.Errorf("Unexpected signing method")
		}
		return []byte(beego.AppConfig.String("jwt::secrets")),nil
	})
	claims,_:=token.Claims.(jwt.MapClaims)
	user_name := claims["user_name"].(string)
	return user_name
}
