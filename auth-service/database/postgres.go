package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectToDB() *gorm.DB {
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")
	// dsn := os.Getenv("DB_URL")
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)

	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = d

	return db
}

func GetDB()
