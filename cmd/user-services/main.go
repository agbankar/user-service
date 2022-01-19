package main

import (
	"context"
	"fmt"
	"user-service/cmd/user-services/controllers"
	"user-service/internal/appserver"
	"user-service/internal/config"
	"user-service/internal/dao"
	"user-service/internal/paths"
	"user-service/internal/service"

	"os"
	"os/signal"
	"syscall"
	"time"
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
	<-sigs
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	server.Stop(ctx)
	defer cancel()
	fmt.Println("Exiting from main")
}
