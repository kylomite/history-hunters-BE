package models

import (
	"historyHunters/internal/db"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestPlayerFields(t *testing.T) {
	player := NewPlayer("testEmail@example.com", "hashed_password", "avatar_1.png")

	if player.Email != "testEmail@example.com" {
		t.Errorf("Expected email to be testEmail@example.com, got %s", player.Email)
	}

	if player.Score != 0 {
		t.Errorf("Expected score to be 0, got %d", player.Score)
	}

	if player.Avatar != "avatar_1.png" {
		t.Errorf("Expected avatar to be png, got %s", player.Avatar)
	}
}

func TestNewPlayerDefaultScore(t *testing.T) {
	player := NewPlayer("default@example.com", "hashed_password", "avatar_1.png")

	if player.Score != 0 {
		t.Errorf("Expected score to default to 0, got %d", player.Score)
	}
}

func TestNewPlayerEmailUniqueness(t *testing.T) {
	err := godotenv.Load("../../../.env.test")
	if err != nil {
		log.Println("Failed to load .env file:", err)
	}

	db, err := db.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()
	
	log.Println("Connecting to DB:", os.Getenv("DB_NAME"))

	_, _ = db.Exec("DELETE FROM players WHERE email = 'unique@example.com'")

	player1 := NewPlayer("unique@example.com", "hashed_password", "avatar_1.png")

	err = player1.Save(db)
	if err != nil {
		t.Fatalf("Failed to create first player: %v", err)
	}

	player2 := NewPlayer("unique@example.com", "another_hashed_password", "avatar_2.png")

	err = player2.Save(db)
	if err == nil {
		t.Errorf("Expected an error due to unique email constraint, but got none")
	}

	if err.Error() != "email already exists" {
		t.Errorf("Expected 'email already exists' error, got %v", err)
	}
}