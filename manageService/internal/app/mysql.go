package app

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"manageService/internal/config"
)

var DB *gorm.DB

func InitMysql() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.PassWord, config.Host, config.Port, config.DB)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return err
}



