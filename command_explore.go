package main

import (
	"fmt"
)

func commandExplore(cfg *config,args ...string) error {
	if len(args)<1{
		return fmt.Errorf("explore requires a location name")
	}
	locationName := args[0]
	// fmt.Print(locationName)
	location,err:=cfg.pokeapiClient.GetLocationInfo(locationName)
	if err!=nil{
		return err
	}
	fmt.Printf("Exploring location %s\n",location.Name)
	fmt.Println("Found Pokemon:")
	for _,encounter:=range location.PokemonEncounters{
		fmt.Printf("- %s\n",encounter.Pokemon.Name)
	}
	return nil
}