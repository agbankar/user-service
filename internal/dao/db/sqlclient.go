package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"user-service/internal/config"
	"user-service/internal/model"
)

func NewMysqlClient(Config *config.Config) (db *gorm.DB, err error) {
	connectionDSN := Config.Db.UserName + ":" + Config.Db.Password + "@tcp(" + Config.Db.Url + ":" + Config.Db.Port + ")" + "/" + Config.Db.Name

	db, err = gorm.Open(mysql.Open(connectionDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(Config.Db.LogLevel)),
	})
	if err != nil {
		fmt.Println(db)
		return nil, err
	}
	initDB(db)
	return db, err

}

// Initialize database
func initDB(db *gorm.DB) {
	db.AutoMigrate(model.User{})

}
