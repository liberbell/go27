package conversion

import (
	"fmt"
	"strconv"
)

func StringToFloat(strings []string) ([]float64, error) {
	for stringIndex, string := range strings {
		floatPrice, err := strconv.ParseFloat(string, 64)
		if err != nil {
			file.Close()
			fmt.Println("Convert to float failed:", err)
		}
		prices[lineIndex] = floatPrice
	}
	job.InputPrices = prices
}
