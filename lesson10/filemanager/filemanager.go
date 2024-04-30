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
		file.Close()
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
		file.Close()
		return nil, errors.New("Failed to read line in file")
	}
	file.Close()
	return lines, nil
}
