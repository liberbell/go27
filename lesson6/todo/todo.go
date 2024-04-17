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

func (note Todo) Display() {
	fmt.Printf("Your note titled %v has the following content: \n\n%v\n\n", note.Title, note.Content)
}

func (note Todo) Save() error {
	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"
	json, err := json.Marshal(note)
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
