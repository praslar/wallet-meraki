package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
	"wallet/config"
	"wallet/internal/repo"
)

type Claims struct {
	XUserID      string `json:"x-user-id"`
	RequiredRole string `json:"x-user-role"`
	jwt.RegisteredClaims
}

type AuthService struct {
	userRepo repo.UserRepo
}

func NewAuthService(userRepo repo.UserRepo) AuthService {
	return AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) GenJWTToken(userID string, key string) (string, error) {
	// Declare the expiration time of the token
	expirationTime := time.Now().Add(24 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		XUserID:      userID,
		RequiredRole: key,
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
func (s *AuthService) ValidJWTToken(token string, requiredRole string) error {
	claims := &Claims{}
	secret := config.LoadEnv().Secret
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return err
	}
	if !tkn.Valid {
		return fmt.Errorf("unauthorized")
	}

	if claims.RequiredRole != requiredRole {
		return fmt.Errorf("unauthorized")
	}

	_, err = s.userRepo.GetUserByID(claims.XUserID)
	if err != nil {
		return fmt.Errorf("unauthorized")
	}

	return nil
}
