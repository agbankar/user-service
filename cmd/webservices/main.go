package main

import (
	"fmt"
	"mocking-goway/cmd/webservices/controllers"
	"mocking-goway/internal/appserver"
	"mocking-goway/internal/config"
	"mocking-goway/internal/dao"
	"mocking-goway/internal/paths"
	"mocking-goway/internal/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	c := &config.Config{Port: "8080"}
	userService := service.NewUserService(dao.NewUserDao())
	userController := controllers.NewUserController(userService)
	server := appserver.NewAppServer(c)
	server.AddGetApi(paths.GetBonus, userController.GetBonus)
	go server.Start()

	//go appserver.Start()
	<-sigs
	fmt.Println("Done")
}
