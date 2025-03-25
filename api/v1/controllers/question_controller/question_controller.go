package question_controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"historyHunters/internal/models/question"

	"github.com/go-chi/chi/v5"
)

func CreateQuestion(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		playerSessionID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid player session ID", http.StatusBadRequest)
			return
		}

		var reqBody struct {
			QuestionText string `json:"question_text"`
		}

		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		q := question.NewQuestion(playerSessionID, reqBody.QuestionText)

		err = q.Save(db)
		if err != nil {
			http.Error(w, "Failed to create question", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(q)
	}
}

func GetQuestionByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		questionID, err := strconv.Atoi(chi.URLParam(r, "question_id"))
		if err != nil {
			http.Error(w, "Invalid question ID", http.StatusBadRequest)
			return
		}

		playerID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid player ID", http.StatusBadRequest)
			return
		}

		q, err := question.FindByID(db, questionID, playerID)
		if err != nil {
			http.Error(w, "Question not found or unauthorized", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(q)
	}
}