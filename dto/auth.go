package dto

import "github.com/google/uuid"

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


type OTPRequest struct {
	Email string `json:"email" form:"email"`
	OTP string `json:"otp" form:"otp"`
}

type VerifyOTP struct {
	OTP string `json:"otp" form:"otp"`
}