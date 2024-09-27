package main

import (
    "encoding/json"
    "net/http"
    "io"
    "github.com/erlint1212/pokedex/internal/pokecache"
    "fmt"
)


type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


func getLocations(c *pokecache.Cache, url string) (LocationArea, error) {
    //url = "https://pokeapi.co/api/v2/location-area/1/"
    if data, OK := c.Get(url); OK {
        fmt.Println("Found in cache!")
        var loc_area LocationArea
        if err := json.Unmarshal(data, &loc_area); err != nil {
            return LocationArea{}, err
        }
        return loc_area, nil
    }
	res, err := http.Get(url)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

    
    data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

    c.Add(url, data)

	var loc_area LocationArea
    if err = json.Unmarshal(data, &loc_area); err != nil {
		return LocationArea{}, err
	}


	return loc_area, nil
}
