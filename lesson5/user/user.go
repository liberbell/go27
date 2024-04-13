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

func (u user) OutputUserDetail() {

	fmt.Println(u.FirstName, u.LastName, u.BirthDate)
}

func (u *user) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
}

func NewUser(firstName, lastName, birthdate string) (*User, error) {
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
