package tools

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID int, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// ambil secret dari env variable JWT_SECRET
	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	return token.SignedString(jwtKey)
}
