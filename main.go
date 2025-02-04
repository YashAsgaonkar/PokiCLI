package main

import (
	"time"

	"github.com/yashasgaonkar/pokedex/utils/pokeapi"
)
func main(){
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	REPL(cfg)
}

