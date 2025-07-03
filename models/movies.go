package models

import (
	"backend-cinemax/config"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Movies struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	BackdropImg string  `json:"backdrop_img"`
	Title       string  `json:"title"`
	Synopsis    string  `json:"synopsis"`
	Popularity  float32 `json:"popularity"`
	Duration    int     `json:"duration"`
	ReleaseDate string  `json:"release_date"`
	Rating      float32 `json:"rating"`
	PosterImg   string  `json:"poster_img"`
	Status      string  `json:"status"` // "now playing", "coming soon", "ended"
	Language    string  `json:"language"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type NowShowingMoviesRequest struct {
	Limit  int    `json:"limit" form:"limit"`
	Page   int    `json:"page" form:"page"`
	SortBy string `json:"sort_by" form:"sort_by"` // e.g., "popularity", "release_date"
	Order  string `json:"order" form:"order"`     // e.g., "asc", "desc"
	Search string `json:"search" form:"search"`
}
type NowShowingMoviesResponse struct {
	Page   int      `json:"page"`
	Limit  int      `json:"limit"`
	Total  int      `json:"total"`
	Movies []Movies `json:"movies"`
}

func NowShowingMovies() ([]Movies, error) {	// req masih belum dipake
	// conncect to db
	conn, err := config.DBConnect()
	if err != nil {
		fmt.Println("NowShowingMovies error connect to db:", err)
		return []Movies{}, err
	}

	// jangan lupa tutup kalo udah selesai
	defer func() {
		conn.Conn().Close(context.Background())
	}()

	rows, err := conn.Query(
		context.Background(),
		`
			SELECT * FROM movies WHERE status = 'now playing'
		`,
	)
	if err != nil {
		fmt.Println("NowShowingMovies error query:", err)
		return []Movies{}, err
	}

	// collect row and map to struct
	// pgx.CollectRows is used to collect rows from the query result into a slice
	movies, err := pgx.CollectRows[Movies](rows, pgx.RowToStructByName)
	if err != nil {
		fmt.Println("NowShowingMovies error collect row:", err)
		return []Movies{}, err
	}
	return movies, nil
}
