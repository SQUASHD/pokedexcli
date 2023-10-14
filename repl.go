package main

import (
	"bufio"
	"fmt"
	"github.com/squashd/pokedexcli/internal/pokeapi"
	"os"
	"strings"
)

type config struct {
	client               pokeapi.Client
	pokedex              map[string]pokeapi.Pokemon
	nextLocationsURL     *string
	previousLocationsURL *string
}

func printUnknownCommand(command string) {
	fmt.Println("Unknown command:", command)
}
func startRepl(c *config) {

	//configString := fmt.Sprintf("BaseURL: %s\nNext: %d\nPrevious: %d\nResultsNum: %d\n", c.BaseURL, c.Next, c.Previous, c.ResultsNum)
	//fmt.Print(configString)
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}

		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(c, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			printUnknownCommand(commandName)
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config, args ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location area you have visited by typing 'explore <location area name>'",
			callback:    commandExplore,
		},

		"catch": {
			name:        "catch",
			description: "Catch a pokemon by typing 'catch <pokemon name>'",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon you have caught by typing 'inspect <pokemon name>'",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays the names of all the pokemon you have caught",
			callback:    commandPokedex,
		},
	}
}
