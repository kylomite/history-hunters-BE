package player

import (
	"database/sql"
	"testing"

	"historyHunters/internal/db"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func SetupTestDB(t *testing.T) *sql.DB {
	err := godotenv.Load("../../../.env.test")
	if err != nil {
		t.Fatalf("Failed to load .env file: %v", err)
	}

	db, err := db.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	return db
}

func CleanupTestDB(t *testing.T, db *sql.DB) {
	_, err := db.Exec("DELETE FROM players")
	if err != nil {
		t.Fatalf("Failed to clean up database: %v", err)
	}
}

func TestPlayerFields(t *testing.T) {
	player := NewPlayer("testEmail@example.com", "hashed_password", "avatar_1.png")

	assert.Equal(t, "testEmail@example.com", player.Email)
	assert.Equal(t, 0, player.Score)
	assert.Equal(t, "avatar_1.png", player.Avatar)
}

func TestNewPlayerMissingRequiredFields(t *testing.T) {
	db := SetupTestDB(t)
	defer db.Close()

	tests := []struct {
		player  *Player
		message string
	}{
		{NewPlayer("", "hashed_password", "avatar_1.png"), "Expected error when creating player without email"},
		{NewPlayer("testEmail@example.com", "", "avatar_1.png"), "Expected error when creating player without password"},
		{NewPlayer("testEmail@example.com", "hashed_password", ""), "Expected error when creating player without avatar"},
	}

	for _, tt := range tests {
		err := tt.player.Save(db)
		if err == nil {
			t.Errorf("Expected error, but got none: %s", tt.message)
		}
	}
}

func TestNewPlayerDefaultScore(t *testing.T) {
	player := NewPlayer("default@example.com", "hashed_password", "avatar_1.png")

	assert.Equal(t, 0, player.Score)
}

func TestNewPlayerEmailUniqueness(t *testing.T) {
	db := SetupTestDB(t)
	defer db.Close()

	CleanupTestDB(t, db)

	player1 := NewPlayer("unique@example.com", "hashed_password", "avatar_1.png")
	err := player1.Save(db)
	if err != nil {
		t.Fatalf("Failed to create first player: %v", err)
	}

	player2 := NewPlayer("unique@example.com", "another_hashed_password", "avatar_2.png")
	err = player2.Save(db)

	if err == nil {
		t.Errorf("Expected error due to unique email constraint, but got none")
	}

	if err.Error() != "email already exists" {
		t.Errorf("Expected 'email already exists' error, got %v", err)
	}
}