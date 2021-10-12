package main

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	os.Setenv("PORT", "8080")
	os.Setenv("DATABASE_SAVE_INTERVAL_SECONDS", "60")
	logger = log.New(os.Stdout, "http: ", log.LstdFlags)
}

func main() {
	LoadDatabase()
	go SaveDatabase()
	StartHttpRouter()
}
