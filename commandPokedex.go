package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config) error {
	if cfg.caughtPokemon == nil {
		return errors.New("you have not captured any pokemon yet")
	}
	fmt.Println("your Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
