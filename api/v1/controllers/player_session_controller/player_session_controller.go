package player_session_controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"historyHunters/internal/models/player_session"
	"github.com/go-chi/chi/v5"
)

func CreatePlayerSession(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newSessionRequest struct {
			PlayerID int `json:"player_id"`
			StageID  int `json:"stage_id"`
			Lives    int `json:"lives"`
		}

		err := json.NewDecoder(r.Body).Decode(&newSessionRequest)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to parse request body: %s", err), http.StatusBadRequest)
			return
		}

		ps := player_session.PlayerSession{
			PlayerID: newSessionRequest.PlayerID,
			StageID:  newSessionRequest.StageID,
			Lives:    newSessionRequest.Lives,
		}

		err = ps.Save(db)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to save player session: %s", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(ps)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to encode response: %s", err), http.StatusInternalServerError)
		}
	}
}

func GetPlayerSessionByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionIDStr := chi.URLParam(r, "session_id")
		sessionID, err := strconv.Atoi(sessionIDStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid session_id: %s", err), http.StatusBadRequest)
			return
		}

		ps, err := player_session.GetByID(db, sessionID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to retrieve player session: %s", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(ps)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to encode response: %s", err), http.StatusInternalServerError)
		}
	}
}

func UpdatePlayerSession(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionIDStr := chi.URLParam(r, "session_id")
		sessionID, err := strconv.Atoi(sessionIDStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid session_id: %s", err), http.StatusBadRequest)
			return
		}

		var ps player_session.PlayerSession
		err = json.NewDecoder(r.Body).Decode(&ps)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to parse request body: %s", err), http.StatusBadRequest)
			return
		}

		ps.ID = sessionID
		err = ps.Update(db)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to update player session: %s", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(ps)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to encode response: %s", err), http.StatusInternalServerError)
		}
	}
}