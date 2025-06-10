package main

import (
	"fm-go-vanillajs-movies/logger"
	"log"
	"net/http"
)

func main() {
	// Initialize logger
	logInstance := initializeLogger()

	http.Handle("/", http.FileServer(http.Dir("public")))

	const addr = ":8080"
	logInstance.Info("Server starting on" + addr)
	err := http.ListenAndServe(addr, nil)
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
