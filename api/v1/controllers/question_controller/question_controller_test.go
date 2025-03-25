package question_controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"historyHunters/internal/models/player"
	"historyHunters/internal/models/stage"
	"historyHunters/internal/models/player_session"
	"historyHunters/internal/models/question"

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

	_, _ = db.Exec("DELETE FROM questions")
	_, _ = db.Exec("DELETE FROM player_sessions")
	_, _ = db.Exec("DELETE FROM players")
	_, _ = db.Exec("DELETE FROM stages")

	_, _ = db.Exec("ALTER SEQUENCE questions_id_seq RESTART WITH 1")
	_, _ = db.Exec("ALTER SEQUENCE player_sessions_id_seq RESTART WITH 1")
	_, _ = db.Exec("ALTER SEQUENCE players_id_seq RESTART WITH 1")
	_, _ = db.Exec("ALTER SEQUENCE stages_id_seq RESTART WITH 1")

	return db
}

func setupRouter(db *sql.DB) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/players", func(r chi.Router) {
		r.Route("/{id}/player_sessions", func(r chi.Router) {
			r.Route("/{id}/questions", func(r chi.Router) {
				r.Post("/", CreateQuestion(db))
				r.Get("/{question_id}", GetQuestionByID(db))
			})
		})
	})

	return router
}

func TestCreateQuestion(t *testing.T) {
	db := setupDB(t)
	defer db.Close()

	player := &player.Player{Email: "test@test.com", PasswordDigest: "hashed_password", Avatar: "avatar.png"}
	err := player.Save(db)
	assert.NoError(t, err)

	stage := &stage.Stage{Title: "Test Stage", BackgroundImg: "bg.png", Difficulty: 2}
	err = stage.Save(db)
	assert.NoError(t, err)

	ps := player_session.NewPlayerSession(player.ID, stage.ID, 3)
	err = ps.Save(db)
	assert.NoError(t, err)

	payload := []byte(`{"question_text": "What is the capital of France?"}`)

	router := setupRouter(db)
	req, _ := http.NewRequest("POST", "/players/1/player_sessions/1/questions", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var createdQuestion question.Question
	err = json.NewDecoder(rr.Body).Decode(&createdQuestion)
	assert.NoError(t, err)
	assert.Equal(t, "What is the capital of France?", createdQuestion.QuestionText)
}

func TestGetQuestionByID(t *testing.T) {
	db := setupDB(t)
	defer db.Close()

	player := &player.Player{Email: "test2@test.com", PasswordDigest: "hashed_password", Avatar: "avatar2.png"}
	err := player.Save(db)
	assert.NoError(t, err)

	stage := &stage.Stage{Title: "Stage 2", BackgroundImg: "bg2.png", Difficulty: 3}
	err = stage.Save(db)
	assert.NoError(t, err)

	ps := player_session.NewPlayerSession(player.ID, stage.ID, 3)
	err = ps.Save(db)
	assert.NoError(t, err)

	testQuestion := question.NewQuestion(ps.ID, "What is 2 + 2?")
	err = testQuestion.Save(db)
	assert.NoError(t, err)

	router := setupRouter(db)
	req, _ := http.NewRequest("GET", "/players/1/player_sessions/1/questions/1", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var fetchedQuestion question.Question
	err = json.NewDecoder(rr.Body).Decode(&fetchedQuestion)
	assert.NoError(t, err)
	assert.Equal(t, "What is 2 + 2?", fetchedQuestion.QuestionText)
}