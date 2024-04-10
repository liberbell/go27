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
	userbirthDate := getUserData("Please enter your birthday(MM/DD/YYYY): ")

	var appUser user
	appUser = user{
		firstName: userfirstName,
		lastName:  userfirstName,
		birthDate: userbirthDate,
		createdAt: time.Now(),
	}

	fmt.Println(firstName, lastName, birthDate)
}

func outputUserDetail(u user) {

	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserData(promptoText string) string {
	fmt.Print(promptoText)
	var value string
	fmt.Scan(&value)
	return value
}