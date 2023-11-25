package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func main() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedexcli>")
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
		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}
