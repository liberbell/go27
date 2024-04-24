package main

import (
	"fmt"

	"golang.org/x/text/number"
)

type transformFn func(int) int

type anotherFn func(int, []string, map[string][]int) ([]int, string)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled, tripled)
	// doubled2 := double(numbers[])
	// fmt.Println(doubled2)
}

func dubleNumbers(numbers *[]int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, val*2)
	}
	return dNumbers
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func getTransformerFunction(numbers, *[]int) transformFn {
	if (*numbers[0]) ==0
	return double

}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
