package main

import (
	"fmt"
)

type customString string

func (text customString) log() {
	fmt.Println(text)
}

func main() {
	title, content, err := getNoteData()
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

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	var value string
	fmt.Scanln(&value)

	return value
}
