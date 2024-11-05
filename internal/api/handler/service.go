package handler

import (
	"ddos/internal/models"
	"net/http"

	_ "ddos/docs"

	"github.com/gin-gonic/gin"
)

// Attack godoc
// @Summary Attack endpoint to initiate load test
// @Description This endpoint initiates a load test based on the given parameters.
// @Tags LoadTest
// @Accept  json
// @Produce  json
// @Param request body models.LoadTestRequest true "Load test parameters"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Router /attack [post]
func (h *Handler) Attack(ctx *gin.Context) {
	var req models.LoadTestRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Generator.Attack(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Load test initiated: "+res)
}
