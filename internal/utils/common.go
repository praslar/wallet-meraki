package utils

import (
	"golang.org/x/crypto/bcrypt"
	"net/mail"
)

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ComparePassword(password string, hash string) bool {
	return hash == password
}

func HashPassword(password string) (string, error) {
	// Tạo salt ngẫu nhiên
	salt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// Trả về password đã được mã hóa và kèm salt
	return string(salt), nil
}
