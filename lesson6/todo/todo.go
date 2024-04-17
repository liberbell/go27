package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Todo struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (todo Todo) Display() {
	fmt.Printf("Your note titled %v has the following content: \n\n%v\n\n", todo.Title, todo.Content)
}

func (todo Todo) Save() error {
	filename := strings.ReplaceAll(todo.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0644)
}

func New(title, content string) (Todo, error) {
	if title == "" {
		return Note{}, errors.New("Invalid input.")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
