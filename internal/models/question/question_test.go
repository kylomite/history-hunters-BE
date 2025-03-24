package question

import (
	"log"
	"testing"
	"fmt"
	"time"

	"historyHunters/internal/db"
	"historyHunters/internal/models/player"
	"historyHunters/internal/models/stage"
	"historyHunters/internal/models/player_session"

	"github.com/joho/godotenv"
)

func TestQuestionFields(t *testing.T) {
	err := godotenv.Load("../../../.env.test")
	if err != nil {
		log.Println("Failed to load .env file:", err)
	}

	db, err := db.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()


	player := &player.Player{
		Email:          fmt.Sprintf("test+%d@example.com", time.Now().UnixNano()),
		PasswordDigest: "hashedpassword",
		Avatar:         "avatar.png",
	}

	err = player.Save(db)
	if err != nil {
		t.Fatalf("Error saving player: %v", err)
	}


	stage := &stage.Stage{
		Title:         fmt.Sprintf("Test Stage %d", time.Now().UnixNano()),
		BackgroundImg: "background.png",
		Difficulty:    3,
	}

	err = stage.Save(db)
	if err != nil {
		t.Fatalf("Error saving stage: %v", err)
	}


	playerSession := player_session.NewPlayerSession(player.ID, stage.ID, 3)
	err = playerSession.Save(db)
	if err != nil {
		t.Fatalf("Error saving player session: %v", err)
	}

	question := NewQuestion(playerSession.ID, "What is the capital of France?")


	err = question.Save(db)
	if err != nil {
		t.Fatalf("Error saving question: %v", err)
	}

	if question.ID == 0 {
		t.Errorf("Expected question ID to be set, got %d", question.ID)
	}

	if question.PlayerSessionID != playerSession.ID {
		t.Errorf("Expected player_session_id to be %d, got %d", playerSession.ID, question.PlayerSessionID)
	}

	if question.QuestionText != "What is the capital of France?" {
		t.Errorf("Expected text to be 'What is the capital of France?', got %s", question.QuestionText)
	}

	err = player.Delete(db)
	if err != nil {
		t.Errorf("Error deleting player: %v", err)
	}

	err = stage.Delete(db)
	if err != nil {
		t.Errorf("Error deleting stage: %v", err)
	}

	err = playerSession.Delete(db)
	if err != nil {
		t.Errorf("Error deleting player session: %v", err)
	}

	err = question.Delete(db)
	if err != nil {
		t.Errorf("Error deleting question: %v", err)
	}
}