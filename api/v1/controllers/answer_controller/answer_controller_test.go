package answer_controller

// import (
// 	"bytes"
// 	"database/sql"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"strconv"
// 	"testing"

// 	"historyHunters/internal/models/answer"
// 	"historyHunters/internal/models/player"
// 	"historyHunters/internal/models/player_session"
// 	"historyHunters/internal/models/question"
// 	"historyHunters/internal/models/stage"

// 	"github.com/go-chi/chi/v5"
// 	"github.com/joho/godotenv"
// 	"github.com/stretchr/testify/assert"
// )

// func setupDB(t *testing.T) *sql.DB {
// 	err := godotenv.Load("../../../../.env.test")
// 	if err != nil {
// 		t.Fatalf("Failed to load .env.test file: %v", err)
// 	}

// 	connStr := os.Getenv("DATABASE_URL")
// 	if connStr == "" {
// 		t.Fatalf("DATABASE_URL is not set")
// 	}

// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		t.Fatalf("Failed to connect to the database: %v", err)
// 	}

// 	_, _ = db.Exec("DELETE FROM players")
// 	_, _ = db.Exec("ALTER SEQUENCE players_id_seq RESTART WITH 1")

// 	return db
// }

// func setupRouter(db *sql.DB) *chi.Mux {
// 	router := chi.NewRouter()

// 	router.Route("/players", func(r chi.Router) {
// 		r.Route("/{player_id}/player_sessions", func(r chi.Router) {
// 			r.Route("/{session_id}/questions", func(r chi.Router) {
// 				r.Route("/{question_id}/answers", func(r chi.Router) {
// 					r.Post("/", CreateAnswer(db))                     // Create an answer
// 					r.Get("/", GetAllAnswers(db))                     // Get all answers for a question
// 					r.Get("/{answer_id}", GetAnswerByID(db))          // Get a specific answer
// 				})
// 			})
// 		})
// 	})

// 	return router
// }

// func TestCreateAnswer(t *testing.T) {
// 	db := setupDB(t)
// 	defer db.Close()

// 	// Seed the data
// 	player := &player.Player{Email: "test@test.com", PasswordDigest: "hashed_password", Avatar: "avatar.png"}
// 	assert.NoError(t, player.Save(db))

// 	stage := &stage.Stage{Title: "Test Stage", BackgroundImg: "bg.png", Difficulty: 3}
// 	assert.NoError(t, stage.Save(db))

// 	ps := player_session.NewPlayerSession(player.ID, stage.ID, 3)
// 	assert.NoError(t, ps.Save(db))

// 	q := question.NewQuestion(ps.ID, "What is the capital of Italy?")
// 	assert.NoError(t, q.Save(db))

// 	// Use the full route for creating answers
// 	payload := []byte(`{
// 		"QuestionID": ` + strconv.Itoa(q.ID) + `,
// 		"AnswerText": "Rome",
// 		"Correct": true
// 	}`)

// 	router := setupRouter(db)

// 	req, _ := http.NewRequest("POST",
// 		"/players/"+strconv.Itoa(player.ID)+
// 			"/player_sessions/"+strconv.Itoa(ps.ID)+
// 			"/questions/"+strconv.Itoa(q.ID)+
// 			"/answers", bytes.NewBuffer(payload))

// 	req.Header.Set("Content-Type", "application/json")
// 	rr := httptest.NewRecorder()

// 	router.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusCreated, rr.Code)

// 	var createdAnswer answer.Answer
// 	err := json.NewDecoder(rr.Body).Decode(&createdAnswer)
// 	assert.NoError(t, err)
// 	assert.Equal(t, q.ID, createdAnswer.QuestionID)
// 	assert.Equal(t, "Rome", createdAnswer.AnswerText)
// 	assert.True(t, createdAnswer.Correct)
// }

// func TestGetAllAnswers(t *testing.T) {
// 	db := setupDB(t)
// 	defer db.Close()

// 	// Seed the data
// 	player := &player.Player{Email: "test3@test.com", PasswordDigest: "hashed_password", Avatar: "avatar3.png"}
// 	assert.NoError(t, player.Save(db))

// 	stage := &stage.Stage{Title: "Stage 3", BackgroundImg: "bg3.png", Difficulty: 4}
// 	assert.NoError(t, stage.Save(db))

// 	ps := player_session.NewPlayerSession(player.ID, stage.ID, 4)
// 	assert.NoError(t, ps.Save(db))

// 	q := question.NewQuestion(ps.ID, "What is the capital of France?")
// 	assert.NoError(t, q.Save(db))

// 	// Create answers associated with the question
// 	a1 := answer.NewAnswer(q.ID, "Paris", true)
// 	assert.NoError(t, a1.Save(db))

// 	a2 := answer.NewAnswer(q.ID, "Lyon", false)
// 	assert.NoError(t, a2.Save(db))

// 	// Create a router with the nested structure
// 	router := setupRouter(db)

// 	// Use the full route path with properly associated IDs
// 	req, _ := http.NewRequest("GET",
// 		"/players/"+strconv.Itoa(player.ID)+
// 			"/player_sessions/"+strconv.Itoa(ps.ID)+
// 			"/questions/"+strconv.Itoa(q.ID)+
// 			"/answers", nil)

// 	rr := httptest.NewRecorder()
// 	router.ServeHTTP(rr, req)

// 	// Assert the response code
// 	assert.Equal(t, http.StatusOK, rr.Code)

// 	// Decode the response body
// 	var fetchedAnswers []answer.Answer
// 	err := json.NewDecoder(rr.Body).Decode(&fetchedAnswers)
// 	assert.NoError(t, err)

// 	// Ensure the expected number of answers is returned
// 	assert.Equal(t, 2, len(fetchedAnswers))
// 	assert.Equal(t, "Paris", fetchedAnswers[0].AnswerText)
// 	assert.Equal(t, "Lyon", fetchedAnswers[1].AnswerText)
// }

// func TestGetAnswerByID(t *testing.T) {
// 	db := setupDB(t)
// 	defer db.Close()

// 	// Seed the database
// 	player := &player.Player{Email: "test2@test.com", PasswordDigest: "hashed_password", Avatar: "avatar2.png"}
// 	assert.NoError(t, player.Save(db))

// 	stage := &stage.Stage{Title: "Stage 2", BackgroundImg: "bg2.png", Difficulty: 2}
// 	assert.NoError(t, stage.Save(db))

// 	ps := player_session.NewPlayerSession(player.ID, stage.ID, 3)
// 	assert.NoError(t, ps.Save(db))

// 	q := question.NewQuestion(ps.ID, "What is 5 + 5?")
// 	assert.NoError(t, q.Save(db))

// 	a := answer.NewAnswer(q.ID, "10", true)
// 	assert.NoError(t, a.Save(db))

// 	// Use the correct route path with nested IDs
// 	router := setupRouter(db)
// 	req, _ := http.NewRequest("GET",
// 		"/players/"+strconv.Itoa(player.ID)+
// 			"/player_sessions/"+strconv.Itoa(ps.ID)+
// 			"/questions/"+strconv.Itoa(q.ID)+
// 			"/answers/"+strconv.Itoa(a.ID), nil)

// 	rr := httptest.NewRecorder()
// 	router.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusOK, rr.Code)

// 	var fetchedAnswer answer.Answer
// 	err := json.NewDecoder(rr.Body).Decode(&fetchedAnswer)
// 	assert.NoError(t, err)
// 	assert.Equal(t, a.ID, fetchedAnswer.ID)
// 	assert.Equal(t, "10", fetchedAnswer.AnswerText)
// 	assert.True(t, fetchedAnswer.Correct)
// }