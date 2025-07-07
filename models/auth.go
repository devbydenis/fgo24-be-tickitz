package models

import (
	"backend-cinemax/config"
	c "backend-cinemax/config"
	"backend-cinemax/dto"
	u "backend-cinemax/utils"
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type IsEmailExistType struct {
	Email string `json:"email"`
}

func IsEmailExist(email string) bool {
	// conncect to db
	conn, err := c.DBConnect()
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

	// check if user exist
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

	//hash password
	hashedPassword, err := u.HashPassword(password)
	if err != nil {
		fmt.Println("InsertUserToDB error hash password:", err)
		return err
	}

	trx, err := conn.Begin(context.Background())
	if err != nil {
		fmt.Println("InsertUserToDB error begin transaction:", err)
		return err
	}
	defer trx.Rollback(context.Background())

	var userId uuid.UUID
	err = trx.QueryRow(context.Background(), `
		INSERT INTO users (id, email, password)
		VALUES ($1, $2, $3)
		RETURNING id;
`, userUUID, email, hashedPassword,
	).Scan(&userId)
	if err != nil && err != pgx.ErrNoRows {
		fmt.Println("InsertUserToDB error query row:", err)
		return err
	}

	_, err = trx.Exec(context.Background(), `
		INSERT INTO profiles (user_id) VALUES ($1);
`, userId,
	)
	if err != nil {
		fmt.Println("InsertUserToDB error insert profile:", err)
		return err
	}

	err = trx.Commit(context.Background())
	if err != nil {
		fmt.Println("InsertUserToDB error commit transaction:", err)
		return err
	}
	
	fmt.Println("InsertUserToDB success insert user and profile")
	return err
}

func MatchUserInDB(email string, password string) ([]dto.LoginRequest, error) {
	// conncect to db
	conn, err := c.DBConnect()
	if err != nil {
		fmt.Println("MatchUserInDB error connet to db:", err)
		return nil, err
	}

	// jangan lupa tutup kalo udah selesai
	defer func(){
		conn.Conn().Close(context.Background())
	}()

	// // hash input password
	// hashedPassword, err := u.HashPassword(password)
	// if err != nil {
	// 	fmt.Println("MatchUserInDB error hash password:", err)
	// 	return nil, err
	// }

	// check if email exist
	rows, err := conn.Query(context.Background(), "SELECT id, email, password FROM users WHERE email = $1", email)
	if err != nil {
		fmt.Println("MatchUserInDB error query:", err)
		return nil, err
	}

	// collect row and map to struxt
	users, err := pgx.CollectRows[dto.LoginRequest](rows, pgx.RowToStructByName)
	if err != nil {
		fmt.Println("MatchUserInDB error collect row:", err)
		return nil, err
	}
	
	// check if user exist
	if len(users) == 0 {
		return nil, fmt.Errorf("user not found")
	}
	
	// compare passwrod from user input with password from db
	if !u.VerifyHashPassword(password, users[0].Password) {
		return nil, fmt.Errorf("invalid password")
	}

	return users, nil
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

func VerifyOTP(reqOTP string) (bool, error) {
	if reqOTP == "" {
		return false, fmt.Errorf("OTP can't be empty")
	}

	redisClient := config.RedisConnect()
	decoded, err := redisClient.Get(context.Background(), "users").Result()
	if err != nil {
		fmt.Println("failed to get otp cahce from redis:", err)
		return false, err
	}

	otp := dto.OTPRequest{}
	err = json.Unmarshal([]byte(decoded), &otp)

	if reqOTP != otp.OTP {
		return false, fmt.Errorf("cache otp and request otp doesnt match")
	}

	return true, nil
}