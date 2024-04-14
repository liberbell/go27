package main

import (
	"fmt"

	"example.com/note/note"
)

type customString string

func (text customString) log() {
	fmt.Println(text)
}

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
}

func getNoteData() (string, string) {
	title := getUserInput("Note tilte: ")

	content := getUserInput("Note content: ")

	return title, content

}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	var value string
	fmt.Scanln(&value)

	return value
}
