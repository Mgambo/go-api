package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPokemons(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pokemon"})
}
