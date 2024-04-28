package main

import (
	"fmt"

	_ "example.com/price-calculator/prices"
)

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
	}
	fmt.Println(result)
}
