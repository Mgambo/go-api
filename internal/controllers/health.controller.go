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

func (h *healthController) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
