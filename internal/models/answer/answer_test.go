package answer

import (
	"log"
	"testing"
	"fmt"
	"time"

	"historyHunters/internal/db"
	"historyHunters/internal/models/player"
	"historyHunters/internal/models/stage"
	"historyHunters/internal/models/player_session"
	"historyHunters/internal/models/question"

	"github.com/joho/godotenv"
)

func TestAnswerFields(t *testing.T) {
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

	testStage := &stage.Stage{
		Title:         fmt.Sprintf("Test Stage %d", time.Now().UnixNano()),
		BackgroundImg: "background.png",
		Difficulty:    3,
	}
	err = testStage.Save(db)
	if err != nil {
		t.Fatalf("Error saving stage: %v", err)
	}

	playerSession := player_session.NewPlayerSession(player.ID, testStage.ID, 3)
	err = playerSession.Save(db)
	if err != nil {
		t.Fatalf("Error saving player session: %v", err)
	}

	question := question.NewQuestion(playerSession.ID, "What is the capital of France?")
	err = question.Save(db)
	if err != nil {
		t.Fatalf("Error saving question: %v", err)
	}

	answer := NewAnswer(question.ID, "Paris", true)
	err = answer.Save(db)
	if err != nil {
		t.Fatalf("Error saving answer: %v", err)
	}

	if answer.ID == 0 {
		t.Errorf("Expected answer ID to be set, got %d", answer.ID)
	}

	if answer.QuestionID != question.ID {
		t.Errorf("Expected question_id to be %d, got %d", question.ID, answer.QuestionID)
	}

	if answer.AnswerText != "Paris" {
		t.Errorf("Expected text to be 'Paris', got %s", answer.AnswerText)
	}

	if !answer.Correct {
		t.Errorf("Expected correct to be true, got %v", answer.Correct)
	}

	err = answer.Delete(db)
	if err != nil {
		t.Errorf("Error deleting answer: %v", err)
	}

	err = question.Delete(db)
	if err != nil {
		t.Errorf("Error deleting question: %v", err)
	}

	err = playerSession.Delete(db)
	if err != nil {
		t.Errorf("Error deleting player session: %v", err)
	}

	err = stage.DeleteStage(db, testStage.ID)
	if err != nil {
		t.Errorf("Error deleting stage: %v", err)
	}

	err = player.DeletePlayer(db, player.ID)
    if err != nil {
        t.Fatalf("Error deleting player: %v", err)
    }
}