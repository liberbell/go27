package main

import (
	"fmt"
)

type transformFn func(int) int

type anotherFn func(int, []string, map[string][]int) ([]int, string)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	more_numbers := []int{6, 1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled, tripled)
	// doubled2 := double(numbers[])
	// fmt.Println(doubled2)

	transformFn1 := getTransformerFunction(&numbers)
	transformFn2 := getTransformerFunction(&more_numbers)

	transformNumbers := transformNumbers(&numbers, transformFn1)
	more_transformNumbers := transformNumbers(&more_numbers, transformFn2)

	fmt.Println(transformNumbers)
	fmt.Println(more_transformNumbers)
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

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 0 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
