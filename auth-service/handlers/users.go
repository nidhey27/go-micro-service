package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nidhey27/auth-service/database"
	"github.com/nidhey27/auth-service/models"
	"golang.org/x/crypto/bcrypt"
)

type requestPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(c *gin.Context) {
	var body requestPayload

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "",
			"data":    make([]string, 0),
			"error":   "Failed to read REQUEST Body",
		})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "",
			"error":   "Failed to Hash Password",
			"data":    make([]string, 0),
		})
		return
	}

	// Create the user
	user := models.User{Email: body.Email, Password: string(passwordHash)}

	result := database.GetDB().Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "",
			"error":   result.Error,
			"data":    make([]string, 0),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "User Created!",
		"error":   "",
		"data":    user,
	})
}
