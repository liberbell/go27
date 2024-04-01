package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	// fmt.Print("Investment Amout: ")
	outputText("Investment Amout: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFunction(investmentAmount, expectedReturnRate, years, inflationRate)
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value(adjusted for Inflation: %.1f\n", futureRealValue)
	// 	fmt.Printf(`Future Value: %.1f\n
	// Future Value(adjusted for Inflation: %.1f\n`, futureValue, futureRealValue)
	// fmt.Println("Future Value: ", futureValue)
	// fmt.Println("Future Value(adjusted for Inflation): ", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFunction(investmentAmount, expectedReturnRate, years, inflationRate float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
