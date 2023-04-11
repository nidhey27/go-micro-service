package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nidhey27/auth-service/config"
	"github.com/nidhey27/auth-service/database"
	"github.com/nidhey27/auth-service/routes"
)

func SetupAndRunApp() error {
	// load env
	err := config.LoadENV()
	if err != nil {
		return err
	}

	// start database
	err = database.ConnectToDB()
	if err != nil {
		return err
	}

	// defer closing database
	defer database.CloseDB()

	// create app
	app := gin.New()

	// setup routes
	routes.SetupRoutes(app)


	port := os.Getenv("PORT")
	app.Run(":" + port)

	return nil

}
