package auth

import (
	"golang-project/internal/modules/auth/delivery/http"
	"golang-project/internal/modules/auth/delivery/http/handler"
	"golang-project/internal/modules/auth/repository"
	"golang-project/internal/modules/auth/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Auth struct {
	App        *gin.Engine
	Handler    *handler.Handler
	Repository *repository.Repository
}

func NewAuth(app *gin.Engine, db *gorm.DB) *Auth {
	var (
		repository = repository.NewRepository(db)
		usecase    = usecase.NewUsecase(repository)
		handler    = handler.NewHandler(app, usecase)
		route      = http.NewRoute(app, handler)
	)

	return &Auth{
		App:        route,
		Handler:    handler,
		Repository: repository,
	}
}
