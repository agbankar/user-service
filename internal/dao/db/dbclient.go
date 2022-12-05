package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"user-service/internal/config"
)

func NewDBClient(config *config.Config) (*gorm.DB, error) {

	db, err := gorm.Open("mysql", "root:test@(localhost:3306)/work")
	return db, err

}
