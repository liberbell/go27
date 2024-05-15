package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashed_pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(hashed_pass), err
}

func CheckPassword(password, hashdPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashdPassword))
}
