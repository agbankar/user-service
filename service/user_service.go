package service

import (
	"fmt"
	"mocking-goway/dao"
	"mocking-goway/model"
)

type UserService interface {
	CalculateBonus(u *model.User) (float32, error)
}
type UserServiceImpl struct {
	UserDao dao.UserDao
}

func NewUserService(dao dao.UserDao) UserService {
	return &UserServiceImpl{UserDao: dao}
}
func (service *UserServiceImpl) CalculateBonus(u *model.User) (float32, error) {
	fmt.Println("Inside service  ", u)
	rating, err := service.UserDao.GetRating(u)
	if err != nil {
		return -1, err
	}
	if rating == "A" {
		return 100, nil
	}
	if rating == "B" {
		return 75, nil
	}
	if rating == "C" {
		return 50, nil
	}
	if rating == "D" {
		return 00, nil
	}
	return 33.00, nil
}
