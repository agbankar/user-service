package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"user-service/cmd/user-services/controllers"
	"user-service/internal/appserver"
	"user-service/internal/config"
	"user-service/internal/dao"
	"user-service/internal/dao/db"
	"user-service/internal/paths"
	"user-service/internal/service"
	"user-service/internal/validations"

	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	c := getConf()
	db, err := db.NewMysqlClient(c)
	if err != nil {
		panic(err.Error())
	}
	validationService := validations.NewValidationService()
	userService := service.NewUserService(dao.NewUserDao(db))
	userController := controllers.NewUserController(userService, validationService)
	server := appserver.NewAppServer(c)
	server.AddGetApi(paths.GetBonus, userController.GetBonus)
	server.AddPostApi(paths.CreateUser, userController.CreateUser)
	server.AddGetApi(paths.GetUserOrders, userController.GetUserOrders)
	go server.Start()
	<-sigs
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	server.Stop(ctx)
	defer cancel()
	fmt.Println("Exiting from main")
	fmt.Println()
}

func getConf() *config.Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("%v", err)
	}

	conf := &config.Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return conf
}
