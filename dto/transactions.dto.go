package dto

import "github.com/google/uuid"

type CreateTransactionRequest struct {
	UserId        uuid.UUID `json:"user_id" db:"user_id"`
	CinemaId      int       `json:"cinema_id" db:"cinema_id"`
	MovieId       int       `json:"movie_id" db:"movie_id"`
	PaymentMethodId int     `json:"payment_methods_id" db:"payment_methods_id"`
	TimeBooking   string    `json:"time_booking" db:"time_booking"`
	DateBooking   string    `json:"date_booking" db:"date_booking"`
	TotalPrice    int   		`json:"total_price" db:"total_price"`
	Status        string    `json:"status" db:"status"`
	Location      string    `json:"location" db:"location"`
	Seats         []string  `json:"seats" db:"seats"`
}
