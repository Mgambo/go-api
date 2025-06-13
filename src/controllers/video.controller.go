package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mgambo/go-api/src/entity"
	"github.com/mgambo/go-api/src/services"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service services.VideoService
}

func NewVideoController(service services.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

// @Tags			Video
// @Summary		get all videos
// @Description	get all the videos with name and content
// @Accept			json
// @Produce		json
// @Success		200	{string}	string				"ok"
// @Failure		400	{object}	handler.APIError	"We need ID!!"
// @Failure		404	{object}	handler.APIError	"Can not find ID"
// @Router			/video [get]
func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

// @Tags			Video
// @Summary		create a new video
// @Description	create a new video with name and content
// @requestBody	{object} entity.Video
// @Accept			json
// @Produce		json
// @Success		200	{string}	entity.Video
// @Failure		400	{object}	handler.APIError	"We need ID!!"
// @Failure		404	{object}	handler.APIError	"Can not find ID"
// @Router			/video [post]
func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
