package main

import (
	"os"
	"atykym.space/app"
	"atykym.space/app/config"
)

func main() {
	configPath, isPresent := os.LookupEnv("CONFIG_PATH")

	if !isPresent {
		panic("CONFIG_PATH env must be present")
	}

	appConfig := config.ReadConfig(configPath)
	var app app.App = app.App{Config: appConfig}
	app.Start()

}
