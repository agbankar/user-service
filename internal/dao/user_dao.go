package dao

import (
	"fmt"
	"user-service/internal/model"
)

type UserDao interface {
	GetRating(User *model.User) (string, error)
}
type UserDaoImpl struct {
}

func NewUserDao() UserDao {
	return &UserDaoImpl{}

}
func (U *UserDaoImpl) GetRating(User *model.User) (string, error) {
	fmt.Println("Inside Dao", User)
	return "A", nil
}
