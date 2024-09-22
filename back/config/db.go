package config

import (
	"back/models"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// connectWithRetry tries to connect to the database with a specified number of retries
func connectWithRetry(dsn string, maxAttempts int) *gorm.DB {
	var db *gorm.DB
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("Database connection successful!")
			return db
		}
		fmt.Printf("Failed to connect to database. Attempt %d/%d. Error: %s\n", attempts, maxAttempts, err)
		time.Sleep(time.Duration(attempts) * time.Second) // Increase wait time with each attempt
	}
	log.Fatalf("Failed to connect to the database after %d attempts.", maxAttempts)
	return nil
}

// ConnectDB initializes the DB connection using the connectWithRetry function
func ConnectDB() {
	dsn := "root:lozinka123@tcp(db:3306)/posts?charset=utf8mb4&parseTime=True&loc=Local"
	DB = connectWithRetry(dsn, 5)
	DB.AutoMigrate(&models.User{}, &models.Post{})
}
