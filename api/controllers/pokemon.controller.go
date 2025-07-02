package controllers

import (
	"fmt"
	"io"
	"net/http"

	"encoding/json"

	"github.com/gin-gonic/gin"
	dto_pokemon "github.com/mgambo/go-api/api/dto/pokemon.dto.go"
	"github.com/mgambo/go-api/api/models"
)

type PokemonController interface {
	GetAllPokemon(c *gin.Context)
}

type pokemonController struct{}

func NewPokemonController() PokemonController {
	return &pokemonController{}
}

// @Tags		Pokemon
// @Summary     Get all the pokemon
// @Description Get all the pokemon with limit 100
// @Accept      json
// @Produce     json
// @Success     200 {object} dto_pokemon.PokemonDataResponse
// @Router      /pokemon [get]
func (h *pokemonController) GetAllPokemon(c *gin.Context) {
	result, error := http.Get("https://pokeapi.co/api/v2/pokemon?limit=100")

	if (error != nil) || (result.StatusCode != http.StatusOK) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch pokemon data",
		})
		return
	}
	defer result.Body.Close() // Close the response body when done

	body, err := io.ReadAll(result.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read response body",
		})
		return
	}

	var data dto_pokemon.PokemonResponse
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error unmarshalling response body:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to parse pokemon data",
		})
		return
	}

	filterPokemon := models.Pokemon{}
	for pokemon := range data.Results {
		// Convert URL to a string
		data.Results[pokemon].URL = fmt.Sprintf("prefix-%s", data.Results[pokemon].URL)
		if data.Results[pokemon].Name == "pikachu" {
			filterPokemon = data.Results[pokemon]
		}
	}
	data.Results = []models.Pokemon{filterPokemon}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
