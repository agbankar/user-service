package dao

import (
	"fmt"
	"user-service/internal/model"
)

type UserDao interface {
	GetRating(string) (string, error)
	CreateUser(*model.User) error
}
type UserDaoImpl struct {
}

func NewUserDao() UserDao {
	return &UserDaoImpl{}

}
func (d *UserDaoImpl) GetRating(UserId string) (string, error) {
	fmt.Println("Inside Dao", UserId)
	return "A", nil
}
func (d *UserDaoImpl) CreateUser(user *model.User) error {
	fmt.Println("inserting user in db")
	return nil

}
