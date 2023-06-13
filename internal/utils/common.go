package utils

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
)

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

<<<<<<< HEAD
func ComparePassword(password string, hash string) bool {
	return hash == password
=======
func ComparePassword(password string, hashedPassword string) bool {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		logrus.Errorf("Fail to hash password: %v", err.Error())
	}

	// So sánh password đã nhập với password đã được mã hóa trong cơ sở dữ liệu
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		logrus.Errorf("Wrong password: %v", err.Error())
	}
	return err == nil
}

func HashPassword(password string) (string, error) {
	// Tạo salt ngẫu nhiên
	salt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// Trả về password đã được mã hóa và kèm salt
	return string(salt), nil
>>>>>>> f42f72261765b586a57e931f5a776a40c861c8d0
}
