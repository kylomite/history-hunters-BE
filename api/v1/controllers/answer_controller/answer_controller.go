package answer_controller

import (
	"encoding/json"
	"database/sql"
	"net/http"
	"strconv"

	"historyHunters/internal/models/answer"

	"github.com/go-chi/chi/v5"
)


func CreateAnswer(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newAnswer answer.Answer

		err := json.NewDecoder(r.Body).Decode(&newAnswer)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if newAnswer.QuestionID == 0 || newAnswer.AnswerText == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		err = newAnswer.Save(db)
		if err != nil {
			http.Error(w, "Failed to create answer", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newAnswer)
	}
}

func GetAllAnswers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the question_id from the URL
		questionIDStr := chi.URLParam(r, "id")
		questionID, err := strconv.Atoi(questionIDStr)
		if err != nil {
			http.Error(w, "Invalid question ID", http.StatusBadRequest)
			return
		}

		// Fetch answers filtered by the question_id
		answers, err := answer.GetAnswersByQuestionID(db, questionID)
		if err != nil {
			http.Error(w, "Failed to retrieve answers", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(answers)
	}
}

func GetAnswerByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		answerIDStr := chi.URLParam(r, "answer_id")
		answerID, err := strconv.Atoi(answerIDStr)
		if err != nil {
			http.Error(w, "Invalid answer ID", http.StatusBadRequest)
			return
		}

		a, err := answer.FindByID(db, answerID)
		if err != nil {
			http.Error(w, "Answer not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(a)
	}
}