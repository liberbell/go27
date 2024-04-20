package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	prices := [4]float64{10.99, 1.22, 33.4, 45.2}
	fmt.Println(prices)
}
