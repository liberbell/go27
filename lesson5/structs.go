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
	userfirstName := getUserData("Please enter your first name: ")
	userlastName := getUserData("Please enter your last name: ")
	userbirthDate := getUserData("Please enter your birthday: ")

	var appUser user
	appUser = user{
		firstName: userfirstName,
		lastName:  userfirstName,
		birthDate: userbirthDate,
		createdAt: time.Now(),
	}

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
