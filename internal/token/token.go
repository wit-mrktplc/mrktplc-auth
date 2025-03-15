package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Token is a struct that holds the necessary dependencies for the Token
type Token struct {
	Secret string
}

// NewToken creates a new instance of the Token
func NewToken(secret string) *Token {
	return &Token{Secret: secret}
}

// CreateToken creates a new JWT token
func (t *Token) CreateToken(email string) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token
	signedToken, err := token.SignedString([]byte(t.Secret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}

// ValidateToken validates a JWT token
func (t *Token) ValidateToken(tokenString string) (string, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(t.Secret), nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to parse token: %w", err)
	}

	// Check if the token is valid
	if !token.Valid {
		return "", fmt.Errorf("token is invalid")
	}

	// Get the email from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("failed to get claims")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get email")
	}

	return email, nil
}
