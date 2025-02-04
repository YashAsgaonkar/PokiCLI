package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/yashasgaonkar/pokedex/utils/pokeapi"
)
type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemons   map[string]pokeapi.Pokemon
}
type cliCommand struct {
	name        string
	description string
	callback    func(*config,...string) error //standard type for REPL CLI's print stuff or else return an error
}
func cleanInput(text string) []string{
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
    return strings.Fields(text)
}
func REPL(cfg *config)  {
	cfg.caughtPokemons = make(map[string]pokeapi.Pokemon)
	fmt.Println("Welcome to the Pokedex!")
    fmt.Println("------------------------")
	fmt.Println("Type 'help' to see the list of commands")
	reader :=bufio.NewReader(os.Stdin) //blocking code
	for{
		fmt.Print("Pokedex >")
		input,_:=reader.ReadString('\n')
		words :=cleanInput(input)
		if len(words)>0{
			command:=words[0]
			args := []string{}
			if len(words) > 1 {
			args = words[1:]
			}
			cmd, exist := getCommands()[command]
			if exist {
				err := cmd.callback(cfg,args...)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Invalid operation")
			}
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Catch a Pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore <location-name>",
			description: "Explore the Pokedex",
			callback:    commandExplore,
		},
		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "Inspect a Pokemon",
			callback:    commandInspect,
		},
		"pokedex":{
			name:        "pokedex",
			description: "List all caught pokemons",
			callback:    commandPokedex,
		},
	}
}
