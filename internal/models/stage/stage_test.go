package models

import (
	// "database/sql"
	// "errors"
	"testing"
	// "time"
	"log"


	"historyHunters/internal/db"
	"github.com/joho/godotenv"
)

func TestStageFields(t *testing.T) {
	stage := NewStage("Test Title", "background.png", 3)

	if stage.Title == "" {
		t.Errorf("Expected title to be set, got %s", stage.Title)
	}

	if stage.BackgroundImg == "" {
		t.Errorf("Expected background_img to be set, got %s", stage.BackgroundImg)
	}

	if stage.Difficulty != 3 {
		t.Errorf("Expected difficulty to be 5, got %d", stage.Difficulty)
	}
}

func TestStageInvalidFields(t *testing.T) {
	err := godotenv.Load("../../../.env.test")
	if err != nil {
		log.Println("Failed to load .env file:", err)
	}

	db, err := db.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	stage1 := NewStage("", "background.png", 5)
	err = stage1.Save(db)
	if err == nil || err.Error() != "title is required" {
		t.Errorf("Expected error 'title is required', got %v", err)
	}

	stage2 := NewStage("Test Title", "", 5)
	err = stage2.Save(db)
	if err == nil || err.Error() != "background image is required" {
		t.Errorf("Expected error 'background image is required', got %v", err)
	}

	stage3 := NewStage("Test Title", "background.png", 0)
	err = stage3.Save(db)
	if err == nil || err.Error() != "difficulty is required" {
		t.Errorf("Expected error 'difficulty is required', got %v", err)
	}
}