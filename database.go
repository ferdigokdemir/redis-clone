package main

import (
	"encoding/json"
	"os"
	"strconv"
	"time"
)

var data map[string]string

var operation chan string = make(chan string, 50)

func SaveDatabase() error {
	interval := os.Getenv("DATABASE_SAVE_INTERVAL_SECONDS")

	if interval == "" {
		interval = "60"
	}

	intervalSeconds, err := strconv.Atoi(interval)

	if err != nil {
		return err
	}

	ticker := time.NewTicker(time.Duration(intervalSeconds) * time.Second)
	for range ticker.C {
		db, _ := json.Marshal(data)
		os.WriteFile("db.json", db, 0644)
		operation <- "DATABASE_SAVED"
		if o := <-operation; o == "STOP_DATABASE_SAVING" {
			println("Stopping database saving")
			ticker.Stop()
			return nil
		}
	}

	return nil
}

func LoadDatabase() error {
	db, err := os.ReadFile("db.json")
	if err != nil {
		db = []byte("{}")
	}
	json.Unmarshal(db, &data)
	operation <- "DATABASE_LOADED"
	return nil
}

func RemoveDatabase() error {
	if _, err := os.Stat("db.json"); err == nil {
		return os.Remove("db.json")
	}
	operation <- "DATABASE_REMOVED"
	return nil
}
