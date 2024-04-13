package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthday(MM/DD/YYYY): ")

	var appUser *user.User

	// appUser, err := &user.User.NewUser(userfirstName, userlastName, userbirthDate)
	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("admin@example.com", "adminpass")

	admin.User.OutputUserDetail()
	admin.User.ClearUserName()
	admin.User.OutputUserDetail()

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
