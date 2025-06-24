package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto_user "github.com/mgambo/go-api/api/dto/user"
	"github.com/mgambo/go-api/api/services"
	utils "github.com/mgambo/go-api/api/utils"
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
// @Summary     Get Users
// @Description Perform get users
// @Accept      application/json
// @Produce     application/json
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

// @Tags		Users
// @Summary     Get Single User by id
// @Description Return user whose id is found
// @Param		id path string true "User ID"
// @Accept      application/json
// @Produce     application/json
// @Success		200 {object} dto_user.UserResponse
// @Router      /users/{id} [get]
func (controller *UsersController) GetSingleUser(c *gin.Context) {
	log.Info().Msg("Find By Id")
	id := c.Param("id")
	response, error := controller.userService.GetUserById(id)
	if error != nil {
		utils.HandleORMError(c, error)
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Tags		Users
// @Summary     Create User
// @Description Perform create user
// @Param		user body dto_user.CreateUserRequest true "Create User"
// @Produce     application/json
// @Success		200 {object} dto_user.UserResponse
// @Router      /users [post]
func (controller *UsersController) CreateUser(c *gin.Context) {
	log.Info().Msg("Create")
	createUserRequest := dto_user.CreateUserRequest{}
	err := c.ShouldBindJSON(&createUserRequest)
	if err != nil {
		log.Error().Msg(err.Error())
	}
	response, error := controller.userService.CreateUser(createUserRequest)
	if error != nil {
		log.Error().Msg(error.Error())
	}

	c.JSON(http.StatusOK, response)
}
