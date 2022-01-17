package service

import (
	"fmt"
	"mocking-goway/model"
	"testing"
)

/**
UserDao dao.UserDao in userserive is to be mock
This is done by creating the userDaoMock struct which implements UserDao interface and assign this interface
to userService object .
           handlFunc is the func to return values as per mock requirement
*/
type userDaoMock struct {
	handlFunc func(User *model.User) (string, error)
}

func (daoMcok userDaoMock) GetRating(User *model.User) (string, error) {
	return daoMcok.handlFunc(User)
}
func TestRegisterUser(t *testing.T) {
	user := model.User{
		Name: "Manish",
		Id:   0,
	}
	mockA := userDaoMock{}
	mockA.handlFunc = func(User *model.User) (string, error) {
		return "A", nil
	}
	service := NewUserService(mockA)
	bonusA, _ := service.CalculateBonus(&user)
	fmt.Println("100% Bonus", bonusA)

	mockB := userDaoMock{}
	mockB.handlFunc = func(User *model.User) (string, error) {
		return "B", nil
	}
	service = NewUserService(mockB)
	bonusB, _ := service.CalculateBonus(&user)
	fmt.Println("75% Bonus", bonusB)
}
