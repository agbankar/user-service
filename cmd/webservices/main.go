package main

import (
	"fmt"
	"mocking-goway/dao"
	"mocking-goway/internal/appserver"
	"mocking-goway/internal/config"
	"mocking-goway/internal/model"
	"mocking-goway/internal/service"
	"net/http"
)

func main() {

	userDao := dao.NewUserDao()
	userService := service.NewUserService(userDao)

	user := model.User{
		Name: "Ashish",
		Id:   0,
	}
	ch := make(chan bool)
	bonus, _ := userService.CalculateBonus(&user)
	fmt.Println(bonus)
	c := &config.Config{Port: "8080"}
	server := appserver.NewAppServer(c)
	go server.StartServer(ch)
	<-ch

}

func GetData(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)

}
