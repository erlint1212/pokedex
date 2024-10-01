package main

import (
    "encoding/json"
    "net/http"
    "io"
    "github.com/erlint1212/pokedex/internal/pokecache"
    "fmt"
)

type LocationAreaEndpoint struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getExplore(c *pokecache.Cache, url string) (LocationAreaEndpoint, error) {
    if data, OK := c.Get(url); OK {
        fmt.Println("Found in cache!")
        var loc_area_end LocationAreaEndpoint
        if err := json.Unmarshal(data, &loc_area_end); err != nil {
            return LocationAreaEndpoint{}, err
        }
        return loc_area_end, nil
    }
	res, err := http.Get(url)
	if err != nil {
		return LocationAreaEndpoint{}, err
	}
	defer res.Body.Close()

    
    data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaEndpoint{}, err
	}

    c.Add(url, data)

	var loc_area_end LocationAreaEndpoint
    if err = json.Unmarshal(data, &loc_area_end); err != nil {
		return LocationAreaEndpoint{}, err
	}


	return loc_area_end, nil
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
