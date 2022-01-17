package service

import (
	"fmt"
	"gotest/model"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	user := model.User{
		Name: "Manish",
		Id:   0,
	}
	fmt.Println(user)
}
