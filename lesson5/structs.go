package main

import (
	"fmt"
)

func main() {
	userfirstName := getUserData("Please enter your first name: ")
	userlastName := getUserData("Please enter your last name: ")
	userbirthDate := getUserData("Please enter your birthday(MM/DD/YYYY): ")

	var appUser *user
	appUser, err := newUser(userfirstName, userlastName, userbirthDate)
	if err != nil {
		fmt.Println(err)
		return
	}

	// outputUserDetail(&appUser)
	appUser.outputUserDetail()
	appUser.clearUserName()
	appUser.outputUserDetail()
}

func getUserData(promptoText string) string {
	fmt.Print(promptoText)
	var value string
	fmt.Scanln(&value)
	return value
}
