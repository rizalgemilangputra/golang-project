package config

import (
	"golang-project/internal/delivery/http/route"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	App    *gin.Engine
	Db     *gorm.DB
	Logger *logrus.Logger
	Viper  *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {

	routeConfig := route.RouteConfig{
		App: config.App,
	}

	routeConfig.Setup()
}
