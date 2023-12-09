package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/davidsoto94/pokedexcli/internal/pokecache"
)

func getPokemon(name, uri string, c pokecache.Cache) (Pokemon, error) {
	var response Pokemon
	res, ok := c.Get(uri)
	if !ok {
		body, err := handleRequests(uri)
		if err != nil {
			return response, err
		}
		c.Add(uri, body)
		res = body
	}
	err := json.Unmarshal(res, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func getLocations[T LocationsResponse | PokemonsInLocation](uri string, c pokecache.Cache) (T, error) {
	var response T
	res, ok := c.Get(uri)
	if !ok {
		body, err := handleRequests(uri)
		if err != nil {
			return response, err
		}
		c.Add(uri, body)
		res = body
	}
	err := json.Unmarshal(res, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func handleRequests(param string) ([]byte, error) {
	res, err := http.Get(param)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}
	return body, nil
}
