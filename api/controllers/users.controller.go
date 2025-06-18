package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mgambo/go-api/api/services"
	"github.com/rs/zerolog/log"
)

type UsersController struct {
	userService services.UserService
}

func NewUsersController(service services.UserService) *UsersController {
	return &UsersController{
		userService: service,
	}
}

// @Tags		Users
// @Summary     Users Check
// @Description Perform Users check
// @Accept      json
// @Produce     json
// @Success     200 {object} map[string]interface{} "{"message": "ok"}"
// @Router      /users [get]
func (controller *UsersController) GetUsers(c *gin.Context) {
	log.Info().Msg("FindAll")
	response, error := controller.userService.GetUsers()
	if error != nil {
		log.Error().Msg(error.Error())
	}
	c.JSON(http.StatusOK, response)
}
