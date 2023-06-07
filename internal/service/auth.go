package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
	"wallet/config"
	"wallet/internal/repo"
)

type AuthService struct {
	userRepo repo.UserRepo
}

func NewAuthService(userRepo repo.UserRepo) AuthService {
	return AuthService{
		userRepo: userRepo,
	}
}

type Claims struct {
	XUserID string `json:"x-user-id"`
	jwt.RegisteredClaims
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

func (s *AuthService) ValidJWTToken(token string, requiredRole string) error {
	claims := &Claims{}
	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
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
	if claims.XUserID == "" {
		return fmt.Errorf("unauthorized")
	}

	_, err = s.userRepo.GetUserByID(claims.XUserID)
	if err != nil {
		return fmt.Errorf("unauthorized")
	}

	return nil
}
