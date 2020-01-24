package main

import (
	"github.com/famous-persons-rest-api/app"
	"github.com/famous-persons-rest-api/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
