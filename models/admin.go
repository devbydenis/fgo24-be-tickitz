package models

import (
	"backend-cinemax/config"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type CastRequest struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type MoviesRequest struct {
	// ID          int     `json:"id"`
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

func InsertToMovieTable(trx pgx.Tx, movie MoviesRequest) (int64, error) {
	// prepare the query
	queryMovies := `	INSERT INTO movies (backdrop_img, title, description, popularity, duration, release_date,
		rating, poster_img, status, language, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW(), NOW())
		RETURNING id;`

	// insert to movies table
	var movieId int64
	err := trx.QueryRow(
		context.Background(),
		queryMovies,
		movie.BackdropImg,
		movie.Title,
		movie.Description,
		movie.Popularity,
		movie.Duration,
		movie.ReleaseDate,
		movie.Rating,
		movie.PosterImg,
		movie.Status,
		movie.Language,
	).Scan(&movieId)
	if err != nil {
		fmt.Println("InsertMovieToDB error query row:", err)
		return 0, err
	}

	return movieId, nil
}

func InsertToGenreTable(trx pgx.Tx, genres []string, movieId int64) error {
	for _, genre := range genres {
		// check each id genres input
		var genreId int64
		err := trx.QueryRow(
			context.Background(),
			`SELECT id FROM genres WHERE name = $1`,
			genre).Scan(&genreId)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				// if genre does not exist, insert new genre
				err = trx.QueryRow(
					context.Background(),
					`INSERT INTO genres (name) VALUES($1) RETURNING id`,
					genre).Scan(&genreId)
				if err != nil {
					fmt.Println("InsertGenreToDB error insert new genre:", err)
					return err
				}
				fmt.Println("created new genre:", genre, "with id:", genreId)
			} else {
				fmt.Println("InsertGenreToDB error query genre row:", err)
				return err
			}
		}

		_, err = trx.Exec(
			context.Background(),
			`INSERT INTO movies_genres (movie_id, genre_id) VALUES ($1, $2)`,
			movieId, genreId)
		if err != nil {
			fmt.Println("InsertGenreToDB error insert genre:", err)
			return err
		}
	}
	
	return nil
}

func InsertToCastsTable(trx pgx.Tx, casts [][]string, movieId int64) error {
	// iterate each cast
	for _, cast := range casts {
		if len(cast) < 2 {
			continue
		}
		actorName := cast[0]
		characterName := cast[1]

		// check if actor already exists
		var castId int64
		err := trx.QueryRow(
			context.Background(),
			`SELECT id FROM casts WHERE actor_name = $1`,
			actorName).Scan(&castId)
		if err != nil {
			// if actir does not exist, insert new actor
			if errors.Is(err, pgx.ErrNoRows) {
				err = trx.QueryRow(
					context.Background(),
					`INSERT INTO casts (actor_name) VALUES($1) RETURNING id`,
					actorName).Scan(&castId)
				if err != nil {
					fmt.Println("InsertCastsToDB error insert new cast:", err)
					return err
				}
			} else {
				fmt.Println("InsertCastsToDB error query cast row:", err)
				return err
			}
		}

		// if actor exists, use the existing id instead
		_, err = trx.Exec(context.Background(), `INSERT INTO movies_casts (movie_id, cast_id, character_name) VALUES ($1, $2, $3)`, movieId, castId, characterName)
		if err != nil {
			fmt.Println("InsertCastsToDB error insert cast:", err)
			return err
		}
	}

	return nil
}

func InsertToDirectorsTable(trx pgx.Tx, directors string, movieId int64) error {
	// check if director already exists
	var directorId int64
	err := trx.QueryRow(
		context.Background(),
		`SELECT id FROM directors WHERE name = $1`,
		directors).Scan(&directorId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = trx.QueryRow(
				context.Background(),
				`INSERT INTO directors (name) VALUES($1) RETURNING id`,
				directors).Scan(&directorId)
			if err != nil {
				fmt.Println("InsertDirectorsToDB error insert new director:", err)
				return err
			}
		} else {
			fmt.Println("InsertDirectorsToDB error query director row:", err)
			return err
		}
	}

	// if director exists, use the existing id instead
	_, err = trx.Exec(context.Background(), `INSERT INTO movies_directors (movie_id, director_id) VALUES ($1, $2)`, movieId, directorId)
	if err != nil {
		fmt.Println("InsertDirectorsToDB error insert director:", err)
		return err
	}

	return nil
}

func CreateMovieWithAllRelations(req MoviesRequest) error {
	// connect to db
	conn, err := config.DBConnect()
	if err != nil {
		return err
	}
	defer func() {
		conn.Conn().Close(context.Background())
	}()

	trx, err := conn.Begin(context.Background())
	if err != nil {
		fmt.Println("CreateMovieWithAllRelations error begin transaction:", err)
		return err
	}
	defer trx.Rollback(context.Background())

	movieId, err := InsertToMovieTable(trx, req)
	if err != nil {
		fmt.Println("CreateMovieWithAllRelations error insert movie:", err)
		return err
	}

	err = InsertToGenreTable(trx,req.Genres, movieId)
	if err != nil {
		fmt.Println("CreateMovieWithAllRelations error insert genres:", err)
		return err
	}

	err = InsertToCastsTable(trx,req.Casts, movieId)
	if err != nil {
		fmt.Println("CreateMovieWithAllRelations error insert casts:", err)
		return err
	}

	err = InsertToDirectorsTable(trx, req.Directors, movieId)
	if err != nil {
		fmt.Println("CreateMovieWithAllRelations error insert directors:", err)
		return err
	}

	err = trx.Commit(context.Background())
	if err != nil {
		fmt.Println("CreateMovieWithAllRelations error commit transaction:", err)
		return err
	}

	return nil
}