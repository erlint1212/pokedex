package main

import (
    //"os/exec"
    "os"
    "fmt"
)

const cliName string = "Pokedex"

type cliCommand struct {
    name        string
    description string
    callback    func() error
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
}

var tmp_struc map[string]cliCommand = map[string]cliCommand{}

func printPrompt() {
    fmt.Print(cliName, " > ")
}

func printUnkown(text string) {
    fmt.Println(text, ": command not found")
}

func commandHelp() error {
    fmt.Printf("Welcome to the %v!\n", cliName)
    fmt.Println("Usage:")
    fmt.Println()
    for _, cmdStruct := range tmp_struc {
        fmt.Println(cmdStruct.name, ": ", cmdStruct.description)
    }
    return nil
}

func commandExit() error {
    os.Exit(0)
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
