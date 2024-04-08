package main

import "fmt"

func main() {
	age := 32

	var agePointer *int
	agePointer = &age

	fmt.Println("Age: ", agePointer)
	fmt.Println("Age: ", *agePointer)
	getAdultYears(agePointer)
	fmt.Println("Age: ", age)
}

func getAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
