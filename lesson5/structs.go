package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u user) outputUserDetail() {

	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name and Last Name and Birth date must be required")
	}
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}

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
