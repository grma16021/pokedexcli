package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func main() {

	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		command := cleanedInput[0]

		if val, ok := commands[command]; ok {
			val.callback()
		} else {
			fmt.Print("Unknown command")
		}

	}

}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println("")

	for _, cmd := range commands {
		fmt.Printf("%s: %s \n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil

}

func cleanInput(text string) []string {
	var words []string
	word := ""

	cleanedSpaces := strings.TrimSpace(text)
	formatedText := strings.ToLower(cleanedSpaces)
	for i := 0; i < len(formatedText); i++ {
		char := formatedText[i]

		if char != 32 {
			word += string(char)
		} else {
			words = append(words, word)
			word = ""
		}
	}
	if len(word) > 0 {
		words = append(words, word)
	}
	return words
}
