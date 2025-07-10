package models

import (
	"backend-cinemax/config"
	"context"
	"encoding/json"
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
	ID          int       `json:"id"`
	BackdropImg string    `json:"backdrop_img"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Popularity  float32   `json:"popularity"`
	Duration    int       `json:"duration"`
	ReleaseDate time.Time `json:"release_date"`
	Rating      float32   `json:"rating"`
	PosterImg   string    `json:"poster_img"`
	Status      string    `json:"status"` // "now playing", "coming soon", "ended"
	Language    string    `json:"language"`
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

type MovieDetail struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Duration    int        `json:"duration"`
	ReleaseDate time.Time  `json:"release_date"`
	BackdropImg string     `json:"backdrop_img"`
	PosterImg   string     `json:"poster_img"`
	Genres      []Genre    `json:"genres"`
	Directors   []Director `json:"directors"`
	Casts       []Cast     `json:"casts"`
}

func GetNowShowingMovies(sortBy, search string, page, limit int) ([]Movies, error) {
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

func GetUpComingMovies() ([]Movies, error) {
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

func GetMovieDetail(id int) (MovieDetail, error) {
	conn, err := config.DBConnect()
	if err != nil {
		fmt.Println("DetailMovies error connect to db:", err)
		return MovieDetail{}, err
	}
	defer conn.Conn().Close(context.Background())

	query := `
    SELECT 
        m.id, m.title, m.description, m.duration,
        m.release_date, m.backdrop_img, m.poster_img,
        (
            SELECT json_agg(json_build_object('id', g.id, 'name', g.name))
            FROM genres g
            JOIN movies_genres mg ON g.id = mg.genre_id
            WHERE mg.movie_id = m.id
        ) AS genres,
        (
            SELECT json_agg(json_build_object('id', d.id, 'name', d.name))
            FROM directors d
            JOIN movies_directors md ON d.id = md.director_id
            WHERE md.movie_id = m.id
        ) AS directors,
        (
            SELECT json_agg(json_build_object(
                'id', c.id,
                'actor_name', c.actor_name,
                'character_name', mc.character_name
            ))
            FROM casts c
            JOIN movies_casts mc ON c.id = mc.cast_id
            WHERE mc.movie_id = m.id
        ) AS casts
    FROM movies m
    WHERE m.id = $1
`

	var movie MovieDetail
	var genresJSON, directorsJSON, castsJSON []byte

	err = conn.QueryRow(context.Background(), query, id).Scan(
		&movie.ID, &movie.Title, &movie.Description, &movie.Duration,
		&movie.ReleaseDate, &movie.BackdropImg, &movie.PosterImg,
		&genresJSON, &directorsJSON, &castsJSON,
	)

	// Unmarshal JSON ke struct
	json.Unmarshal(genresJSON, &movie.Genres)
	json.Unmarshal(directorsJSON, &movie.Directors)
	json.Unmarshal(castsJSON, &movie.Casts)

	return movie, nil
}
