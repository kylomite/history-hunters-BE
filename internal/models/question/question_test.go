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

	err = player.DeletePlayer(db, player.ID)
    if err != nil {
        t.Fatalf("Error deleting player: %v", err)
    }

	err = stage.DeleteStage(db, testStage.ID)
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

func TestFindByID(t *testing.T) {
	err := godotenv.Load("../../../.env.test")
	if err != nil {
		log.Println("Failed to load .env file:", err)
	}

	db, err := db.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Create a player
	player := &player.Player{
		Email:          fmt.Sprintf("test+%d@example.com", time.Now().UnixNano()),
		PasswordDigest: "hashedpassword",
		Avatar:         "avatar.png",
	}

	err = player.Save(db)
	if err != nil {
		t.Fatalf("Error saving player: %v", err)
	}

	// Create a stage
	testStage := &stage.Stage{
		Title:         fmt.Sprintf("Test Stage %d", time.Now().UnixNano()),
		BackgroundImg: "background.png",
		Difficulty:    3,
	}

	err = testStage.Save(db)
	if err != nil {
		t.Fatalf("Error saving stage: %v", err)
	}

	// Create a player session
	playerSession := player_session.NewPlayerSession(player.ID, testStage.ID, 3)
	err = playerSession.Save(db)
	if err != nil {
		t.Fatalf("Error saving player session: %v", err)
	}

	// Create a question
	question := NewQuestion(playerSession.ID, "What is the capital of France?")
	err = question.Save(db)
	if err != nil {
		t.Fatalf("Error saving question: %v", err)
	}

	// Test finding the question by ID
	foundQuestion, err := FindByID(db, question.ID, player.ID)
	if err != nil {
		t.Fatalf("Error finding question: %v", err)
	}

	// Assertions
	if foundQuestion.ID != question.ID {
		t.Errorf("Expected ID %d, got %d", question.ID, foundQuestion.ID)
	}

	if foundQuestion.QuestionText != "What is the capital of France?" {
		t.Errorf("Expected question text 'What is the capital of France?', got %s", foundQuestion.QuestionText)
	}

	if foundQuestion.PlayerSessionID != playerSession.ID {
		t.Errorf("Expected player_session_id %d, got %d", playerSession.ID, foundQuestion.PlayerSessionID)
	}

	// Cleanup
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
		t.Errorf("Error deleting player: %v", err)
	}
}

func TestUnauthorizedQuestionAccess(t *testing.T) {
	err := godotenv.Load("../../../.env.test")
	if err != nil {
		log.Println("Failed to load .env file:", err)
	}

	db, err := db.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Create player 1
	player1 := &player.Player{
		Email:          fmt.Sprintf("player1+%d@example.com", time.Now().UnixNano()),
		PasswordDigest: "hashedpassword",
		Avatar:         "avatar1.png",
	}
	err = player1.Save(db)
	if err != nil {
		t.Fatalf("Error saving player 1: %v", err)
	}

	// Create player 2
	player2 := &player.Player{
		Email:          fmt.Sprintf("player2+%d@example.com", time.Now().UnixNano()),
		PasswordDigest: "hashedpassword",
		Avatar:         "avatar2.png",
	}
	err = player2.Save(db)
	if err != nil {
		t.Fatalf("Error saving player 2: %v", err)
	}

	// Create stage
	testStage := &stage.Stage{
		Title:         "Test Stage",
		BackgroundImg: "background.png",
		Difficulty:    3,
	}
	err = testStage.Save(db)
	if err != nil {
		t.Fatalf("Error saving stage: %v", err)
	}

	// Create a player session for player 1
	session1 := player_session.NewPlayerSession(player1.ID, testStage.ID, 3)
	err = session1.Save(db)
	if err != nil {
		t.Fatalf("Error saving player 1 session: %v", err)
	}

	// Create a player session for player 2
	session2 := player_session.NewPlayerSession(player2.ID, testStage.ID, 3)
	err = session2.Save(db)
	if err != nil {
		t.Fatalf("Error saving player 2 session: %v", err)
	}

	// Create question for player 1's session
	question := NewQuestion(session1.ID, "What is the capital of France?")
	err = question.Save(db)
	if err != nil {
		t.Fatalf("Error saving question: %v", err)
	}

	// Attempt to access player 1's question using player 2's ID
	_, err = FindByID(db, question.ID, player2.ID)
	if err == nil {
		t.Fatalf("Expected unauthorized access error, but got nil")
	}

	if err.Error() != "question not found or unauthorized access" {
		t.Errorf("Expected unauthorized access error, got %v", err)
	}

	// Cleanup
	err = question.Delete(db)
	if err != nil {
		t.Errorf("Error deleting question: %v", err)
	}

	err = session1.Delete(db)
	if err != nil {
		t.Errorf("Error deleting player 1 session: %v", err)
	}

	err = session2.Delete(db)
	if err != nil {
		t.Errorf("Error deleting player 2 session: %v", err)
	}

	err = stage.DeleteStage(db, testStage.ID)
	if err != nil {
		t.Errorf("Error deleting stage: %v", err)
	}

	err = player1.DeletePlayer(db, player1.ID)
	if err != nil {
		t.Errorf("Error deleting player 1: %v", err)
	}

	err = player2.DeletePlayer(db, player2.ID)
	if err != nil {
		t.Errorf("Error deleting player 2: %v", err)
	}
}