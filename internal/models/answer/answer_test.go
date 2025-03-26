package answer

import (
	"fmt"
	"log"
	"testing"
	"time"

	"historyHunters/internal/db"
	"historyHunters/internal/models/player"
	"historyHunters/internal/models/player_session"
	"historyHunters/internal/models/question"
	"historyHunters/internal/models/stage"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
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

func TestFindAnswerByID(t *testing.T) {
	err := godotenv.Load("../../../.env.test")
	if err != nil {
		t.Fatalf("Failed to load .env file: %v", err)
	}

	db, err := db.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Create test player
	player := &player.Player{
		Email:          fmt.Sprintf("test+%d@example.com", time.Now().UnixNano()),
		PasswordDigest: "hashedpassword",
		Avatar:         "avatar.png",
	}
	err = player.Save(db)
	assert.NoError(t, err)

	// Create test stage
	testStage := &stage.Stage{
		Title:         fmt.Sprintf("Test Stage %d", time.Now().UnixNano()),
		BackgroundImg: "bg.png",
		Difficulty:    2,
	}
	err = testStage.Save(db)
	assert.NoError(t, err)

	// Create test player session
	ps := player_session.NewPlayerSession(player.ID, testStage.ID, 3)
	err = ps.Save(db)
	assert.NoError(t, err)

	// Create test question
	q := question.NewQuestion(ps.ID, "What is 3 + 3?")
	err = q.Save(db)
	assert.NoError(t, err)

	// Create test answer
	answer := NewAnswer(q.ID, "6", true)
	err = answer.Save(db)
	assert.NoError(t, err)

	// Find the answer by ID
	foundAnswer, err := FindByID(db, answer.ID)
	assert.NoError(t, err)
	assert.NotNil(t, foundAnswer)

	// Verify the properties
	assert.Equal(t, answer.ID, foundAnswer.ID)
	assert.Equal(t, answer.QuestionID, foundAnswer.QuestionID)
	assert.Equal(t, "6", foundAnswer.AnswerText)
	assert.True(t, foundAnswer.Correct)

	// Clean up
	err = answer.Delete(db)
	assert.NoError(t, err)

	err = q.Delete(db)
	assert.NoError(t, err)

	err = ps.Delete(db)
	assert.NoError(t, err)

	err = stage.DeleteStage(db, testStage.ID)
	assert.NoError(t, err)

	err = player.DeletePlayer(db, player.ID)
	assert.NoError(t, err)
}