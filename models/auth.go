package models

import (
	"backend-cinemax/config"
	c "backend-cinemax/config"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type RegisterRequest struct {
	Email 			string `json:"email" form:"email"`
	Password 		string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type LoginRequest struct {
	ID 		  	uuid.UUID `json:"id" form:"id"`
	Email 	  string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
}

type ForgotPasswordRequest struct {
	Email string `form:"email" json:"email"`
}

type ChangePasswordRequest struct {
	Email 				string `json:"email" form:"email"`
	NewPassword 		string `json:"new_password" form:"new_password"`
	ConfirmNewPassword 	string `json:"confirm_new_password" form:"confirm_new_password"`
}

type IsEmailExistType struct {
	Email string `json:"email"`
}

type OTPRequest struct {
	Email string `json:"email" form:"email"`
	OTP string `json:"otp" form:"otp"`
}

type VerifyOTP struct {
	OTP string `json:"otp" form:"otp"`
}

func IsEmailExist(email string) bool {
	// conncect to db
	conn, err := config.DBConnect()
	if err != nil {
		fmt.Println("IsEmailExist error connet to db:", err)
		return false
	}
	// jangan lupa tutup kalo udah selesai
	defer func() {
		conn.Conn().Close(context.Background())
	}()

	// check if email exist
	rows, err := conn.Query(context.Background(), `SELECT email FROM users WHERE email = $1`, email,)
	if err != nil {
		fmt.Println("IsEmailExist error query:", err)
		return false
	}

	// collect row and map to struxt
	users, err := pgx.CollectRows[IsEmailExistType](rows, pgx.RowToStructByName)
	if err != nil {
		fmt.Println("IsEmailExist error collect row:", err)
		return false
	}

	fmt.Println("IsEmailExist users:", users)
	if len(users) > 0 {
		return true
	}

	return false
}

func InsertUserToDB(email string, password string, userUUID uuid.UUID) error {
	conn, err := c.DBConnect()
	if err != nil {
		return err
	}
	defer func(){
		conn.Conn().Close(context.Background())
	}()

	_, err = conn.Exec(context.Background(), `
		INSERT INTO users (id, email, password) VALUES ($1, $2, $3);
	`, userUUID, email, password,
)

	_, err = conn.Exec(context.Background(), `
		INSERT INTO profiles (user_id) VALUES ($1);
	`, userUUID,
)

	return err
}

func MatchUserInDatabase(email string, password string) bool {
	// conncect to db
	conn, err := c.DBConnect()
	if err != nil {
		fmt.Println("MatchUserInDatabase error connet to db:", err)
		return false
	}

	// jangan lupa tutup kalo udah selesai
	defer func(){
		conn.Conn().Close(context.Background())
	}()

	// check if email exist
	rows, err := conn.Query(context.Background(), "SELECT id, email, password FROM users WHERE email = $1 AND password = $2", email, password,)
	if err != nil {
		fmt.Println("MatchUserInDatabase error query:", err)
		return false
	}

	// collect row and map to struxt
	users, err := pgx.CollectRows[LoginRequest](rows, pgx.RowToStructByName)
	if err != nil {
		fmt.Println("MatchUserInDatabase error collect row:", err)
		return false
	}
	
	fmt.Println("MatchUserInDatabase users:", users)
	if len(users) == 0 {
		return false
	}

	return true
}

func UpdateUserPassword(email string, newPassword string) error {
	conn, err := c.DBConnect()
	if err != nil {
		return err
	}
	defer func(){
		conn.Conn().Close(context.Background())
	}()

	_, err = conn.Exec(context.Background(), `
		UPDATE users SET password = $1 WHERE email = $2;
	`, newPassword, email,
	)

	return err
}