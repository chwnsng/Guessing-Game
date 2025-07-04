package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// creating a sample secret
var jwtSecret = []byte("placeholder-secret-which-should-be-changed")

// type Claims struct {
// 	Username string `json:"username"`
// 	jwt.RegisteredClaims
// }

func CreateToken(username string) (string, error) {
	// creating header.payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(), // Let the token be valid for 24 hours
		})

	// encoding & signing the token --> header64.payload64.signature
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// verifying tokens
func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return err
	}

	// check if token is valid
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	// check if claims is valid
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("invalid claims format")
	}

	// check if token hasn't expried
	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return fmt.Errorf("token expried!")
		}
	}

	return nil
}
