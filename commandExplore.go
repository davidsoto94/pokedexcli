package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config) error {
	if cfg.param1 == "" {
		return errors.New("area not provided")
	}
	res, err := getLocations[PokemonsInLocation](fmt.Sprintf("%s/location-area/%s", cfg.baseUrl, cfg.param1), cfg.cache)
	if err != nil {
		return err
	}
	cfg.currentPokemons = make(map[string]string)
	for _, val := range res.PokemonEncounters {
		cfg.currentPokemons[val.Pokemon.Name] = val.Pokemon.Name
		fmt.Println(val.Pokemon.Name)
	}
	return nil
}
