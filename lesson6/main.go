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
		fmt.Println(err)
		return
	}

	content, err := getUserInput("Note content: ")
	if err != nil {
		fmt.Println(err)
		return
	}

}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note tilte: ")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	content, err := getUserInput("Note content: ")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	return title, content, nil

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
