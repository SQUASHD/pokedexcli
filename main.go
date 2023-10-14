package main

import (
	"github.com/squashd/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeclient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	c := &config{
		client:  pokeclient,
		pokedex: map[string]pokeapi.Pokemon{},
	}
	startRepl(c)
}
