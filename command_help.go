package main

import "fmt"

func commandHelp(c *config, args ...string) error {
	helpString := "Help text"
	fmt.Println(helpString)
	for _, command := range getCommands() {
		fmt.Printf("%s - %s\n", command.name, command.description)
	}
	return nil
}
