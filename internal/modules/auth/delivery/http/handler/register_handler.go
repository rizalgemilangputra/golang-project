package handler

import (
	"golang-project/internal/modules/auth/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	var request model.RegisterMemberRequestModel

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.Usecase.Register(request)

	c.JSON(http.StatusOK, gin.H{"message": "Member created successfully", "user": request})
}
