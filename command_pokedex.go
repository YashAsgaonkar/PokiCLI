package main

import (
	"fmt"
)

func commandPokedex(cfg *config,args ...string) error {
	if len(cfg.caughtPokemons) < 1 {
		fmt.Println("You haven't caught any pokemon yet")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for poki:=range(cfg.caughtPokemons){
		fmt.Printf("-%s\n", poki)
	}
	return nil
}