package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mgambo/go-api/internal/controllers"
)

var (
	healthController controllers.HealthController = controllers.NewHealthController()
)

func setupHealthRoute(server *gin.Engine, apiPath string) {
	api := server.Group(apiPath + "/health")

	// @Tags		Health
	// @Summary     Health Check
	// @Description Perform health check
	// @Tags        health
	// @Accept      json
	// @Produce     json
	// @Success     200 {object} map[string]interface{} "{"message": "ok"}"
	// @Router      /health [get]
	api.GET("/", healthController.HealthCheck)
}
