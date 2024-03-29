package db

import (
	"fmt"
	"os"

	tweetRepository "github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/repositories/tweet"
	userRepository "github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/repositories/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseConnection() (*gorm.DB, error) {
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	// Migrate the schema
	//nolint: errcheck
	db.AutoMigrate(&userRepository.UserDBModel{}, &tweetRepository.DBModel{})

	return db, nil
}
