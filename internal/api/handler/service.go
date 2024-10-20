package handler

import (
	"ddos/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Attack(ctx *gin.Context) {
	var req models.LoadTestRequest

	h.Generator.Attack(req)
}
