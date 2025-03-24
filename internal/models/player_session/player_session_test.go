package models

import (
	"database/sql"
	"log"
	"testing"

	"historyHunters/internal/db"
	"github.com/joho/godotenv"
)

func TestPlayerSessionFields(t *testing.T) {
	err := godotenv.Load("../../../.env.test")
	if err != nil {
		log.Println("Failed to load .env file:", err)
	}

	db, err := db.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	playerSession := NewPlayerSession(1, 1, 3)

	if playerSession.PlayerID == 0 {
		t.Errorf("Expected player_id to be set, got %d", playerSession.PlayerID)
	}

	if playerSession.StageID == 0 {
		t.Errorf("Expected stage_id to be set, got %d", playerSession.StageID)
	}

	if playerSession.Lives <= 0 {
		t.Errorf("Expected lives to be greater than 0, got %d", playerSession.Lives)
	}

	err = playerSession.Save(db)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestPlayerSessionInvalidFields(t *testing.T) {
	err := godotenv.Load("../../../.env.test")
	if err != nil {
		log.Println("Failed to load .env file:", err)
	}

	db, err := db.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	session1 := NewPlayerSession(0, 1, 3)
	err = session1.Save(db)
	if err == nil || err.Error() != "player_id is required" {
		t.Errorf("Expected error 'player_id is required', got %v", err)
	}

	session2 := NewPlayerSession(1, 0, 3)
	err = session2.Save(db)
	if err == nil || err.Error() != "stage_id is required" {
		t.Errorf("Expected error 'stage_id is required', got %v", err)
	}

	session3 := NewPlayerSession(1, 1, 0)
	err = session3.Save(db)
	if err == nil || err.Error() != "lives must be greater than 0" {
		t.Errorf("Expected error 'lives must be greater than 0', got %v", err)
	}
}