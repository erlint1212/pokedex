package main

import (
    "encoding/json"
    "net/http"
    "io"
)

/*
type NamedAPIResource struct {
    name    string `json:"name"`
    url     string  `json:"url"`
}

type VersionDetails struct {
    rate        int                 `json:"rate"`
    version     NamedAPIResource    `json:"version"`
}

type LocationArea struct {
    id                      int                     `json:"id"`
    name                    string                  `json:"name"`
    game_index              int                     `json:"game_index"`
    encounter_method_rates  []struct{
        encounter_method    NamedAPIResource  `json:"encounter_method"`
        version_details     []struct {
            rate        int                 `json:"rate"`
            version     NamedAPIResource    `json:"version"`
        } `json:"version_details"`
    }   `json:"encounter_method_rates"`
    location                NamedAPIResource        `json:"location"`
    names                   []struct {
        name        string  `json:"name"`
        language    NamedAPIResource   `json:"language"`
    }   `json:"names"`
    pokemon_encounters      []struct {
        pokemon     NamedAPIResource    `json:"pokemon"`
        version_details     []struct{
            version             NamedAPIResource    `json:"version"`
            max_chance          int                 `json:"max_chance"`
            encounter_details   []struct{
                min_level   int     `json:"min_level"`
                max_level   int     `json:"max_level"`
                condition_values    []NamedAPIResource  `json:"condition_values"`
                chance      int     `json:"chance"`
                method      NamedAPIResource    `json:"method"`
            }  `json:"encounter_details"`
        }  `json:"version_details"`
    }          `json:"pokemon_encounters"`
}
*/

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


func getLocations(url string) (LocationArea, error) {
    //url = "https://pokeapi.co/api/v2/location-area/1/"
	res, err := http.Get(url)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

    
    data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	var loc_area LocationArea
    if err = json.Unmarshal(data, &loc_area); err != nil {
		return LocationArea{}, err
	}
    /*
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&loc_area); err != nil {
		return LocationArea{}, fmt.Errorf("error decoding response body: %w", err)
	}
    */

	return loc_area, nil
}
