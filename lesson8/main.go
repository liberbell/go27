package main

import (
	"fmt"
)

type transformFn func(int) int

type anotherFn func(int, []string, map[string][]int) ([]int, string)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	// more_numbers := []int{6, 1, 2, 3, 4}
	// doubled := transformNumbers(&numbers, double)
	double := createTransformer(2)
	// tripled := transformNumbers(&numbers, triple)
	triple := createTransformer(3)

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled, tripled)
	// doubled2 := double(numbers[])
	// fmt.Println(doubled2)

	// transformerFn1 := getTransformerFunction(&numbers)
	// transformerFn2 := getTransformerFunction(&more_numbers)

	// transformedNumbers := transformNumbers(&numbers, transformerFn1)
	// more_transformedNumbers := transformNumbers(&more_numbers, transformerFn2)

	// fmt.Println(transformedNumbers)
	// fmt.Println(more_transformedNumbers)

	// transformed := transformNumbers(&numbers, func(number int) int {
	// 	return number * 2
	// })
	// fmt.Println(transformed)

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

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
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
