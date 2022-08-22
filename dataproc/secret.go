package dataproc

import (
	"golang.org/x/crypto/bcrypt"
)

// SecretHashEncode hash bcrypt encode
func SecretHashEncode(password string) string {

	secret, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(secret)
}

// SecretHashCompare hash bcrypt compare
func SecretHashCompare(hashPassword, inputPassword string) bool {

	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(inputPassword)); err != nil {
		return false
	}

	return true
}
