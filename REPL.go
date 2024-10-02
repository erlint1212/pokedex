package main

import (
    //"os/exec"
    "os"
    "fmt"
    "github.com/erlint1212/pokedex/internal/pokecache"
    "math/rand"
    //"math"
)

const cliName string = "Pokedex"
var Pokedex map[string]PokemonEndpoint = map[string]PokemonEndpoint{}

type cliCommand struct {
    name        string
    description string
    callback    func(conf *config, c *pokecache.Cache, arg string) error
}

type config struct {
    Next        any
    Previous    any
}


var commands map[string]cliCommand = map[string]cliCommand{
        "help": cliCommand{
            name:        "help",
            description: "Displays a help message",
            callback:    commandHelp,
        },
        "exit": cliCommand{
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
        "map": cliCommand{
            name:        "map",
            description: "The map command displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the next 20 locations, and so on. The idea is that the map command will let us explore the world of Pokemon.",
            callback:    commandMap,

        },
        "mapb": cliCommand{
            name:        "mapb",
            description: "Displays the previous 20 locations. It's a way to go back.",
            callback:    commandMapb,
        },
        "explore": cliCommand{
            name:        "explore",
            description: "Shows all the locations for a area",
            callback:    commandExplore,
        },
        "catch": cliCommand{
            name:        "catch",
            description: "Try to catch given pokemon, catch <pokemon name>, chance based on base exp",
            callback:    commandCatch,
        },
        "inspect": cliCommand{
            name:        "inspect",
            description: "Inspect your caught pokemon",
            callback:    commandInspect,
        },
}

var tmp_struc map[string]cliCommand = map[string]cliCommand{}

func printPrompt() {
    fmt.Print(cliName, " > ")
}

func printUnkown(text string) {
    fmt.Println(text, ": command not found")
}

func commandHelp(conf *config, c *pokecache.Cache, arg string) error {
    fmt.Printf("Welcome to the %v!\n", cliName)
    fmt.Println("Usage:")
    fmt.Println()
    for _, cmdStruct := range tmp_struc {
        fmt.Println(cmdStruct.name, ": ", cmdStruct.description)
    }
    return nil
}

func commandExit(conf *config, c *pokecache.Cache, arg string) error {
    os.Exit(0)
    return nil 
}

func commandMap(conf *config, c *pokecache.Cache, arg string) error {
    next_url, ok := conf.Next.(string)
    if !ok {
        return fmt.Errorf("Next url must be a string")
    }
    loc, err := getLocations(c ,next_url) 
    if err != nil {
        return err
    }
    conf.Next = loc.Next
    conf.Previous = loc.Previous
    for _, loc := range loc.Results {
        fmt.Println(loc.Name)
    }
    return nil
}

func commandMapb(conf *config, c *pokecache.Cache, arg string) error {
    prev_url, ok := conf.Previous.(string)
    if !ok {
        fmt.Println("Start of map, Previous = nil")
        return nil
    }
    loc, err := getLocations(c, prev_url) 
    if err != nil {
        return err
    }
    conf.Next = loc.Next
    conf.Previous = loc.Previous
    for _, loc := range loc.Results {
        fmt.Println(loc.Name)
    }
    return nil
}

func commandExplore(conf *config, c *pokecache.Cache, arg string) error {
    fmt.Printf("Exploring %s...\n", arg)
    url := "https://pokeapi.co/api/v2/location-area/" + arg
    loc_end, err := getExplore(c, url)
    if err != nil {
        return err
    }
    fmt.Println("Found Pokemon:")
    for _, pokemon_encounter := range loc_end.PokemonEncounters {
        fmt.Println(" - ", pokemon_encounter.Pokemon.Name)
    }
    return nil
}

func commandCatch(conf *config, c *pokecache.Cache, arg string) error {
    pokemon_name := arg
    fmt.Printf("Throwing a Pokeball at %s...\n", pokemon_name)
    url := "https://pokeapi.co/api/v2/pokemon/" + pokemon_name
    poke_end, err := getPokemon(c, url)
    if err != nil {
        return fmt.Errorf("Pokemon not found : %v", err)
    }
    num := rand.Intn(100)
    //chance := int((1 / (1 + math.Exp(-float64(poke_end.BaseExperience))))*100)
    chance := int((float64(poke_end.BaseExperience) / float64(390))*100)
    switch {
    case num >= chance:
        fmt.Printf("%s was caught!\n", pokemon_name)
        Pokedex[pokemon_name] = poke_end
        return nil
    case num < chance:
        fmt.Printf("%s escaped!\n", pokemon_name)
        return nil
    default:
        return fmt.Errorf("Something went wrong with the rand.Intn function")
    }
}

func commandInspect(conf *config, c *pokecache.Cache, arg string) error {
    pokemon_name := arg
    data, OK := Pokedex[pokemon_name]
    if !OK {
        fmt.Println("you have not caught that pokemon")
        return nil
    }
    fmt.Println("Name:", data.Name)
    fmt.Println("Height:", data.Height)
    fmt.Println("Weight:", data.Weight)
    fmt.Println("Stats:")
    for _, stat := range data.Stats {
        fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
    }
    fmt.Println("Types:")
    for _, item := range data.Types {
        fmt.Printf("  - %s\n", item.Type.Name)
    }
    return nil
}

/*
func clearScreen() error{
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    if err := cmd.Run(); err != nil {
        return err
    }
    return nil
}

func handleInvalidCmd(text string) {
    defer printUnkown(text)
}

func handleCmd(text string) {
    handleInvalidCmd(text)
}

*/
