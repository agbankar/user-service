package main

import (
	"gotest/dao"
	"gotest/model"
	"gotest/service"
)

func main() {

	userDao := dao.NewUserDao()
	userService := service.NewUserService(userDao)

	user := model.User{
		Name: "Ashish",
		Id:   0,
	}
	userService.RegisterUser(&user)
}
