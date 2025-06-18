package main

import (
	"database/sql"
	"fm-go-vanillajs-movies/data"
	"fm-go-vanillajs-movies/handlers"
	"fm-go-vanillajs-movies/logger"
	"log"
	"net/http"
	"os"

	// _ "github.com/golang-migrate/migrate/v4/database/postgres"
	// _ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Initialize logger
	logInstance := initializeLogger()

	// migrations
	// err := runMigrations(logInstance)
	// if err != nil {
	// 	logInstance.Error("Error to run migrations: ", err)
	// 	panic("error to run migrations")
	// }

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file was available")
	}

	// connect to the db
	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		log.Fatal("DATABASE_URL not set")
	}
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to the DB: %v", err)
	}
	defer db.Close()

	// repositories
	movieRepo, err := data.NewMovieRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Failed to initialize the movie repository: %v", err)
	}

	// handlers
	movieHandler := handlers.NewMovieHandlers(logInstance, movieRepo)

	// endpoints
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)

	// Handler for static files(frontend)
	// Must be at the end, after all other handlers
	http.Handle("/", http.FileServer(http.Dir("public")))

	const addr = ":8080"
	logInstance.Info("Server starting on" + addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		logInstance.Error("Server failed to start", err)
		log.Fatalf("Server has failed: %v", err)
	}
}

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movie-service.log")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	defer logInstance.Close()

	return logInstance
}

// func runMigrations(logger *logger.Logger) error {
// 	m, err := migrate.New(
// 		"file://migrations",
// 		"postgres://postgres:postgres@localhost:5433/movies_db?sslmode=disable",
// 	)

// 	if err != nil {
// 		logger.Error("Error to initialize the migrate tool: ", err)
// 		return err
// 	}

// 	if err := m.Up(); err != nil {
// 		logger.Error("Error to run the migrations up: ", err)
// 		return err
// 	}

// 	return nil
// }
