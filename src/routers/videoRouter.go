package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mgambo/go-api/src/controllers"
	"github.com/mgambo/go-api/src/services"
)

var (
	videoService    services.VideoService       = services.New()
	videoController controllers.VideoController = controllers.NewVideoController(videoService)
)

func setupVideoRoute(server *gin.Engine, apiPath string) {
	api := server.Group(apiPath + "/video")

	api.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})

	api.POST("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.Save(ctx))
	})
}
