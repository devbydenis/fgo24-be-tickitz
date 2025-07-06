package dto

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
