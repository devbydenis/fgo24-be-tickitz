package models

import (
	"backend-cinemax/config"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type User struct {
	ID string `json:"id" form:"id"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Username string `json:"username" form:"username"`
	Role string `json:"role" form:"role"`	
}

type IsEmailExistType struct {
	Email string `json:"email"`
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
	rows, err := conn.Query(context.Background(), "SELECT email FROM users WHERE email = $1", email,)
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