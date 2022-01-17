package main

import (
	"fmt"
	"mocking-goway/dao"
	"mocking-goway/model"
	"mocking-goway/service"
)

func main() {
	userDao := dao.NewUserDao()
	userService := service.NewUserService(userDao)

	user := model.User{
		Name: "Ashish",
		Id:   0,
	}
	bonus, _ := userService.CalculateBonus(&user)
	fmt.Println(bonus)
}
