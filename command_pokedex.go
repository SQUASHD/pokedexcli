package main

import (
	"fmt"
)

func commandPokedex(c *config, args ...string) error {
	if len(c.pokedex) == 0 {
		fmt.Println("You haven't caught any pokemon yet")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range c.pokedex {
		fmt.Printf(" – %s\n", pokemon.Name)
	}
	return nil
}
