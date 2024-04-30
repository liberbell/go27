package filemanager

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("An error occureed: ", err)
		return
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
}
