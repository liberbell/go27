package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	dubled := dubleNumbers(&numbers)

	fmt.Println(dubled)
}

func dubleNumbers(numbers *[]int) int {
	dNumbers := []int{}
	for _, vl := range *numbers {
		dNumbers = (dNumbers, val*2)
	}
	return dNumbers[]
}
