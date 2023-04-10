package dao

import (
	"fmt"
	"gorm.io/gorm"
	"user-service/internal/model"
)

type UserDao interface {
	GetRating(string) (string, error)
	CreateUser(*model.User) error
}
type UserDaoImpl struct {
	Db *gorm.DB
}

func NewUserDao(Db *gorm.DB) UserDao {
	return &UserDaoImpl{Db: Db}

}
func (d *UserDaoImpl) GetRating(UserId string) (string, error) {
	fmt.Println("Inside Dao", UserId)
	return "A", nil
}
func (d *UserDaoImpl) CreateUser(user *model.User) error {
	fmt.Println("inserting user in db", user)
	tx := d.Db.Create(&user)
	return tx.Error

}
