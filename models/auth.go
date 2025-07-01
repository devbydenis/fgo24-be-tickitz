package models

import (
	c "backend-cinemax/config"
	"context"

	"github.com/google/uuid"
)

type RegisterRequest struct {
	Email 			string `json:"email" form:"email"`
	Password 		string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type LoginRequest struct {
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