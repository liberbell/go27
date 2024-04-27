package main

import "fmt"

func main() {
	// fact := factorial(5)
	// fmt.Println(fact)
	numbers := []int{1, 10, 15}
	// sum := sumup(numbers
	sum := sumup(1, 10, 15, 40, -5)
	anothersum := sumup(1, numbers...)

	fmt.Println(sum)
	fmt.Println(anothersum)
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
	// result := 1

	// for i := 1; i <= number; i++ {
	// 	result = result * i
	// }

	// return result
}

func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}
