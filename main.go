package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		firstWord := cleanedInput[0]

		fmt.Println("Your command was: " + firstWord)
	}

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
