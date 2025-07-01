package utils

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Errors string `json:"errors,omitempty"`
	Token string `json:"token,omitempty"`
}