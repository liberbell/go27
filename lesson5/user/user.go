package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

type Admin struct {
	email    string
	password string
}

func (u User) OutputUserDetail() {

	fmt.Println(u.FirstName, u.LastName, u.BirthDate)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name and Last Name and Birth date must be required")
	}
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthdate,
		CreatedAt: time.Now(),
	}, nil
}
