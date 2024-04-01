package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Input Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Input Expence: ")
	fmt.Scan(&expenses)

	fmt.Print("Input Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

func getUserInput(infoText string) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
}
