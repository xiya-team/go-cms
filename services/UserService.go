package services

import (
	"go-cms/models"
)

type UserService struct {
	BaseService
}

func NewUserService() (userService *UserService) {
	return &UserService{}
}

func (c *UserService) FindByUserName(user_name string) models.User{
	userModel := models.NewUser()
	user, _ :=userModel.FindByUserName(user_name)
	return user
}

func (s *UserService) FindByUserId(id int) (user models.User){
	userModel := models.NewUser()
	user_data,_ :=userModel.FindById(id)
	return user_data
}