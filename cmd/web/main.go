package main

import (
	"golang-project/internal/config"
)

func main() {
	viper := config.NewViper()
	logger := config.NewLogger(viper)
	db := config.NewGorm(viper, logger)
	app := config.NewGin(viper)

	config.Bootstrap(&config.BootstrapConfig{
		App:    app,
		Db:     db,
		Logger: logger,
		Viper:  viper,
	})

	app.Run(":8080")
}
