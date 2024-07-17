package auth

import (
	"golang-project/internal/modules/auth/delivery/http"

	"github.com/gin-gonic/gin"
)

type AuthModule struct {
	Http *gin.Engine
}

func NewAuthModule(app *gin.Engine) *AuthModule {
	var http = http.NewRoute(app)

	return &AuthModule{
		Http: http,
	}
}
