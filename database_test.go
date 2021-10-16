package main

import (
	"os"
	"testing"
)

func TestLoadDatabase(t *testing.T) {
	t.Log("Testing LoadDatabase")

	t.Run("returns success result", func(t *testing.T) {
		err := LoadDatabase()
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})
}

func TestRemoveDatabase(t *testing.T) {
	t.Log("Testing RemoveDatabase")

	t.Run("returns success result", func(t *testing.T) {
		err := RemoveDatabase()
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})
}

func TestSaveDatabase(t *testing.T) {
	t.Log("Testing SaveDatabase")

	os.Setenv("DATABASE_SAVE_INTERVAL_SECONDS", "1")

	t.Run("returns success result", func(t *testing.T) {
		if o := <-operation; o != "DATABASE_SAVED" {
			operation <- "STOP_DATABASE_SAVING"
		}
		err := SaveDatabase()

		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})
}
