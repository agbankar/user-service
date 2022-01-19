package controllers

import (
	"fmt"
	"mocking-goway/internal/service"
	"net/http"
)

type UserController struct {
	UserService service.UserService
}

func GetBonus(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Go the request in controller")
	writer.Write([]byte("ok"))
	writer.WriteHeader(http.StatusOK)

}
