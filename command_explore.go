package main

import (
	"errors"
	"fmt"
)

func commandExplore(c *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("no city name provided")
	}
	cityName := args[0]
	location, err := c.client.ExploreLocation(cityName)
	if err != nil {
		return err
	}
	for _, pokemon := range location.PokemonEncounters {
		fmt.Printf("%s\n", pokemon.Pokemon.Name)
	}
	return nil
}
