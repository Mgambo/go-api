package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController interface {
	FindAll(c *gin.Context)
}

type usersController struct{}

func NewUsersController() UsersController {
	return &usersController{}
}

// @Tags		Users
// @Summary     Users Check
// @Description Perform Users check
// @Accept      json
// @Produce     json
// @Success     200 {object} map[string]interface{} "{"message": "ok"}"
// @Router      /Users [get]
func (h *usersController) FindAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
