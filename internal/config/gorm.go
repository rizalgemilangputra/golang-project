package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGorm(viper *viper.Viper, log *logrus.Logger) *gorm.DB {
	host := viper.GetString("database.host")
	database := viper.GetString("database.database")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	port := viper.GetString("database.port")
	sslMode := viper.GetString("database.sslMode")
	timeZone := viper.GetString("database.timeZone")

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, username, password, database, port, sslMode, timeZone)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	return db

}
