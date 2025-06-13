package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController interface {
	HealthCheck(c *gin.Context)
}

type healthController struct{}

func NewHealthController() HealthController {
	return &healthController{}
}

// @Tags		Health
// @Summary     Health Check
// @Description Perform health check
// @Accept      json
// @Produce     json
// @Success     200 {object} map[string]interface{} "{"message": "ok"}"
// @Router      /health [get]
func (h *healthController) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
