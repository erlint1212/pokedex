package main
import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    // solves the A ref B ref A problem
    tmp_struc = commands
    config_var := config{"https://pokeapi.co/api/v2/location-area/", nil}

    reader := bufio.NewScanner(os.Stdin)
    printPrompt()

    for reader.Scan() {
        text := reader.Text()
        command, OK := commands[text]
        if !OK {
           printUnkown(text) 
           printPrompt()
           continue
        }
        if err := command.callback(&config_var); err != nil {
            fmt.Println(err)
        }
        printPrompt()
    }
    fmt.Println()
}
