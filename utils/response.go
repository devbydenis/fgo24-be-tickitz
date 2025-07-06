package utils

type Response struct {
	Status  int    `json:"status,omitempty"`
	Success bool `json:"success"`
	Message string `json:"message"`
	Errors string `json:"errors,omitempty"`
	Token string `json:"token,omitempty"`
	OTP string `json:"otp,omitempty"`
	Total int64 `json:"total,omitempty"`
	Result interface{} `json:"result,omitempty"`
}