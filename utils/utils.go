package utils

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

var errNotFound = errors.New("value not found")
var errInvalidEmail = errors.New("invalid email")

func Contains(s []string, str string) error {
	for _, v := range s {
		if v == str {
			return nil
		}
	}

	return errNotFound
}

func IsEmailValid(e string) error {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(e) {
		return errInvalidEmail
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
