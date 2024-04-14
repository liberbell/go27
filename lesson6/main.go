package main

import "fmt"

type customString string

func (text customString) log() {
	fmt.Println(text)
}

func main() {
	var name customString = "max"

	name.log()

}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	var value string
	fmt.Scan(&value)

	return value
}
