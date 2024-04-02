package main

import "fmt"

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	revenue := getUserInput("Input Revenue: ")
	expenses := getUserInput("Input Expence: ")
	taxRate := getUserInput("Input Tax Rate: ")

	ebt, profit, ratio := calculateFinancial(revenue, expenses, taxRate)

	fmt.Printf("ebt: %.1f\n", ebt)
	fmt.Printf("profit: %.1f\n", profit)
	fmt.Printf("ratio: %.1f\n", ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancial(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
