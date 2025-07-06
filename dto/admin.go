package dto

import "time"

type CastRequest struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type MoviesRequest struct {
	ID          int     `json:"id"`
	BackdropImg string     `json:"backdrop_img"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Popularity  float32    `json:"popularity"`
	Duration    int        `json:"duration"`
	ReleaseDate string     `json:"release_date"`
	Rating      float32    `json:"rating"`
	PosterImg   string     `json:"poster_img"`
	Status      string     `json:"status"` // "now playing", "coming soon", "ended"
	Language    string     `json:"language"`
	Genres      []string   `json:"genres"`
	Casts       [][]string `json:"casts"`
	Directors   string     `json:"directors"`
}

type MoviesResponse struct {
	ID          int64     		`json:"id"`
	BackdropImg string    		`json:"backdrop_img"`
	Title       string    		`json:"title"`
	Description string    		`json:"description"`
	Popularity  float32   		`json:"popularity"`
	Duration    int       		`json:"duration"`
	ReleaseDate time.Time 		`json:"release_date"`
	Rating      float32   		`json:"rating"`
	PosterImg   string    		`json:"poster_img"`
	Status      string    		`json:"status"` // "now playing", "coming soon", "ended"
	Language    string    		`json:"language"`
	Genres      []string  		`json:"genres"`
	Casts       [][]string 		`json:"casts"`
	Directors   string    		`json:"directors"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}