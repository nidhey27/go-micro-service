package database

import "github.com/nidhey27/auth-service/models"

func SyncDatabase() {
	GetDB().AutoMigrate(&models.User{})
}
