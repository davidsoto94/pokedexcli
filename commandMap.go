package main

import (
	"errors"
	"fmt"
)

func commandMap(config *config) error {
	if config.next == "" {
		config.next = "https://pokeapi.co/api/v2/location-area/"
	}
	res, err := getLocations[LocationsResponse](config.next, config.cache)
	if err != nil {
		return err
	}
	config.next = res.Next
	config.previous = res.Previous
	for _, val := range res.Results {
		fmt.Println(val.Name)
	}
	return nil
}

func commandMapb(config *config) error {
	if config.previous == nil {
		return errors.New("can't go back")
	}
	res, err := getLocations[LocationsResponse](*config.previous, config.cache)
	if err != nil {
		return err
	}
	config.next = res.Next
	config.previous = res.Previous
	for _, val := range res.Results {
		fmt.Println(val.Name)
	}
	return nil
}
