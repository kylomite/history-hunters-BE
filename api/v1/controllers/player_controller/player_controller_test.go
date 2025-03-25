package player_controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"historyHunters/internal/model"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func setupDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", "your_db_connection_string")
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	return db
}

func TestGetAllPlayers(t *testing.T) {
	db := setupDB(t)
	defer db.Close()

	req, _ := http.NewRequest("GET", "/players", nil)
	rr := httptest.NewRecorder()

	handler := GetAllPlayers(db)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetPlayerByID(t *testing.T) {
	db := setupDB(t)
	defer db.Close()

	player := model.Player{Email: "test@test.com", PasswordDigest: "hashed_password", Avatar: "avatar.png", Score: 10}
	player.Save(db)

	req, _ := http.NewRequest("GET", "/players/1", nil)
	rr := httptest.NewRecorder()

	handler := GetPlayerByID(db)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestUpdatePlayer(t *testing.T) {
	db := setupDB(t)
	defer db.Close()

	player := model.Player{Email: "update@test.com", PasswordDigest: "hashed_password", Avatar: "avatar.png", Score: 20}
	player.Save(db)

	payload := []byte(`{"email": "updated@test.com", "score": 30}`)
	req, _ := http.NewRequest("PATCH", "/players/1", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()

	handler := UpdatePlayer(db)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeletePlayer(t *testing.T) {
	db := setupDB(t)
	defer db.Close()

	player := model.Player{Email: "delete@test.com", PasswordDigest: "hashed_password", Avatar: "avatar.png"}
	player.Save(db)

	req, _ := http.NewRequest("DELETE", "/players/1", nil)
	rr := httptest.NewRecorder()

	handler := DeletePlayer(db)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}