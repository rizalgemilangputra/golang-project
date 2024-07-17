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
	Db     *gorm.DB
	Logger *logrus.Logger
	Viper  *viper.Viper
}

func NewApp(app *App) {
	modules := modules.Modules{
		App:    app.App,
		Db:     app.Db,
		Logger: app.Logger,
	}
	modules.Setup()
}
