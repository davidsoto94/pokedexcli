package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func getNextLocations(next string) (Response, error) {
	if next == "" {
		next = "https://pokeapi.co/api/v2/location-area/"
	}
	body, err := handleRequests(next)
	if err != nil {
		return Response{}, err
	}
	reponse := Response{}
	err = json.Unmarshal(body, &reponse)
	if err != nil {
		return Response{}, err
	}
	return reponse, nil
}

func getPreviousLocations(previous string) (Response, error) {
	body, err := handleRequests(previous)
	if err != nil {
		return Response{}, err
	}
	reponse := Response{}
	err = json.Unmarshal(body, &reponse)
	if err != nil {
		return Response{}, errors.New(err.Error())
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
