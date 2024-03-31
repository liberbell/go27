package main

import "fmt"

func main() {
	var revenue float64
	var expences float64
	var taxRate float64

	fmt.Print("Input Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Input Expence: ")
	fmt.Scan(&expences)

	fmt.Print("Input Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expences
	profit := ebt * (1 - taxRate)
	ratio := ebt / profit

	fmt.Print(ebt)
	fmt.Print(profit)
	fmt.Print(ratio)
}
