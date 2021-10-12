package main

import (
	"encoding/json"
	"os"
	"strconv"
	"time"
)

var data map[string]string

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
	}
	return nil
}

func LoadDatabase() error {
	db, _ := os.ReadFile("db.json")
	return json.Unmarshal(db, &data)
}
