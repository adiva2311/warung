package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/gommon/log"
)

type jwtCustomClaims struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	PhoneNumber string `json:"phone_number"`
	jwt.RegisteredClaims
}

// var secretKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(id int, email string, role string, phone_number string) (string, error) {
	// Set custom claims
	customClaims := &jwtCustomClaims{
		ID:          uint(id),
		Email:       email,
		Role:        role,
		PhoneNumber: phone_number,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)

	// Generate encoded token and send it as response.
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Error("Unable to generate token")
		return "", err
	}

	return signedToken, nil
}

func GetSecretKey() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}
