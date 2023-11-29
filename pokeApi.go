package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/davidsoto94/pokedexcli/internal/pokecache"
)

func getLocations[T LocationsResponse | PokemonsInLocation](uri string, c pokecache.Cache) (T, error) {
	var reponse T
	res, ok := c.Get(uri)
	if !ok {
		body, err := handleRequests(uri)
		if err != nil {
			return reponse, err
		}
		c.Add(uri, body)
		res = body
	}
	err := json.Unmarshal(res, &reponse)
	if err != nil {
		return reponse, err
	}
	return reponse, nil
}

func handleRequests(param string) ([]byte, error) {
	res, err := http.Get(param)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, errors.New(fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body))
	}
	if err != nil {
		return nil, err
	}
	return body, nil
}
