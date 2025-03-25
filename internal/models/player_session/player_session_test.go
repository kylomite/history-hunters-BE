package player_session

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	"historyHunters/internal/db"
	"historyHunters/internal/models/player"
	"historyHunters/internal/models/stage"

	"github.com/joho/godotenv"
)

func setupTestDB(t *testing.T) *sql.DB {
	err := godotenv.Load("../../../.env.test")
	if err != nil {
		log.Println("Failed to load .env file:", err)
	}

	db, err := db.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	return db
}

func createTestPlayerAndStage(t *testing.T, db *sql.DB) (*player.Player, *stage.Stage) {
	testPlayer := player.NewPlayer(
		fmt.Sprintf("test+%d@example.com", time.Now().UnixNano()),
		"hashed_password",
		"avatar.png",
	)
	err := testPlayer.Save(db)
	if err != nil {
		t.Fatalf("Failed to create test player: %v", err)
	}

	// Create test stage
	testStage := &stage.Stage{
		Title:         fmt.Sprintf("Test Stage %d", time.Now().UnixNano()),
		BackgroundImg: "background.png",
		Difficulty:    3,
	}
	err = testStage.Save(db)
	if err != nil {
		t.Fatalf("Failed to create test stage: %v", err)
	}

	return testPlayer, testStage
}

func cleanupTestData(t *testing.T, db *sql.DB, player *player.Player, argStage *stage.Stage, session *PlayerSession) {
	if session != nil {
		sessionErr := session.Delete(db)
		if sessionErr != nil {
			t.Errorf("Failed to delete player session: %v", sessionErr)
		}
	}

	playerErr := player.DeletePlayer(db, player.ID)
	if playerErr != nil {
		t.Errorf("Failed to delete player: %v", playerErr)
	}

	stageErr := stage.DeleteStage(db, argStage.ID) 
	if stageErr != nil {
		t.Errorf("Failed to delete stage: %v", stageErr)
	}
}
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

	player := player.NewPlayer("deletePlayer@example.com", "hashed_password", "avatar.png")

    err = player.Save(db)
    if err != nil {
        t.Fatalf("Error saving player: %v", err)
    }

    stageTitle := fmt.Sprintf("Test Stage %d", time.Now().UnixNano())
    testStage := &stage.Stage{
        Title:        stageTitle,
        BackgroundImg: "background.png",
        Difficulty:   3,
    }

    err = testStage.Save(db)
    if err != nil {
        t.Fatalf("Error saving stage: %v", err)
    }

    playerSession := NewPlayerSession(player.ID, testStage.ID, 3)

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
func TestCreatePlayerSession(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	player, stage := createTestPlayerAndStage(t, db)

	session := NewPlayerSession(player.ID, stage.ID, 3)
	err := session.Save(db)
	if err != nil {
		t.Fatalf("Failed to create player session: %v", err)
	}

	if session.ID == 0 {
		t.Errorf("Expected session ID to be set, got 0")
	}
	if session.PlayerID != player.ID {
		t.Errorf("Expected player_id %d, got %d", player.ID, session.PlayerID)
	}
	if session.StageID != stage.ID {
		t.Errorf("Expected stage_id %d, got %d", stage.ID, session.StageID)
	}
	if session.Lives != 3 {
		t.Errorf("Expected lives to be 3, got %d", session.Lives)
	}

	cleanupTestData(t, db, player, stage, session)
}

func TestGetPlayerSessionByID(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	player, stage := createTestPlayerAndStage(t, db)

	session := NewPlayerSession(player.ID, stage.ID, 3)
	err := session.Save(db)
	if err != nil {
		t.Fatalf("Failed to create player session: %v", err)
	}

	fetchedSession, err := GetByID(db, session.ID)
	if err != nil {
		t.Fatalf("Failed to fetch player session: %v", err)
	}

	if fetchedSession.ID != session.ID {
		t.Errorf("Expected session ID %d, got %d", session.ID, fetchedSession.ID)
	}
	if fetchedSession.PlayerID != player.ID {
		t.Errorf("Expected player ID %d, got %d", player.ID, fetchedSession.PlayerID)
	}
	if fetchedSession.StageID != stage.ID {
		t.Errorf("Expected stage ID %d, got %d", stage.ID, fetchedSession.StageID)
	}
	if fetchedSession.Lives != 3 {
		t.Errorf("Expected lives to be 3, got %d", fetchedSession.Lives)
	}

	cleanupTestData(t, db, player, stage, session)
}

func TestUpdatePlayerSession(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	player, stage := createTestPlayerAndStage(t, db)

	// Create player session
	session := NewPlayerSession(player.ID, stage.ID, 3)
	err := session.Save(db)
	if err != nil {
		t.Fatalf("Failed to create player session: %v", err)
	}

	session.Lives = 5
	err = session.Update(db)
	if err != nil {
		t.Fatalf("Failed to update player session: %v", err)
	}

	updatedSession, err := GetByID(db, session.ID)
	if err != nil {
		t.Fatalf("Failed to fetch updated player session: %v", err)
	}

	if updatedSession.Lives != 5 {
		t.Errorf("Expected lives to be 5, got %d", updatedSession.Lives)
	}

	cleanupTestData(t, db, player, stage, session)
}