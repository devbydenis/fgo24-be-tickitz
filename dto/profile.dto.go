package dto

import (
	"time"

	"github.com/google/uuid"
)

type GetProfileResponse struct {
	Email          string `json:"email" form:"email" db:"email"`
	Username       string `json:"username" form:"username" db:"username"`
	Firstname      string `json:"firstname" form:"firstname" db:"firstname"`
	Lastname       string `json:"lastname" form:"lastname" db:"lastname"`
	PhoneNumber    string `json:"phone_number" form:"phone_number" db:"phone_number"`
	Gender         string `json:"gender" form:"gender" db:"gender"`
	ProfilePicture string `json:"profile_picture" form:"profile_picture" db:"profile_picture"`
}

type UpdateProfileRequest struct {
	Username       *string `json:"username" form:"username" db:"username"`
	Firstname      *string `json:"firstname" form:"firstname" db:"firstname"`
	Lastname       *string `json:"lastname" form:"lastname" db:"lastname"`
	PhoneNumber    *string `json:"phone_number" form:"phone_number" db:"phone_number"`
	Gender         *string `json:"gender" form:"gender" db:"gender"`
	ProfilePicture *string `json:"profile_picture" form:"profile_picture" db:"profile_picture"`
}

type GetHistoryResponse struct {
	TransactionId   uuid.UUID    `json:"transaction_id" form:"transaction_id" db:"transaction_id"`
	CinemaName      string `json:"cinema_name" form:"cinema_name" db:"cinema_name"`
	DateBooking time.Time `json:"date_booking" form:"date_booking" db:"date_booking"`
	TimeBooking time.Time `json:"time_booking" form:"time_booking" db:"time_booking"`
	MovieName   string `json:"movie_name" form:"movie_name" db:"movie_name"`
	Seats       []string `json:"seats" form:"seats" db:"seats"`
	TotalPrice  int    `json:"total_price" form:"total_price" db:"total_price"`
	Status      string `json:"status" form:"status" db:"status"`
}
