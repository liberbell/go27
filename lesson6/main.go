package main

import (
	"errors"
	"fmt"
)

type customString string

func (text customString) log() {
	fmt.Println(text)
}

func main() {
	title, err := getUserInput("Note tilte: ")
	if err != nil {
		return errors
	}
	content := getUserInput("Note content: ")

}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)

	var value string
	fmt.Scan(&value)

	if value == "" {
		return "", errors.New("Invalid input.")
	}

	return value, nil
}
