package utils

import (
	"fmt"
	"math/rand/v2"
	// "net/smtp"
	"os"

	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func GenerateToken(id uuid.UUID, email string) (string, error) {
	godotenv.Load()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    id,
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
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

func SendEmailOTP(emailRecipient, otp string) error {
	godotenv.Load()

	const CONFIG_SMTP_PORT = 587
	const CONFIG_SENDER_NAME = "wachingcinemax@gmail.com"

	mailer := gomail.NewMessage()
    mailer.SetHeader("From", CONFIG_SENDER_NAME)
    mailer.SetHeader("To", emailRecipient, "rahmadidenis@gmail.com")
    mailer.SetHeader("Subject", "OTP Verification!")
    mailer.SetBody("text/html", otp)

	dialer := gomail.NewDialer(
			os.Getenv("CONFIG_SMTP_HOST"),
			CONFIG_SMTP_PORT,
			os.Getenv("CONFIG_AUTH_EMAIL"),
			os.Getenv("CONFIG_AUTH_PASSWORD"),
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
			return err
	}

	return nil

	// // SMTP
	// 	smtpHost := os.Getenv("CONFIG_SMTP_HOST")
  //   smtpPort := os.Getenv("CONFIG_SMTP_PORT")
  //   senderEmail := os.Getenv("CONFIG_AUTH_EMAIL")      
  //   senderPassword := os.Getenv("CONFIG_AUTH_PASSWORD")    

  //   auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)

  //   subject := "Subject: Your OTP Code\r\n"
  //   body := fmt.Sprintf("Your OTP code is: %s", otp)
  //   msg := []byte(subject + "\r\n" + body)

  //   err := smtp.SendMail(
  //       smtpHost+":"+smtpPort,
  //       auth,
  //       senderEmail,
	// 			[]string{emailRecipient},
  //       msg,
  //   )
  //   return err
}