package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthday: ")

	fmt.Println(firstName, lastName, birthDate)
}

func outputUserDetail(firstName, lastName, birthDate string) {

	fmt.Println(firstName, lastName, birthDate)
}

func getUserData(promptoText string) string {
	fmt.Print(promptoText)
	var value string
	fmt.Scan(&value)
	return value
}
