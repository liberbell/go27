package main

import "fmt"

func main() {
	var revenue int
	var expences int
	var taxRate float64

	fmt.Print("Input Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Input Expence: ")
	fmt.Scan(&expences)

	fmt.Print("Input Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expences
	profit := float64(ebt) * (1 - taxRate)
}
