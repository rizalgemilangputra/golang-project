package main

import (
	"golang-project/internal/config"
)

func main() {
	viper := config.NewViper()
	logger := config.NewLogger(viper)
	db := config.NewGorm(viper, logger)
	app := config.NewGin(viper)

	config := config.App{
		App:    app,
		DB:     db,
		Logger: logger,
		Viper:  viper,
	}
	config.Create()

	app.Run(":8080")
}
