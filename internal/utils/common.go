package utils

import "net/mail"

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ComparePassword(password string, hash string) bool {
	return hash == password
}
