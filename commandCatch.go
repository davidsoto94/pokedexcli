package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config) error {
	if cfg.param1 == "" {
		return errors.New("pokemon not provided")
	}
	if cfg.currentPokemons == nil {
		return errors.New("you have not explore any area yet")
	}
	_, ok := cfg.currentPokemons[cfg.param1]

	if !ok {
		return errors.New("pokemon not in the current area")
	}

	pokemon, err := getPokemon(cfg.param1, fmt.Sprintf("%s/pokemon/%s", cfg.baseUrl, cfg.param1), cfg.cache)
	if err != nil {
		return err
	}
	res := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")
	if cfg.caughtPokemon == nil {
		cfg.caughtPokemon = map[string]Pokemon{}
	}
	cfg.caughtPokemon[pokemon.Name] = pokemon

	return nil
}
