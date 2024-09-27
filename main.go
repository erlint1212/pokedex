package main
import (
    "fmt"
    "bufio"
    "os"
    "github.com/erlint1212/pokedex/internal/pokecache"
    "time"
)

func main() {
    // solves the A ref B ref A problem
    tmp_struc = commands
    config_var := config{"https://pokeapi.co/api/v2/location-area/", nil}
    cache := pokecache.NewCache(5 * 60 * 1000 *time.Millisecond)

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
        if err := command.callback(&config_var, &cache); err != nil {
            fmt.Println(err)
        }
        printPrompt()
    }
    fmt.Println()
}
