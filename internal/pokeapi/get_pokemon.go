package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName
	if pokemonName == "" {
		return Pokemon{}, errors.New("no pokemon name provided")
	}
	if val, ok := c.cache.Get(pokemonName); ok {
		pokemonData := Pokemon{}
		err := json.Unmarshal(val, &pokemonData)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonData, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonData := Pokemon{}
	err = json.Unmarshal(dat, &pokemonData)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(pokemonName, dat)
	return pokemonData, nil
}
