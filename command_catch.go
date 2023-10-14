package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(c *config, args ...string) error {

	if len(args) < 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s\n", pokemonName)
	pokemon, err := c.client.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	if res > 30 {
		fmt.Printf("You caught %s!\n", pokemonName)
		c.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("Oh no! %s escaped!\n", pokemonName)
	}

	return nil
}
