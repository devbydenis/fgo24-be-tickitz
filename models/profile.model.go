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
	// conncect to db
	conn, err := config.DBConnect()
	if err != nil {
		return err
	}
	defer func(){
		conn.Conn().Close(context.Background())
	}()

	// set default value if request is empty
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

func UploadPhoto(userId, fileName string) error {
	// conncect to db
	conn, err := config.DBConnect()
	if err != nil {
		return err
	}
	defer func(){
		conn.Conn().Close(context.Background())
	}()
	
	fmt.Println(fileName)

	// set default value if request is empty
	query := `
		UPDATE profiles 
		SET profile_picture = $1 
		WHERE user_id = $2;
		`

	_, err = conn.Exec(
			context.Background(), 
			query, 
			fileName, 
			userId,
	)

	return err
}

func GetHistory(userId string) ([]dto.GetHistoryResponse, error) {
		// conncect to db
	conn, err := config.DBConnect()
	if err != nil {
		return []dto.GetHistoryResponse{}, err
	}
	defer func(){
		conn.Conn().Close(context.Background())
	}()

	query := `
		SELECT 
			t.id AS transaction_id,
			c.name AS cinema_name,
			t.date_booking,
			t.time_booking,
			m.title AS movie_name,
			(
				SELECT json_agg(td.seat) AS seats
				FROM transaction_detail td
				JOIN transactions t ON td.transaction_id = t.id
			),
			t.total_price,
			th.status
		FROM 
			transactions t
		JOIN transaction_detail td ON t.id = td.transaction_id
		JOIN transaction_history th ON t.id = th.transaction_id
		JOIN cinemas c ON t.cinema_id = c.id
		JOIN movies m ON t.movie_id = m.id
		WHERE 
			t.user_id = $1
		GROUP BY 
			t.id, t.user_id, c.name, t.date_booking, t.time_booking, m.title, t.total_price, th.status;
	`

	rows, err := conn.Query(context.Background(), query, userId)
	if err != nil {
		fmt.Println("GetHistory error exec row:", err)
		return []dto.GetHistoryResponse{}, err
	}
	defer rows.Close()

	var histories []dto.GetHistoryResponse
	for rows.Next() {
		var history dto.GetHistoryResponse
		err = rows.Scan(
			&history.TransactionId,
			&history.CinemaName,
			&history.DateBooking,
			&history.TimeBooking,
			&history.MovieName,
			&history.Seats,
			&history.TotalPrice,
			&history.Status,
		)

		if err != nil {
			fmt.Println("GetHistory error scan row:", err)
			return []dto.GetHistoryResponse{}, err
		}
		
		histories = append(histories, history)
		fmt.Println("GetHistory: ", histories)
	}

	return histories, nil
}