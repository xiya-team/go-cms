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
	tokenString,_ :=token.SignedString([]byte(beego.AppConfig.String("token_secrets")))
	
	return tokenString
}

func CheckToken(tokenString string) (b bool, t *jwt.Token) {
	kv := strings.Split(tokenString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		fmt.Println("AuthString invalid:", tokenString)
		return false, nil
	}
	
	t, err := jwt.Parse(kv[1], func(*jwt.Token) (interface{}, error) {
		return []byte(beego.AppConfig.String("token_secrets")), nil
	})
	
	fmt.Println(err)
	
	if err != nil {
		fmt.Println("转换为jwt claims失败.", err)
		return false, nil
	}
	return true, t
}

func GetUserIdByToken(tokenString string)  float64{
	token,_ :=jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _,ok :=token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,fmt.Errorf("Unexpected signing method")
		}
		return []byte(beego.AppConfig.String("token_secrets")),nil
	})
	claims,_:=token.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	return id
}

func GetUserNameByToken(tokenString string)  string{
	token,_ :=jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _,ok :=token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,fmt.Errorf("Unexpected signing method")
		}
		return []byte(beego.AppConfig.String("token_secrets")),nil
	})
	claims,_:=token.Claims.(jwt.MapClaims)
	user_name := claims["user_name"].(string)
	return user_name
}
