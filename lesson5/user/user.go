package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u User) OutputUserDetail() {

	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name and Last Name and Birth date must be required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}
