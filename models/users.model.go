package models

import (
	
)

type User struct {
	ID string `json:"id" form:"id"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	// Username string `json:"username" form:"username"`
	Role string `json:"role" form:"role"`	
}

