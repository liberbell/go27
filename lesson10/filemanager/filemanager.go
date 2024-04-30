package filemanager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("An error occureed: ", err)
		return nil, errors.New("Failed to open file")
	}
	scanner := bufio.NewScanner(file)
	var lines []string

	if scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Reading content file: ", err)
		file.Close()
		return
	}
	file.Close()
	return lines
}
