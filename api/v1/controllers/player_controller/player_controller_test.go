package player_controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"historyHunters/internal/models/player"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func setupDB(t *testing.T) *sql.DB {
	err := godotenv.Load("../../../../.env.test")
	if err != nil {
		t.Fatalf("Failed to load .env.test file: %v", err)
	}

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		t.Fatalf("DATABASE_URL is not set")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	_, _ = db.Exec("DELETE FROM players")
	_, _ = db.Exec("ALTER SEQUENCE players_id_seq RESTART WITH 1")

	return db
}

func setupRouter() *chi.Mux {
	router := chi.NewRouter()
	return router
}

func TestCreatePlayer(t *testing.T) {
	db := setupDB(t)
	defer db.Close()

	payload := []byte(`{
		"email": "newplayer@test.com",
		"password_digest": "hashed_password",
		"avatar": "avatar.png"
	}`)

	router := setupRouter()
	router.Post("/players", CreatePlayer(db))

	req, _ := http.NewRequest("POST", "/players", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var createdPlayer player.Player
	err := json.NewDecoder(rr.Body).Decode(&createdPlayer)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}

	assert.Equal(t, "newplayer@test.com", createdPlayer.Email)
	assert.Equal(t, "avatar.png", createdPlayer.Avatar)

	var dbPlayer player.Player
	err = db.QueryRow(`SELECT email, avatar, score FROM players WHERE email = $1`, createdPlayer.Email).
		Scan(&dbPlayer.Email, &dbPlayer.Avatar, &dbPlayer.Score)
	if err != nil {
		t.Fatalf("Error querying database: %v", err)
	}

	assert.Equal(t, "newplayer@test.com", dbPlayer.Email)
	assert.Equal(t, "avatar.png", dbPlayer.Avatar)
	assert.Equal(t, 0, dbPlayer.Score) // Default score value
}

func TestGetAllPlayers(t *testing.T) {
	db := setupDB(t)
	defer db.Close()

	router := setupRouter()
	router.Get("/players", GetAllPlayers(db))

	req, _ := http.NewRequest("GET", "/players", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetPlayerByID(t *testing.T) {
	db := setupDB(t)
	defer db.Close()

	player := player.Player{Email: "test@test.com", PasswordDigest: "hashed_password", Avatar: "avatar.png", Score: 10}
	player.Save(db)

	router := setupRouter()
	router.Get("/players/{id}", GetPlayerByID(db))

	req, _ := http.NewRequest("GET", "/players/1", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestUpdatePlayer(t *testing.T) {
	db := setupDB(t)
	defer db.Close()

	player := player.Player{Email: "update@test.com", PasswordDigest: "hashed_password", Avatar: "avatar.png", Score: 20}
	player.Save(db)

	router := setupRouter()
	router.Patch("/players/{id}", UpdatePlayer(db))

	payload := []byte(`{"email": "updated@test.com", "score": 30}`)
	req, _ := http.NewRequest("PATCH", "/players/1", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeletePlayer(t *testing.T) {
	db := setupDB(t)
	defer db.Close()

	player := player.Player{Email: "delete@test.com", PasswordDigest: "hashed_password", Avatar: "avatar.png"}
	player.Save(db)

	router := setupRouter()
	router.Delete("/players/{id}", DeletePlayer(db))

	req, _ := http.NewRequest("DELETE", "/players/1", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}