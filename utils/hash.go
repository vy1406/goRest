package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) // 14 or bcrypt.DefaultCost

	return string(bytes), err
}
