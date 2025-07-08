package models

import (
	"backend-cinemax/config"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type Genre struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name"`
}

type Director struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Cast struct {
	ID            int    `json:"id" db:"id"`
	ActorName     string `json:"actor_name" db:"actor_name"`
	CharacterName string `json:"character_name" db:"character_name"`
}

type Movies struct {
	ID          int     `json:"id"`
	BackdropImg string  `json:"backdrop_img"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Popularity  float32 `json:"popularity"`
	Duration    int     `json:"duration"`
	ReleaseDate time.Time  `json:"release_date"`
	Rating      float32 `json:"rating"`
	PosterImg   string  `json:"poster_img"`
	Status      string  `json:"status"` // "now playing", "coming soon", "ended"
	Language    string  `json:"language"`
	// CreatedAt   string  `json:"created_at"`
	// UpdatedAt   string  `json:"updated_at"`
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

func NowShowingMovies(sortBy, search string, page, limit int) ([]Movies, error) {
	conn, err := config.DBConnect()
	if err != nil {
		fmt.Println("NowShowingMovies error connect to db:", err)
		return []Movies{}, err
	}

	defer func() {
		conn.Conn().Close(context.Background())
	}()

	fmt.Println("sortBy:", sortBy, "search:", search, "page:", page, "limit:", limit)

	if page == 0 {
		page = 1
	}

	// Whitelist untuk kolom yang boleh di-sort
	validSortColumns := map[string]bool{
		"popularity":   true,
		"release_date": true,
		"rating":       true,
		"title":        true,
	}

	if !validSortColumns[sortBy] {
		sortBy = "popularity"
	}

	// Query dengan OR condition untuk handle empty search
	query := fmt.Sprintf(`
		SELECT 
			id, 
			title, 
			description, 
			popularity, 
			status, 
			duration, 
			release_date, 
			rating, 
			poster_img, 
			backdrop_img, 
			language 
		FROM movies 
		WHERE status = 'now playing'
			AND ($1 = '' OR title ILIKE $1)
		ORDER BY %s DESC
		LIMIT $2
		OFFSET $3
	`, sortBy)

	// Format search parameter
	searchParam := ""
	if search != "" {
		searchParam = "%" + search + "%"
	}

	rows, err := conn.Query(
		context.Background(),
		query,
		searchParam,
		limit,
		(page-1)*limit,
	)
	if err != nil {
		fmt.Println("NowShowingMovies error query:", err)
		return []Movies{}, err
	}

	movies, err := pgx.CollectRows[Movies](rows, pgx.RowToStructByName)
	if err != nil {
		fmt.Println("NowShowingMovies error collect row:", err)
		return []Movies{}, err
	}

	return movies, nil
}


func UpComingMovies() ([]Movies, error) {
	conn, err := config.DBConnect()
	if err != nil {
		fmt.Println("UpComingMovies error connect to db:", err)
		return []Movies{}, err
	}
	defer conn.Conn().Close(context.Background())

	query := `SELECT id, title, description, popularity, status, duration, release_date, rating, poster_img, backdrop_img, language FROM movies WHERE status = 'coming soon'`

	rows, err := conn.Query(
		context.Background(), query)
	if err != nil {
		fmt.Println("UpComingMovies error query:", err)
		return []Movies{}, err
	}

	movies, err := pgx.CollectRows[Movies](rows, pgx.RowToStructByName)
	if err != nil {
		fmt.Println("UpComingMovies error collect row:", err)
		return []Movies{}, err
	}

	return movies, nil
}