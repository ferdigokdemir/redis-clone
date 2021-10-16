package main

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	// Set Environment Variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	os.Setenv("PORT", port)
	os.Setenv("DATABASE_SAVE_INTERVAL_SECONDS", "60")
	os.Setenv("ENABLE_CORS", "true")
	logger = log.New(os.Stdout, "http: ", log.LstdFlags)
}

func main() {
	LoadDatabase()
	go SaveDatabase()
	StartHttpRouter()
}
