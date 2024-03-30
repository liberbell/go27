package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	years := 10.0

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println(futureValue)
}
