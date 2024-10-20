package router

import (
	"ddos/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func NewGin(h *handler.Handler) *gin.Engine {
	router := gin.Default()

	router.GET("/attack", h.Attack)

	return router
}
