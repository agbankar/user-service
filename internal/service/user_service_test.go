package service

import (
	"errors"
	"gopkg.in/go-playground/assert.v1"
	"mocking-goway/internal/model"
	"testing"
)

var user model.User

func init() {
	user = model.User{
		Name: "Ashish",
		Id:   1,
	}

}

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
func TestBonusA(t *testing.T) {
	mockA := &userDaoMock{}
	mockA.handlFunc = func(User *model.User) (string, error) {
		return "A", nil
	}
	service := NewUserService(mockA)
	bonusA, _ := service.CalculateBonus(&user)
	assert.IsEqual(100, bonusA)

}

func TestBonusB(t *testing.T) {
	mockB := userDaoMock{}
	mockB.handlFunc = func(User *model.User) (string, error) {
		return "B", nil
	}
	service := NewUserService(mockB)
	bonusB, _ := service.CalculateBonus(&user)
	assert.IsEqual(75, bonusB)

}

func TestBonusC(t *testing.T) {
	mockC := userDaoMock{}
	mockC.handlFunc = func(User *model.User) (string, error) {
		return "C", nil
	}
	service := NewUserService(mockC)
	bonusC, _ := service.CalculateBonus(&user)
	assert.IsEqual(50, bonusC)

}
func TestBonusD(t *testing.T) {
	mockD := userDaoMock{}
	mockD.handlFunc = func(User *model.User) (string, error) {
		return "C", nil
	}
	service := NewUserService(mockD)
	bonusD, _ := service.CalculateBonus(&user)
	assert.IsEqual(0, bonusD)

}
func TestBonusInvalid(t *testing.T) {
	mockInvalid := userDaoMock{}
	mockInvalid.handlFunc = func(User *model.User) (string, error) {
		return "AA", nil
	}
	service := NewUserService(mockInvalid)
	bonusD, _ := service.CalculateBonus(&user)
	assert.IsEqual(-1, bonusD)

}
func TestBonusError(t *testing.T) {
	mockError := userDaoMock{}
	mockError.handlFunc = func(User *model.User) (string, error) {
		return "A", errors.New("Testing error")
	}
	service := NewUserService(mockError)
	errorBonus, _ := service.CalculateBonus(&user)
	assert.IsEqual(-1, errorBonus)

}
