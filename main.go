package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/davidsoto94/pokedexcli/internal/pokecache"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	cfg := config{
		baseUrl:         "https://pokeapi.co/api/v2",
		next:            "",
		previous:        nil,
		param1:          "",
		currentPokemons: nil,
		cache:           pokecache.NewCache(5 * time.Minute),
	}
	for {
		fmt.Print("Pokedex >")
		reader.Scan()
		text := cleanInput(reader.Text())
		if len(text) == 0 {
			continue
		}
		command, ok := getCommands()[text[0]]
		if !ok {
			fmt.Println(text, ": command not found")
			continue
		}
		if len(text) > 1 {
			cfg.param1 = text[1]
		} else {
			cfg.param1 = ""
		}
		err := command.callback(&cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}
