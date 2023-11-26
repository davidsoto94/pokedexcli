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
		next:     "",
		previous: nil,
		cache:    pokecache.NewCache(5 * time.Minute),
	}
	for {
		fmt.Print("pokedexcli >")
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
		err := command.callback(&cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}
