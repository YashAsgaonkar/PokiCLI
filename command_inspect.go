package main

import (
	"fmt"
)

func commandInspect(cfg *config,args ...string) error {
	if len(args)<1{
		return fmt.Errorf("inspect requires a pokemon name")
	}
	pokiInfo,ok:=cfg.caughtPokemons[args[0]]
	if !ok{
		fmt.Printf("Pokemon not caught yet\n")
		return nil
	}
	fmt.Printf("Inspecting %s\n",pokiInfo.Name)
	fmt.Printf("Height: %d\n",pokiInfo.Height)
	fmt.Printf("Weight: %d\n",pokiInfo.Weight)
	fmt.Println("Stats:")
	for _,s:=range pokiInfo.Stats{
		fmt.Printf("- %s: %d\n",s.Stat.Name,s.BaseStat)
	}
	fmt.Println("Types:")
	for _,t:=range pokiInfo.Types{
		fmt.Printf("- %s\n",t.Type.Name)
	}
	fmt.Println("Abilities:")
	for _,a:=range pokiInfo.Abilities{
		fmt.Printf("- %s\n",a.Ability.Name)
	}
	return nil
}