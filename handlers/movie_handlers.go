package handlers

import (
	"encoding/json"
	"fm-go-vanillajs-movies/data"
	"fm-go-vanillajs-movies/logger"
	"fm-go-vanillajs-movies/models"
	"fmt"
	"net/http"
	"strconv"
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
		mh.logger.Error("Get top movies error", err)
		http.Error(w, fmt.Sprintf("Internal error getting top movies: %v", err), 500)
		return
	}

	mh.writeJSONResponse(w, movies)
}

func (mh *MovieHandlers) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := mh.storage.GetRandomMovies()
	if err != nil {
		mh.logger.Error("Get random movies: ", err)
		http.Error(w, "Error getting random movies", 500)
		return
	}

	mh.writeJSONResponse(w, movies)
}

func (mh *MovieHandlers) writeJSONResponse(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&data)
	if err != nil {
		mh.logger.Error("writeJSONResponse: ", err)
		http.Error(w, "error: internal error", http.StatusInternalServerError)
		return err
	}
	return nil
}

func (mh *MovieHandlers) handleStorageError(w http.ResponseWriter, err error, context string) bool {
	if err != nil {
		if err == data.ErrMovieNotFound {
			http.Error(w, context, http.StatusNotFound)
			return true
		}
		mh.logger.Error(context, err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return true
	}
	return false
}

func (mh *MovieHandlers) parseID(w http.ResponseWriter, idStr string) (int, bool) {
	id, err := strconv.Atoi((idStr))
	if err != nil {
		mh.logger.Error("Invalid id format", err)
		http.Error(w, "invalid id", http.StatusBadRequest)
		return 0, false
	}
	return id, true
}

func (mh *MovieHandlers) SearchMovies(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	order := r.URL.Query().Get("order")
	genreStr := r.URL.Query().Get("genre")

	var genre *int
	if genreStr != "" {
		genreInt, ok := mh.parseID(w, genreStr)
		if !ok {
			return
		}
		genre = &genreInt
	}

	var movies []models.Movie
	var err error
	if query == "" {
		if mh.writeJSONResponse(w, make([]models.Movie, 0)) == nil {
			mh.logger.Info("Served empty slice because query is empty too.")
			return
		}
		return
	}
	if query != "" {
		movies, err = mh.storage.SearchMoviesByName(query, order, genre)
	}
	fmt.Println("movies search: ", movies)
	if mh.handleStorageError(w, err, "failed to get movies") {
		return
	}
	if mh.writeJSONResponse(w, movies) == nil {
		mh.logger.Info("successfully served movies")
	}
}

func (mh *MovieHandlers) GetMovie(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/movies/"):]
	id, ok := mh.parseID(w, idStr)
	if !ok {
		return
	}

	movie, err := mh.storage.GetMovieById(id)
	if mh.handleStorageError(w, err, "failed to get movie by id") {
		return
	}
	if mh.writeJSONResponse(w, movie) == nil {
		mh.logger.Info("Successfully server moview with ID: " + idStr)
	}
}

func (mh *MovieHandlers) GetGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := mh.storage.GetAllGenres()
	if mh.handleStorageError(w, err, "failed to get genres") {
		return
	}
	if mh.writeJSONResponse(w, genres) == nil {
		mh.logger.Info("Successfully served genres")
	}
}
