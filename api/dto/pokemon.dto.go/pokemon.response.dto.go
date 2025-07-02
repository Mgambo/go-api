package dto_pokemon

import "github.com/mgambo/go-api/api/models"

type PokemonResponse struct {
	Count    int              `json:"count"`
	Next     *string          `json:"next,omitempty"`
	Previous *string          `json:"previous,omitempty"`
	Results  []models.Pokemon `json:"results"`
}

type PokemonDataResponse struct {
	Data PokemonResponse `json:"data"`
}
