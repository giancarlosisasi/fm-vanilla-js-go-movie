package handlers

import (
	"encoding/json"
	"fm-go-vanillajs-movies/data"
	"fm-go-vanillajs-movies/logger"
	"fm-go-vanillajs-movies/models"
	"fmt"
	"net/http"
)

type MovieHandlers struct {
	logger  *logger.Logger
	storage data.MovieStorage
}

// factory to create a new handlers
func NewMovieHandlers(logger *logger.Logger, storage data.MovieStorage) *MovieHandlers {
	return &MovieHandlers{logger: logger, storage: storage}
}

func (mh *MovieHandlers) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := mh.storage.GetTopMovies()
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal error getting top movies: %v", err), 500)
		mh.logger.Error("Get top movies error", err)
		return
	}

	mh.writeJSONResponse(w, movies)
}

func (mh *MovieHandlers) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID:          1,
			TMDB_ID:     181,
			Title:       "The hacker",
			ReleaseYear: 1984,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 1, FirstName: "Max"}},
		},
	}

	mh.writeJSONResponse(w, movies)
}

func (mh *MovieHandlers) writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&data)
	if err != nil {
		mh.logger.Error("writeJSONResponse: ", err)
		http.Error(w, "error: internal error", http.StatusInternalServerError)
	}
}
