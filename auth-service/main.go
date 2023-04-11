package main

import "github.com/nidhey27/auth-service/app"

func main() {
	// setup and run app
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
