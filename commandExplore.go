package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config) error {
	if cfg.param1 == "" {
		return errors.New("Area not provided")
	}
	res, err := getLocations[PokemonsInLocation](fmt.Sprintf("%s/location-area/%s", cfg.baseUrl, cfg.param1), cfg.cache)
	if err != nil {
		return err
	}
	for _, val := range res.PokemonEncounters {
		fmt.Println(val.Pokemon.Name)
	}
	return nil
}
