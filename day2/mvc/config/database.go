package config

import (
	"fmt"
	"mvc/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var config = map[string]string{
	"user":     "root",
	"password": "",
	"host":     "127.0.0.1:3306",
	"dbname":   "agmc",
}

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config["user"],
		config["password"],
		config["host"],
		config["dbname"],
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	autoMigrate(db)

	return db
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
