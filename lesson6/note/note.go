package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content: \n\n%v\n\n", note.title, note.content)
}

func (note Note) Save() error {
	filename := strings.ReplaceAll(note.title, " ", "_")
	filename = strings.ToLower(filename)
	json, err := json.Marshal(filename)
	if err != nil {
		return err
	}
	os.WriteFile(filename, json, 0644)
}

func New(title, content string) (Note, error) {
	if title == "" {
		return Note{}, errors.New("Invalid input.")
	}

	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}
