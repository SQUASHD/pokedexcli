package main

import (
	"errors"
	"fmt"
)

func commandInspect(c *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]
	val, ok := c.pokedex[pokemonName]
	if ok {
		res, err := c.client.GetPokemon(val.Name)
		if err != nil {
			return err
		}
		fmt.Printf("Name: %s\n", res.Name)
		fmt.Printf("Height: %d\n", res.Height)
		fmt.Printf("Weight: %d\n", res.Weight)
		fmt.Println("Stats:")
		for _, v := range res.Stats {
			fmt.Printf("  %s: %d\n", v.Stat.Name, v.BaseStat)
		}
		fmt.Println("Types:")
		for _, v := range res.Types {
			fmt.Println(v.Type.Name)
		}
	} else {
		fmt.Println("You haven't caught that pokemon")
	}
	return nil
}
