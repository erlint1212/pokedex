package main

import (
    //"os/exec"
    "os"
    "fmt"
    "github.com/erlint1212/pokedex/internal/pokecache"
)

const cliName string = "Pokedex"

type cliCommand struct {
    name        string
    description string
    callback    func(conf *config, c *pokecache.Cache) error
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
}

var tmp_struc map[string]cliCommand = map[string]cliCommand{}

func printPrompt() {
    fmt.Print(cliName, " > ")
}

func printUnkown(text string) {
    fmt.Println(text, ": command not found")
}

func commandHelp(conf *config, c *pokecache.Cache) error {
    fmt.Printf("Welcome to the %v!\n", cliName)
    fmt.Println("Usage:")
    fmt.Println()
    for _, cmdStruct := range tmp_struc {
        fmt.Println(cmdStruct.name, ": ", cmdStruct.description)
    }
    return nil
}

func commandExit(conf *config, c *pokecache.Cache) error {
    os.Exit(0)
    return nil 
}

func commandMap(conf *config, c *pokecache.Cache) error {
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

func commandMapb(conf *config, c *pokecache.Cache) error {
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
