package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type mapData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func FetchLocations(api string) (mapData, string, string, error) {

	var mapDat mapData

	//var api = "https://pokeapi.co/api/v2/location-area/"

	resp, err := http.Get(api)
	if err != nil {
		return mapDat, "", "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body, &mapDat)
	if err != nil {
		fmt.Println(err)
	}

	for _, location := range mapDat.Results {
		fmt.Println(location.Name)
	}

	prev, ok := mapDat.Previous.(string)
	if !ok {
		fmt.Println("error cant cast prev as string")
		prev = ""
	}

	return mapDat, mapDat.Next, prev, nil
}

func FetchPreviousLocations(url string) error {
	var mapDat mapData

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &mapDat)
	if err != nil {
		return err
	}

	for _, location := range mapDat.Results {
		fmt.Println(location.Name)
	}
	return nil
}
