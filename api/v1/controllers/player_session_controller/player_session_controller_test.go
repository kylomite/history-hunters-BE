package player_session_controller

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"historyHunters/internal/models/player_session"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"
// 	"time"

// 	_ "github.com/lib/pq"
// 	"github.com/go-chi/chi/v5"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/joho/godotenv"
// 	"os"
// )

// func setupDB(t *testing.T) *sql.DB {
// 	// Load the environment variables from .env.test
// 	err := godotenv.Load("../../../../.env.test")
// 	if err != nil {
// 		t.Fatalf("Failed to load .env.test file: %v", err)
// 	}

// 	// Get the database URL from environment variables
// 	connStr := os.Getenv("DATABASE_URL")
// 	if connStr == "" {
// 		t.Fatalf("DATABASE_URL is not set")
// 	}

// 	// Open the database connection
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		t.Fatalf("Failed to connect to the database: %v", err)
// 	}

// 	// Clean up the database by deleting existing records
// 	_, _ = db.Exec("DELETE FROM player_sessions")
// 	_, _ = db.Exec("ALTER SEQUENCE player_sessions_id_seq RESTART WITH 1")

// 	return db
// }

// func setupRouter() *chi.Mux {
// 	// Set up the router for handling requests
// 	router := chi.NewRouter()
// 	return router
// }

// func TestCreatePlayerSession(t *testing.T) {
// 	// Setup the database and defer closing the connection
// 	db := setupDB(t)
// 	defer db.Close()

// 	// Set up the router and define the POST route
// 	r := setupRouter()
// 	r.Post("/players/1/player_sessions", CreatePlayerSession(db))

// 	// Define the test payload
// 	payload := `{
// 		"player_id": 1,
// 		"stage_id": 1,
// 		"lives": 3
// 	}`

// 	// Create the request for creating a player session
// 	req := httptest.NewRequest(http.MethodPost, "/players/1/player_sessions", strings.NewReader(payload))
// 	resp := httptest.NewRecorder()

// 	// Perform the request
// 	r.ServeHTTP(resp, req)

// 	// Check response code
// 	assert.Equal(t, http.StatusCreated, resp.Code)

// 	// Check if player session is returned correctly
// 	var ps player_session.PlayerSession
// 	err := json.NewDecoder(resp.Body).Decode(&ps)
// 	assert.NoError(t, err)
// 	assert.Equal(t, 1, ps.PlayerID)
// 	assert.Equal(t, 2, ps.StageID)
// 	assert.Equal(t, 3, ps.Lives)
// 	assert.NotZero(t, ps.ID) // ID should be set by the database

// 	// Check if the player session was inserted into the database
// 	var dbSession player_session.PlayerSession
// 	err = db.QueryRow(`SELECT id, player_id, stage_id, lives FROM player_sessions WHERE id = $1`, ps.ID).
// 		Scan(&dbSession.ID, &dbSession.PlayerID, &dbSession.StageID, &dbSession.Lives)
// 	if err != nil {
// 		t.Fatalf("Error querying database: %v", err)
// 	}

// 	assert.Equal(t, ps.ID, dbSession.ID)
// 	assert.Equal(t, ps.PlayerID, dbSession.PlayerID)
// 	assert.Equal(t, ps.StageID, dbSession.StageID)
// 	assert.Equal(t, ps.Lives, dbSession.Lives)
// }

// func TestGetPlayerSessionByID(t *testing.T) {
// 	// Setup the database and defer closing the connection
// 	db := setupDB(t)
// 	defer db.Close()

// 	// Set up the router and define the GET route
// 	r := setupRouter()
// 	r.Get("/player_sessions/{session_id}", GetPlayerSessionByID(db))

// 	// Mock an existing session to retrieve
// 	mockSession := &player_session.PlayerSession{
// 		ID:        1,
// 		PlayerID:  1,
// 		StageID:   2,
// 		Lives:     3,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}
// 	_, err := db.Exec(`INSERT INTO players (id) VALUES ($1)`, 1)
// 	if err != nil {
// 		t.Fatalf("Error inserting mock player: %v", err)
// 	}
	
// 	// Insert the mock session into the database
// 	_, err = db.Exec(`INSERT INTO player_sessions (id, player_id, stage_id, lives, created_at, updated_at) 
// 		VALUES ($1, $2, $3, $4, $5, $6)`,
// 		mockSession.ID, mockSession.PlayerID, mockSession.StageID, mockSession.Lives, mockSession.CreatedAt, mockSession.UpdatedAt)
// 	if err != nil {
// 		t.Fatalf("Error inserting mock player session: %v", err)
// 	}

// 	// Create the request for getting a player session
// 	req := httptest.NewRequest(http.MethodGet, "/player_sessions/1", nil)
// 	resp := httptest.NewRecorder()

// 	// Perform the request
// 	r.ServeHTTP(resp, req)

// 	// Assert response
// 	assert.Equal(t, http.StatusOK, resp.Code)

// 	var ps player_session.PlayerSession
// 	err = json.NewDecoder(resp.Body).Decode(&ps)
// 	assert.NoError(t, err)
// 	assert.Equal(t, mockSession.ID, ps.ID)
// 	assert.Equal(t, mockSession.PlayerID, ps.PlayerID)
// 	assert.Equal(t, mockSession.StageID, ps.StageID)
// 	assert.Equal(t, mockSession.Lives, ps.Lives)
// }

// func TestUpdatePlayerSession(t *testing.T) {
// 	// Setup the database and defer closing the connection
// 	db := setupDB(t)
// 	defer db.Close()

// 	// Set up the router and define the PATCH route
// 	r := setupRouter()
// 	r.Patch("/player_sessions/{session_id}", UpdatePlayerSession(db))

// 	// Mock an existing session to update
// 	mockSession := &player_session.PlayerSession{
// 		ID:        1,
// 		PlayerID:  1,
// 		StageID:   2,
// 		Lives:     3,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	_, err := db.Exec(`INSERT INTO players (id) VALUES ($1)`, 1)
// 	if err != nil {
// 		t.Fatalf("Error inserting mock player: %v", err)
// 	}

// 	// Insert the mock session into the database
// 	_, err = db.Exec(`INSERT INTO player_sessions (id, player_id, stage_id, lives, created_at, updated_at) 
// 		VALUES ($1, $2, $3, $4, $5, $6)`,
// 		mockSession.ID, mockSession.PlayerID, mockSession.StageID, mockSession.Lives, mockSession.CreatedAt, mockSession.UpdatedAt)
// 	if err != nil {
// 		t.Fatalf("Error inserting mock player session: %v", err)
// 	}

// 	// Define the updated payload
// 	payload := `{
// 		"lives": 2
// 	}`

// 	// Create the request for updating the player session
// 	req := httptest.NewRequest(http.MethodPatch, "/player_sessions/1", strings.NewReader(payload))
// 	resp := httptest.NewRecorder()

// 	// Perform the request
// 	r.ServeHTTP(resp, req)

// 	// Assert response
// 	assert.Equal(t, http.StatusOK, resp.Code)

// 	var ps player_session.PlayerSession
// 	err = json.NewDecoder(resp.Body).Decode(&ps)
// 	assert.NoError(t, err)
// 	assert.Equal(t, mockSession.ID, ps.ID)
// 	assert.Equal(t, mockSession.PlayerID, ps.PlayerID)
// 	assert.Equal(t, mockSession.StageID, ps.StageID)
// 	assert.Equal(t, 2, ps.Lives) // Updated value
// }