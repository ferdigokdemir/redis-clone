package main

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	// Set Environment Variables
	os.Setenv("PORT", "80")
	os.Setenv("DATABASE_SAVE_INTERVAL_SECONDS", "60")
	os.Setenv("ENABLE_CORS", "true")
	logger = log.New(os.Stdout, "http: ", log.LstdFlags)
}

func main() {
	LoadDatabase()
	go SaveDatabase()
	StartHttpRouter()
}
