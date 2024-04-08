package main

import "fmt"

func main() {
	age := 32

	fmt.Println("Age: ", age)
	adultYears := getAdultYears()
	fmt.Println(adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}
