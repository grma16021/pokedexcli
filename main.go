package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/grma16021/pokedexcli/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     string
	Previous string
}

var commands map[string]cliCommand

var conf = &config{}

var api = "https://pokeapi.co/api/v2/location-area/"

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
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map",
			description: "Displays the previous 20 locations",
			callback:    commandMapB,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	if err := scanner.Err(); err != nil {
		fmt.Println("error reading input")
	}

	test := &config{}

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		command := cleanedInput[0]

		if val, ok := commands[command]; ok {
			val.callback(test)
		} else {
			fmt.Print("Unknown command")
		}

	}

}

func commandMapB(conf *config) error {
	url := conf.Previous

	if url == "" {
		fmt.Println("previous is empty")
		return fmt.Errorf("bøg")
	}

	internal.FetchPreviousLocations(conf.Previous)
	return nil
}

func commandMap(conf *config) error {

	if conf.Next == "" {
		_, n, p, err := internal.FetchLocations(api)
		if err != nil {
			fmt.Println(err)
		}
		conf.Next = n
		conf.Previous = p
		//fmt.Printf("conf is: %v", conf.Next)
		return nil
	} else {
		_, n, p, err := internal.FetchLocations(conf.Next)
		if err != nil {
			fmt.Println(err)
		}
		conf.Next = n
		conf.Previous = p
		//fmt.Printf("conf is: %v", conf.Next)
		return nil
	}

}

func commandHelp(*config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println("")

	for _, cmd := range commands {
		fmt.Printf("%s: %s \n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit(*config) error {
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
