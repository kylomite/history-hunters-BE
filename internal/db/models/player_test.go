package models

import (
	"testing"
	"time"
	"historyHunters/internal/db"
)


func TestPlayerFields(t *testing.T) {
	p := &Player{
		ID:             1,
		Email:          "testEmail@example.com",
		PasswordDigest: "hashed_password",
		Avatar:         "/assets/avatar_imgs/avatar_1.png",
		Score:          0,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	// Test email
	if p.Email != "testEmail@example.com" {
		t.Errorf("Expected email to be testEmail@example.com, got %s", p.Email)
	}

	// Test score
	if p.Score != 0 {
		t.Errorf("Expected score to be 0, got %d", p.Score)
	}

	// Test avatar path
	if p.Avatar != "/assets/avatar_imgs/avatar_1.png" {
		t.Errorf("Expected avatar to be /assets/avatar_imgs/avatar_1.png, got %s", p.Avatar)
	}
}
func TestNewPlayerDefaultScore(t *testing.T) {
	player, err := NewPlayer("test@example.com", "hashed_password", "/assets/avatar_imgs/avatar_1.png")
	if err != nil {
		t.Fatalf("failed to create player: %v", err)
	}

	if player.Score != 0 {
		t.Errorf("expected score to default to 0, got %d", player.Score)
	}
}

func TestNewPlayerEmailUniqueness (t *testing.T) {
	testDB := db.InitTestDB()

	// Create the first player
	player1 := &Player{
		Email:          "unique@example.com",
		PasswordDigest: "hashed_password",
		Avatar:         "/assets/avatar_imgs/avatar_1.png",
		Score:          0,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := testDB.Create(player1).Error; err != nil {
		t.Fatalf("Failed to create first player: %v", err)
	}

	player2 := &Player{
		Email:          "unique@example.com", // Same email
		PasswordDigest: "another_hashed_password",
		Avatar:         "/assets/avatar_imgs/avatar_2.png",
		Score:          100,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err := testDB.Create(player2).Error

	if err == nil {
		t.Errorf("Expected an error due to unique email constraint, but got none")
	}

	if err.Error() != "email already exists" {
		t.Errorf("Expected 'email already exists' error, got %v", err)
	}
}