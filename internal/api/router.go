package router

import (
	"ddos/internal/api/handler"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Voting service
// @version 1.0
// @description Delivery service
// @BasePath /
// @in header
// @name Authorization
func NewGin(h *handler.Handler) *gin.Engine {
	router := gin.Default()

	router.GET("/attack", h.Attack)

	swaggerURL := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, swaggerURL))

	return router
}
