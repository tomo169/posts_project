package config

import (
	"fmt"

	"back/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDB() {
	dsn := "root:lozinka123@tcp(127.0.0.1:3306)/posts?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Faild to connect to the database:", err)
		return
	}
	fmt.Println("Database connection successful!")

	DB.AutoMigrate(&models.User{}, &models.Post{})
}
