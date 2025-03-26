package answer_controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"historyHunters/internal/models/answer"
	"historyHunters/internal/models/player"
	"historyHunters/internal/models/player_session"
	"historyHunters/internal/models/question"
	"historyHunters/internal/models/stage"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
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

func setupRouter(db *sql.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Post("/answers", CreateAnswer(db))
	router.Get("/answers/{question_id}", GetAnswerByID(db))
	return router
}

func TestCreateAnswer(t *testing.T) {
	db := setupDB(t)
	defer db.Close()

	player := &player.Player{Email: "test@test.com", PasswordDigest: "hashed_password", Avatar: "avatar.png"}
	assert.NoError(t, player.Save(db))

	stage := &stage.Stage{Title: "Test Stage", BackgroundImg: "bg.png", Difficulty: 3}
	assert.NoError(t, stage.Save(db))

	ps := player_session.NewPlayerSession(player.ID, stage.ID, 3)
	assert.NoError(t, ps.Save(db))

	q := question.NewQuestion(ps.ID, "What is the capital of Italy?")
	assert.NoError(t, q.Save(db))

	payload := []byte(`{
		"QuestionID": ` + strconv.Itoa(q.ID) + `,
		"AnswerText": "Rome",
		"Correct": true
	}`)

	router := setupRouter(db)

	req, _ := http.NewRequest("POST", "/answers", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var createdAnswer answer.Answer
	err := json.NewDecoder(rr.Body).Decode(&createdAnswer)
	assert.NoError(t, err)
	assert.Equal(t, q.ID, createdAnswer.QuestionID)
	assert.Equal(t, "Rome", createdAnswer.AnswerText)
	assert.True(t, createdAnswer.Correct)
}

func TestGetAnswerByID(t *testing.T) {
	db := setupDB(t)
	defer db.Close()

	player := &player.Player{Email: "test2@test.com", PasswordDigest: "hashed_password", Avatar: "avatar2.png"}
	assert.NoError(t, player.Save(db))

	stage := &stage.Stage{Title: "Stage 2", BackgroundImg: "bg2.png", Difficulty: 2}
	assert.NoError(t, stage.Save(db))

	ps := player_session.NewPlayerSession(player.ID, stage.ID, 3)
	assert.NoError(t, ps.Save(db))

	q := question.NewQuestion(ps.ID, "What is 5 + 5?")
	assert.NoError(t, q.Save(db))

	a := answer.NewAnswer(q.ID, "10", true)
	assert.NoError(t, a.Save(db))

	router := setupRouter(db)
	req, _ := http.NewRequest("GET", "/answers/"+strconv.Itoa(a.ID), nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var fetchedAnswer answer.Answer
	err := json.NewDecoder(rr.Body).Decode(&fetchedAnswer)
	assert.NoError(t, err)
	assert.Equal(t, a.ID, fetchedAnswer.ID)
	assert.Equal(t, "10", fetchedAnswer.AnswerText)
	assert.True(t, fetchedAnswer.Correct)
}