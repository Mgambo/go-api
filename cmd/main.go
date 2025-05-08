package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mgambo/go-api/internal/routers"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /api/v1
func main() {
	r := routers.SetupRouter()

	r.Use(ErrorHandlerMiddleware())

	r.Use(gin.Recovery())
	// Listen and Server in 0.0.0.0:3000
	r.Run(":3000")
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
