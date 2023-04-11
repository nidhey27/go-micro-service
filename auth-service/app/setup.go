package app

import "github.com/nidhey27/auth-service/config"

func SetupAndRunApp() error {
	// load env
	err := config.LoadENV()
	if err != nil {
		return err
	}
	return nil

}
