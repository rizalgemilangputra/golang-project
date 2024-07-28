package modules

import (
	"golang-project/internal/modules/auth"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Modules struct {
	App    *gin.Engine
	DB     *gorm.DB
	Logger *logrus.Logger
}

func (m *Modules) Setup() {
	auth.NewAuth(m.App, m.DB)
}
