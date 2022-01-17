package dao

import (
	"fmt"
	"gotest/model"
)

type UserDao interface {
	InsertUser(User *model.User)
}
type UserDaoImpl struct {
}

func NewUserDao() UserDao {
	return &UserDaoImpl{}

}
func (U *UserDaoImpl) InsertUser(User *model.User) {
	fmt.Println("Inside Dao", User)
}
