package service

import (
	"errors"
	"user-service/internal/dao"
	"user-service/internal/model"
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
	rating, err := service.UserDao.GetRating(u)
	if err != nil {
		return -1, errors.New("Not able to fetch the data from DB")
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
		return 0, nil
	}
	return -1, errors.New("Not able to fetch the data from DB")
}
