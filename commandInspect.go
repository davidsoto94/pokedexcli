package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config) error {
	if cfg.param1 == "" {
		return errors.New("pokemon not provided")
	}
	pokemon, ok := cfg.caughtPokemon[cfg.param1]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	return nil
}
