package models

import (
	"backend-cinemax/config"
	"backend-cinemax/dto"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func GetUserByUserId(userId string) (dto.GetProfileResponse, error) {
		// conncect to db
	conn, err := config.DBConnect()
	if err != nil {
		fmt.Println("IsEmailExist error connet to db:", err)
		return dto.GetProfileResponse{}, err
	}
	// jangan lupa tutup kalo udah selesai
	defer func() {
		conn.Conn().Close(context.Background())
		}()
		
		// check if email exist
		rows, err := conn.Query(
			context.Background(),
		`
			SELECT 
				u.email, 
				COALESCE(p.username, '') as username,
				COALESCE(p.firstname, '') as firstname, 
				COALESCE(p.lastname, '') as lastname, 
				COALESCE(p.phone_number, '') as phone_number, 
				COALESCE(p.gender, '') as gender, 
				COALESCE(p.profile_picture, '') as profile_picture
			FROM users u
			LEFT JOIN profiles p ON u.id = p.user_id
			WHERE u.id = $1;
		`, 
		userId,
	)
	if err != nil {
		fmt.Println("GetUserByEmail error query:", err)
		return dto.GetProfileResponse{}, err
	}
	
	// collect row and map to struct
	users, err := pgx.CollectRows[dto.GetProfileResponse](rows, pgx.RowToStructByName)
	fmt.Println("GetUserByEmail users:", users)
	if err != nil {
		fmt.Println("GetUserByEmail error collect row:", err)
		return dto.GetProfileResponse{}, err
	}
	// check if user found
	if len(users) == 0 {
		return dto.GetProfileResponse{}, fmt.Errorf("user not found")
	}

	return users[0], nil
}

func UpdateUser(userId string, req dto.UpdateProfileRequest) error {
	conn, err := config.DBConnect()
	if err != nil {
		return err
	}
	defer func(){
		conn.Conn().Close(context.Background())
	}()

		// fmt.Println("UpdateUser req:", req.Firstname)

	query := `
		UPDATE profiles 
		SET username = COALESCE($1, username),
				firstname = COALESCE($2, firstname),
				lastname = COALESCE($3, lastname),
				phone_number = COALESCE($4, phone_number),
				gender = COALESCE($5, gender),
				profile_picture = COALESCE($6, profile_picture) 
		WHERE user_id = $7;
		`

	_, err = conn.Exec(
			context.Background(), 
			query, 
			req.Username, 
			req.Firstname, 
			req.Lastname, 
			req.PhoneNumber, 
			req.Gender, 
			req.ProfilePicture, 
			userId,
	)

	return err
}