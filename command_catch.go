package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
    fmt.Printf("Throwing a pokeball at %s\n", args[0])
	pokiInfo, err := cfg.pokeapiClient.GetPokemonInfo(args[0])
	if err!=nil {
		fmt.Printf("No pokemon with that name found\n")
		return err
	}
	exp:=pokiInfo.BaseExperience
	if rand.Intn(exp+1)<60{    //catch logic
		cfg.caughtPokemons[pokiInfo.Name] = pokiInfo
		fmt.Printf("You caught %s\n",pokiInfo.Name)
	}else{
		fmt.Printf("%s broke free\n",pokiInfo.Name)
	}
	return nil
}