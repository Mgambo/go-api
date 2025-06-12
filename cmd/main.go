package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mgambo/go-api/docs"
	"github.com/mgambo/go-api/internal/routers"
)

func init() {
	godotenv.Load()
}

// @title			Swagger API
// @version			1.0
// @description	Api Documentation.
// @host			localhost:3000
// @BasePath		/api/v1
func main() {
	r := routers.SetupRouter()

	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Title = os.Getenv("APP_NAME")
	if os.Getenv("HOST") != "" {
		docs.SwaggerInfo.Host = os.Getenv("HOST")
	} else {
		docs.SwaggerInfo.Host = "localhost:" + os.Getenv("PORT")
	}

	r.Use(ErrorHandlerMiddleware())

	r.Use(gin.Recovery())
	r.Run(":" + os.Getenv("PORT"))
}

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}
