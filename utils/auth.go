package utils

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/smtp"
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

func SendEmailOTP(emailRecipient string, body string) {
	// redisClient := config.RedisConnect()
	// decoded, err := redisClient.Get(context.Background(), "users").Result()
	
	// otp := models.OTPRequest{}
	// err = json.Unmarshal([]byte(decoded), &otp)
	// if err != nil {
		
	// }

	from := "wachingcinemax@gmail.com"
	pass := "cinemax123."
	to := emailRecipient

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		"This is your OTP: " + body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	
	log.Print("sent, visit " + emailRecipient)
}