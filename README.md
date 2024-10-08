# Pokedex

## Guided project from boot.dev [link](https://www.boot.dev/courses/build-pokedex-cli)

### Description

Build a REPL pokedex on the command line in Go. Learn how to use JSON, make network requests, and implement caching.

A REPL, or Read-Eval-Print Loop, is a simple interactive programming environment that takes user input, evaluates it, and returns the result to the user. In this guided project, you'll build a Pokedex-like REPL in Go that uses the PokeAPI to fetch data about Pokemon. It's a great way to put your Go knowledge to the test and learn valuable skills like HTTP networking and data serialization.

### Learning goals

* Learn how to parse JSON in Go
* Practice making HTTP requests in Go
* Learn how to build a CLI tool that makes interacting with a back-end server easier
* Get hands-on practice with local Go development and tooling
* Learn about caching and how to use it to improve performance

### Improvements

* Make it so cache and API functions only send relevant information instead of the whole API response, memory efficency.

### Extension ideas

* Update the CLI to support the "up" arrow to cycle through previous commands
* Simulate battles between pokemon
* Add more unit tests
* Refactor your code to organize it better and make it more testable
* Keep pokemon in a "party" and allow them to level up
* Allow for pokemon that are caught to evolve after a set amount of time
* Persist a user's Pokedex to disk so they can save progress between sessions
* Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
* Random encounters with wild pokemon
* Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon

