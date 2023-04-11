package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {

	routes := app.Group("/api/v1")
	routes.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version":     "v1",
			"author":      "Nidhey Indurkar",
			"description": "Auth Service",
		})
	})
}
