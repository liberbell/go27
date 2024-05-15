package hash

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashed_pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(hashed_pass), err
}
