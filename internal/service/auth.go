package service

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
	"wallet/config"
)

type Claims struct {
	XUserID string `json:"x-user-id"`
	jwt.RegisteredClaims
}

type AuthService struct {
}

func NewAuthService() AuthService {
	return AuthService{}
}

func (s *AuthService) GenJWTToken(userID string) (string, error) {
	// Declare the expiration time of the token
	expirationTime := time.Now().Add(24 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		XUserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	secret := config.LoadEnv().Secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, err
}
