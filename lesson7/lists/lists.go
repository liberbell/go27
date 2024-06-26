package lists

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	prices := []float64{2.99, 10.99}
	fmt.Println(prices[1])
	prices[1] = 12.99
	// prices[2] = 21.99
	updatePrice := append(prices, 21.99, 31.99, 43.99)
	prices = updatePrice[1:]
	fmt.Println(prices)

	discountPrice := []float64{89.99, 101.99, 45.99}
	prices = append(prices, discountPrice...)
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string
// 	productNames = [4]string{"A book"}
// 	prices := [4]float64{10.99, 1.22, 33.4, 45.2}
// 	fmt.Println(prices)
// 	productNames[2] = "A Carpet"
// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	featurePrices := prices[1:]
// 	featurePrices[0] = 199.99
// 	highlightPrices := featurePrices[:1]
// 	fmt.Println(featurePrices)
// 	fmt.Println(highlightPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(highlightPrices), cap(highlightPrices))

// 	highlightPrices = highlightPrices[:3]
// 	fmt.Println(highlightPrices)
// 	fmt.Println(len(highlightPrices), cap(highlightPrices))
// }
