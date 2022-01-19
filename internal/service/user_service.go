package service

import (
	"errors"
	"user-service/internal/dao"
	"user-service/internal/model"
)

type UserService interface {
	CalculateBonus(string) (float32, error)
	CreateUser(*model.User) error
}
type UserServiceImpl struct {
	UserDao dao.UserDao
}

func NewUserService(dao dao.UserDao) UserService {
	return &UserServiceImpl{UserDao: dao}
}
func (s *UserServiceImpl) CalculateBonus(UserId string) (float32, error) {
	rating, err := s.UserDao.GetRating(UserId)
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

// CreateUser TODO: Not implemented
func (s *UserServiceImpl) CreateUser(user *model.User) error {
	return nil

}
