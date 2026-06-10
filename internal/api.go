package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/grma16021/pokedexcli/internal/pokecache"
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

type pokemonLocationData struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func FetchLocations(api string, cache *pokecache.Cache) (mapData, string, string, error) {

	var mapDat mapData
	var body []byte
	if cached, ok := cache.Get(api); ok {
		body = cached
	} else {

		resp, err := http.Get(api)
		if err != nil {
			return mapDat, "", "", err
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		cache.Add(api, body)

	}

	err := json.Unmarshal(body, &mapDat)
	if err != nil {
		fmt.Println(err)
	}

	for _, location := range mapDat.Results {
		fmt.Println(location.Name)
	}

	prev, ok := mapDat.Previous.(string)
	if !ok {
		//fmt.Println("error cant cast prev as string")
		prev = ""
	}

	return mapDat, mapDat.Next, prev, nil
}

func FetchPreviousLocations(url string) error {
	var mapDat mapData
	var data []byte

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
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

func FetchPokemonLocation(name string, cache *pokecache.Cache) error {
	var pokemonLocationDat = pokemonLocationData{}
	var body []byte

	api := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", name)
	if cached, ok := cache.Get(api); ok {
		body = cached
	} else {

		res, err := http.Get(api)

		if err != nil {
			return err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		cache.Add(api, body)
	}

	err := json.Unmarshal(body, &pokemonLocationDat)
	if err != nil {
		return err
	}

	for _, name := range pokemonLocationDat.PokemonEncounters {
		fmt.Println("- " + name.Pokemon.Name)
	}

	return nil

}

func FetchPokemonInfo(name string, cache *pokecache.Cache) (Pokemon, error) {
	var Pmon = Pokemon{}
	var body []byte

	api := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
	if cached, ok := cache.Get(api); ok {
		body = cached
	} else {
		res, err := http.Get(api)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}

		cache.Add(api, body)
	}

	err := json.Unmarshal(body, &Pmon)
	if err != nil {
		return Pokemon{}, err
	}

	return Pmon, nil
}
