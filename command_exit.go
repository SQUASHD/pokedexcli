package main

import (
	"fmt"
	"os"
)

func commandExit(c *config, args ...string) error {
	exitString := "Exiting pokedex"
	fmt.Println(exitString)
	os.Exit(0)
	return nil
}
