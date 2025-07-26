package models

import (
	"backend-cinemax/config"
	"backend-cinemax/dto"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func CreateTransaction(req dto.CreateTransactionRequest) error {
		// connect to db
	conn, err := config.DBConnect()
	if err != nil {
		return err
	}
	defer func() {
		conn.Conn().Close(context.Background())
	}()

	// start trx
	trx, err := conn.Begin(context.Background())
	if err != nil {
		fmt.Println("CreateBooking error begin transaction:", err)
		return err
	}
	defer trx.Rollback(context.Background())

	// insert to transactions table
	var transaction struct {
		Id uuid.UUID
	}
	queryToTransactionsTable := `
		INSERT INTO transactions (
			user_id, 
			cinema_id, 
			movie_id, 
			payment_methods_id,
			time_booking,
			date_booking, 
			total_price
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
	`
	err = conn.QueryRow(context.Background(), queryToTransactionsTable,
		req.UserId,
		req.CinemaId,
		req.MovieId,
		req.PaymentMethodId,
		req.TimeBooking, 
		req.DateBooking, 
		req.TotalPrice, 
	).Scan(&transaction.Id)

	if err != nil {
		fmt.Println("CreateBooking error exec row:", err)
		return err
	}
	
	// insert to transaction_detail table
	queryToTransactionDetail:=`
		INSERT INTO transaction_detail (transaction_id, seat)
		VALUES ($1, $2);
	`
	for _, seat := range req.Seats {
		_, err = conn.Exec(context.Background(), queryToTransactionDetail, transaction.Id, seat)
		if err != nil {
			fmt.Println("Create transaction detail error exec row:", err)
			return err
		}
	}

	// insert to transaction_history table
	queryToTransactionHistory := `
		INSERT INTO transaction_history (transaction_id, status)
		VALUES ($1, $2);
	`
	_, err = conn.Exec(context.Background(), queryToTransactionHistory, transaction.Id, req.Status)
	if err != nil {
		fmt.Println("CreateBooking error exec row:", err)
		return err
	}

	// commit trx
	err = trx.Commit(context.Background())
	if err != nil {
		fmt.Println("CreateBooking error commit transaction:", err)
		return err
	}

	return nil
}