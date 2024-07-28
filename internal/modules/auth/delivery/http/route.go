package http

import (
	"golang-project/internal/modules/auth/delivery/http/handler"

	"github.com/gin-gonic/gin"
)

func NewRoute(app *gin.Engine, handler *handler.Handler) *gin.Engine {
	app.POST("/register", handler.Register)
	app.POST("/login")
	return app
}
