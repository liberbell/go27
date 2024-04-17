package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Printf("Your note titled %v has the following content: \n\n%v\n\n", todo.Text, todo.Content)
}

func (todo Todo) Save() error {
	filename := "todo.json"
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Invalid input.")
	}

	return Todo{
		Text: content,
	}, nil
}
