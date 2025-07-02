package routers

import (
	"net/http"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// gin-swagger middleware
// swagger embed files

var db = make(map[string]string)

func setupApiRoute(server *gin.Engine, apiPath string) {
	setupHealthRoute(server, apiPath)
	setupUserRoute(server, apiPath)
	setupPokemonRoute(server, apiPath)
	// setupVideoRoute(server, apiPath)
}

func setupAuthRoute(server *gin.Engine) {
	authorized := server.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:3000/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})
}

func SetupRouter() *gin.Engine {
	apiPath := "/api/v1"
	// Disable Console Color
	// gin.DisableConsoleColor()
	server := gin.Default()
	setupApiRoute(server, apiPath)
	setupAuthRoute(server)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return server
}
