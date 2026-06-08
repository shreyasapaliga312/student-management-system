package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
)

type Payload struct {
	Username string
	Email    string
	Id       uint
}

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Id       uint   `json:"id"`
	jwt.StandardClaims
}

var jwtSecret string

var validate = validator.New()

func ValidateStruct(s interface{}) error {
	if err := validate.Struct(s); err != nil {
		return err
	}
	return nil
}

func GenerateJWTToken(payload Payload) (string, error) {
	if jwtSecret = os.Getenv("JWT_SECRET"); jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable not provided!\n")
	}

	key := []byte(jwtSecret)

	expirationTime := time.Now().Add(5 * 24 * 60 * time.Minute) // Will expire in 5 days

	claims := &Claims{
		Id:       payload.Id,
		Username: payload.Username,
		Email:    payload.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	SignedToken, err := unsignedToken.SignedString(key)
	if err != nil {
		return "", err
	}
	return SignedToken, nil
}
func VerifyJWTToken(strToken string) (*Claims, error) {
	if jwtSecret = os.Getenv("JWT_SECRET"); jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable not provided!\n")
	}

	key := []byte(jwtSecret)

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(strToken, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return claims, fmt.Errorf("invalid token signature")
	}

	if !token.Valid {
		return claims, fmt.Errorf("invalid token")
	}

	return claims, nil
}
