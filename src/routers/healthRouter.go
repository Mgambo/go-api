package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mgambo/go-api/src/controllers"
)

var (
	healthController controllers.HealthController = controllers.NewHealthController()
)

func setupHealthRoute(server *gin.Engine, apiPath string) {
	api := server.Group(apiPath + "/health")

	api.GET("/", healthController.HealthCheck)
}
