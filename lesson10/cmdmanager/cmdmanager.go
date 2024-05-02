package cmdmanager

import "fmt"

type CMDmanager struct {
}

func (cmd CMDmanager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your price. Confirm every price with Enter")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan()

		if price == 0 {
			break
		}
		prices = append(prices, price)
	}
	return prices
}

func (cmd CMDmanager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}