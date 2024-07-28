package handler

import (
	"golang-project/internal/modules/auth/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	App     *gin.Engine
	Usecase *usecase.Usecase
}

func NewHandler(app *gin.Engine, usecase *usecase.Usecase) *Handler {
	handler := &Handler{
		App:     app,
		Usecase: usecase,
	}

	return handler
}
