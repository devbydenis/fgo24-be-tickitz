package utils

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func GenerateToken(id uuid.UUID, email string) (string, error) {
	godotenv.Load()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    id,
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte("SECRET_KEY"))
	if err != nil {
		fmt.Println("GenerateToken error:", err)
		return "", err
	}

	return tokenString, nil
}

func GenerateOTP() string {
	result := 0
	for {
		randomNumber := rand.IntN(9999)
		if len(strconv.Itoa(randomNumber)) == 4 {
			result = randomNumber
			break
		}
	}
	return strconv.Itoa(result)
}