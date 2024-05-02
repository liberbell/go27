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
		prices = append(prices, price)
	}
}

func (cmd CMDmanager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}
