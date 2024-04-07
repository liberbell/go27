package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading balance file: ", err)
		return 1000, errors.New("Failed to find file.")
	}
	valueText := string(data)
	balance, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		fmt.Println("Error parsing balance: ", err)
		return 1000, errors.New("Failed to parse stored value")
	}
	return balance, nil
}

func writeFloatTofile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
