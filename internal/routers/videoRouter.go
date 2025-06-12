package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mgambo/go-api/internal/controllers"
	"github.com/mgambo/go-api/internal/services"
)

var (
	videoService    services.VideoService       = services.New()
	videoController controllers.VideoController = controllers.NewVideoController(videoService)
)

func setupVideoRoute(server *gin.Engine, apiPath string) {
	api := server.Group(apiPath + "/video")

	//	@Tags			Video
	//	@Summary		get all videos
	//	@Description	get all the videos with name and content
	//	@Accept			json
	//	@Produce		json
	//	@Success		200	{string}	string				"ok"
	//	@Failure		400	{object}	handler.APIError	"We need ID!!"
	//	@Failure		404	{object}	handler.APIError	"Can not find ID"
	//	@Router			/video [get]
	api.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})

	//	@Tags			Video
	//	@Summary		create a new video
	//	@Description	create a new video with name and content
	//	@requestBody	{object} entity.Video
	//	@Accept			json
	//	@Produce		json
	//	@Success		200	{string}	entity.Video
	//	@Failure		400	{object}	handler.APIError	"We need ID!!"
	//	@Failure		404	{object}	handler.APIError	"Can not find ID"
	//	@Router			/video [post]
	api.POST("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.Save(ctx))
	})
}
