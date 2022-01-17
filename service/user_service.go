package service

import (
	"fmt"
	"gotest/dao"
	"gotest/model"
)

type UserService interface {
	RegisterUser(u *model.User)
}
type UserServiceImpl struct {
	UserDao dao.UserDao
}

func NewUserService(dao dao.UserDao) UserService {
	return &UserServiceImpl{UserDao: dao}
}
func (service *UserServiceImpl) RegisterUser(u *model.User) {
	fmt.Println("Inside service  ", u)
	service.UserDao.InsertUser(u)
}
