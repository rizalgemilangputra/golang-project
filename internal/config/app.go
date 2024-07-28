package config

import (
	"golang-project/internal/modules"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type App struct {
	App    *gin.Engine
	DB     *gorm.DB
	Logger *logrus.Logger
	Viper  *viper.Viper
}

func (app *App) Create() {
	modules := modules.Modules{
		App:    app.App,
		DB:     app.DB,
		Logger: app.Logger,
	}
	modules.Setup()
}
