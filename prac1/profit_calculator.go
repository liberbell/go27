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
	if err != nil {
		fmt.Println("Revenue input error", err)
		return
	}

	expenses, err := getUserInput("Input Expence: ")
	if err != nil {
		fmt.Println("Expence input error", err)
		return
	}

	taxRate, err := getUserInput("Input Tax Rate: ")
	if err != nil {
		fmt.Println("Tax rate input error", err)
		return
	}

	ebt, profit, ratio := calculateFinancial(revenue, expenses, taxRate)

	fmt.Printf("ebt: %.1f\n", ebt)
	fmt.Printf("profit: %.1f\n", profit)
	fmt.Printf("ratio: %.1f\n", ratio)
	storeResult(ebt, profit, ratio)
}

func storeResult(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f", ebt, profit, ratio)
	os.WriteFile("result.txt", []byte(results), 0644)
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
