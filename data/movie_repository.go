package data

import (
	"database/sql"
	"errors"
	"fm-go-vanillajs-movies/logger"
	"fm-go-vanillajs-movies/models"

	_ "github.com/lib/pq"
)

type MovieRepository struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewMovieRepository(db *sql.DB, log *logger.Logger) (*MovieRepository, error) {
	return &MovieRepository{
		db:     db,
		logger: log,
	}, nil
}

const defaultLimit = 20

func (r *MovieRepository) GetTopMovies() ([]models.Movie, error) {
	// fetch movies
	query := `
                SELECT id, tmdb_id, title, tagline, release_year, overview, score,
                popularity, language, poster_url, trailer_url
                FROM movies
                ORDER BY popularity DESC
                LIMIT $1
        `

	return r.getMovies(query)
}

func (r *MovieRepository) GetRandomMovies() ([]models.Movie, error) {
	query := `
                SELECT id, tmdb_id, title, tagline, release_year, overview, score,
                popularity, language, poster_url, trailer_url
                FROM movies
                ORDER BY random() DESC
                LIMIT $1
        `

	return r.getMovies(query)
}

func (r *MovieRepository) GetMovieById(id int) (models.Movie, error) {
	return models.Movie{}, nil
}

func (r *MovieRepository) SearchMoviesByName(name string) ([]models.Movie, error) {
	return make([]models.Movie, 2), nil
}

func (r *MovieRepository) GetAllGenres() ([]models.Genre, error) {
	return make([]models.Genre, 1), nil
}

func (r *MovieRepository) getMovies(query string) ([]models.Movie, error) {
	rows, err := r.db.Query(query, defaultLimit)
	if err != nil {
		r.logger.Error("Failed to query movies", err)
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var m models.Movie
		var tmdbID sql.NullInt32
		var tagline, overview, language, posterURL, trailerURL sql.NullString
		var score, popularity sql.NullFloat64
		var releaseYear sql.NullInt32

		if err := rows.Scan(
			&m.ID, &tmdbID, &m.Title, &tagline, &releaseYear,
			&overview, &score, &popularity, &language,
			&posterURL, &trailerURL,
		); err != nil {
			r.logger.Error("Failed to scan movie row", err)
			return nil, err
		}

		// Convert nullable fields to pointers
		if tmdbID.Valid {
			m.TMDB_ID = int(tmdbID.Int32)
		}
		if tagline.Valid {
			m.Tagline = tagline.String
		}
		if releaseYear.Valid {
			m.ReleaseYear = int(releaseYear.Int32)
		}
		if overview.Valid {
			m.Overview = &overview.String
		}
		if score.Valid {
			scoreFloat := float32(score.Float64)
			m.Score = &scoreFloat
		}
		if popularity.Valid {
			popularityFloat := float32(popularity.Float64)
			m.Popularity = &popularityFloat
		}
		if language.Valid {
			m.Language = &language.String
		}
		if posterURL.Valid {
			m.PosterURL = &posterURL.String
		}
		if trailerURL.Valid {
			m.TrailerURL = &trailerURL.String
		}

		movies = append(movies, m)
	}

	if err = rows.Err(); err != nil {
		r.logger.Error("Error iterating movie rows", err)
		return nil, err
	}

	return movies, nil
}

var (
	ErrMovieNotFound = errors.New("movie not found")
)
