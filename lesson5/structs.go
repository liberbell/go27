package main

import "fmt"

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthday: ")

	fmt.Println(firstName, lastName, birthDate)
}

func getUserData(promptoText string) string {
	fmt.Println(promptoText)
	var value string
	fmt.Scan(&value)
	return value
}
