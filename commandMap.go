package main

import (
	"errors"
	"fmt"
)

func commandMap(config *config) error {
	res, err := getNextLocations(config.next)
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
		return errors.New("Can't go back")
	}
	res, err := getPreviousLocations(*config.previous)
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
