package main

import (
	"fmt"
	"os/user"

	"example.com/structs/user"
)

func main() {
	userfirstName := getUserData("Please enter your first name: ")
	userlastName := getUserData("Please enter your last name: ")
	userbirthDate := getUserData("Please enter your birthday(MM/DD/YYYY): ")

	var appUser *user.User

	// appUser, err := &user.User.NewUser(userfirstName, userlastName, userbirthDate)
	appUser = &user.User{
		FirstName: "Max",
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	// outputUserDetail(&appUser)
	appUser.OutputUserDetail()
	appUser.ClearUserName()
	appUser.OutputUserDetail()
}

func getUserData(promptoText string) string {
	fmt.Print(promptoText)
	var value string
	fmt.Scanln(&value)
	return value
}
