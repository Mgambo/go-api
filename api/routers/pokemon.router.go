package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mgambo/go-api/api/controllers"
)

var (
	pokemonController controllers.PokemonController = controllers.NewPokemonController()
)

func setupPokemonRoute(server *gin.Engine, apiPath string) {
	api := server.Group(apiPath + "/pokemon")

	api.GET("/", pokemonController.GetAllPokemon)
}
