package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mgambo/go-api/api/controllers"
	"github.com/mgambo/go-api/api/repositories"
	"github.com/mgambo/go-api/api/services"
	"github.com/mgambo/go-api/internal/database"
)

func setupUserRoute(server *gin.Engine, apiPath string) {
	// validate
	validate := validator.New()

	// Repository
	userRepository := repositories.NewUserRepository(database.Db)

	// Service
	userService := services.NewUserServiceImpl(userRepository, validate)
	usersController := controllers.NewUsersController(userService)

	api := server.Group(apiPath + "/users")

	api.GET("/", usersController.GetUsers)
	api.GET("/:id", usersController.GetSingleUser)
	api.POST("/", usersController.CreateUser)
}
