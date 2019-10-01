package services

import (
	"go-cms/models"
)

func FindByUserName(user_name string) models.User{
	userModel := models.NewUser()
	user, _ :=userModel.FindByUserName(user_name)
	return user
}