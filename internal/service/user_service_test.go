package service

import (
	"errors"
	"fmt"
	"gopkg.in/go-playground/assert.v1"
	"testing"
	"user-service/internal/model"
)

func init() {

	fmt.Println("Initialing tests")
}

//UserDao dao.UserDao in userserive is to be mocked
//This is done by creating the userDaoMock struct which implements UserDao interface and assign this interface
//to userService object. handlFunc is the func to return values as per mock requirement

type userDaoMock struct {
	handlFunc func(UserID string) (string, error)
}

func (daoMcok *userDaoMock) GetUserOrders(UserId int) (model.User, error) {
	return model.User{}, nil

	//TODO implement me
}

func (daoMcok *userDaoMock) GetRating(UserID string) (string, error) {
	return daoMcok.handlFunc(UserID)
}
func (daoMcok *userDaoMock) CreateUser(User *model.User) error {
	fmt.Println("Creating user", User)
	return nil

}
func TestBonusA(t *testing.T) {
	mockDao := &userDaoMock{}
	mockDao.handlFunc = func(UserID string) (string, error) {
		return "A", nil

	}
	service := NewUserService(mockDao)
	bonus, _ := service.CalculateBonus("456")
	assert.Equal(t, bonus, float32(100))
}
func TestBonusB(t *testing.T) {
	mockDao := &userDaoMock{}
	mockDao.handlFunc = func(UserID string) (string, error) {
		return "B", nil

	}
	service := NewUserService(mockDao)
	bonus, _ := service.CalculateBonus("456")
	assert.Equal(t, bonus, float32(75))

}
func TestBonusC(t *testing.T) {
	mockDao := &userDaoMock{}
	mockDao.handlFunc = func(UserID string) (string, error) {
		return "C", nil
	}
	service := NewUserService(mockDao)
	bonus, _ := service.CalculateBonus("456")
	assert.Equal(t, bonus, float32(50))

}
func TestBonusD(t *testing.T) {
	mockDao := &userDaoMock{}
	mockDao.handlFunc = func(UserID string) (string, error) {
		return "D", nil
	}
	service := NewUserService(mockDao)
	bonus, _ := service.CalculateBonus("456")
	assert.Equal(t, bonus, float32(0))

}
func TestBonusInvalid(t *testing.T) {
	mockDao := &userDaoMock{}
	mockDao.handlFunc = func(UserID string) (string, error) {
		return "AAA", nil
	}
	service := NewUserService(mockDao)
	bonus, _ := service.CalculateBonus("456")
	assert.Equal(t, bonus, float32(-1))

}
func TestBonusError(t *testing.T) {
	mockDao := &userDaoMock{}
	mockDao.handlFunc = func(UserID string) (string, error) {
		return "A", errors.New("error")
	}
	service := NewUserService(mockDao)
	bonus, _ := service.CalculateBonus("456")
	assert.Equal(t, bonus, float32(-1))
	fmt.Println()

}
