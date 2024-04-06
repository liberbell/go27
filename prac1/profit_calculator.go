package main

import (
	"fmt"
	"os"

	"errors"
)

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	revenue, err := getUserInput("Input Revenue: ")
	panic(err)

	expenses, err := getUserInput("Input Expence: ")
	panic(err)

	taxRate, err := getUserInput("Input Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateFinancial(revenue, expenses, taxRate)

	fmt.Printf("ebt: %.1f\n", ebt)
	fmt.Printf("profit: %.1f\n", profit)
	fmt.Printf("ratio: %.1f\n", ratio)
}

func storeResult(ebt, profit, ratio float64) {
	os.WriteFile("result.txt")
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		fmt.Println("Invalid input value")
		return 0, errors.New("Value must be positive")
	}
	return userInput, nil
}

func calculateFinancial(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
