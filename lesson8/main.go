package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	dubled := dubleNumbers(&numbers)

	fmt.Println(dubled)
	doubled2 := double(numbers[])
	fmt.Println(doubled2)
}

func dubleNumbers(numbers *[]int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, val*2)
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}
