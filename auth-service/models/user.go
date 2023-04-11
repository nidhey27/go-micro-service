package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName     string `json:"user_name" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Role         int    `json:"role" binding:"required"`
	RefreshToken string `json:"refresh_token"`
}
