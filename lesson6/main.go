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
	var name customString = "max"

	name.log()

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
