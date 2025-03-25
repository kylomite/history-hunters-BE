package player_controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"historyHunters/internal/models/player"
	"github.com/go-chi/chi/v5"
)

func CreatePlayer(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newPlayerRequest struct {
			Email          string `json:"email"`
			PasswordDigest string `json:"password_digest"`
			Avatar         string `json:"avatar"`
		}

		err := json.NewDecoder(r.Body).Decode(&newPlayerRequest)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		player := player.NewPlayer(newPlayerRequest.Email, newPlayerRequest.PasswordDigest, newPlayerRequest.Avatar)

		err = player.Save(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(player)
	}
}

func GetAllPlayers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		players, err := player.GetAllPlayers(db)
		if err != nil {
			http.Error(w, "Failed to fetch players", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(players)
	}
}

func GetPlayerByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid player ID", http.StatusBadRequest)
			return
		}

		player, err := player.FindPlayerByID(db, id)
		if err != nil {
			http.Error(w, "Player not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(player)
	}
}

func UpdatePlayer(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		var p player.Player
		err = json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		p.ID = id
		err = p.Update(db)
		if err != nil {
			http.Error(w, "Failed to update player", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(p)
	}
}

func DeletePlayer(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		player, err := player.FindPlayerByID(db, id)
		if err != nil {
			http.Error(w, "Player not found", http.StatusNotFound)
			return
		}

		err = player.DeletePlayer(db, id)
		if err != nil {
			http.Error(w, "Failed to delete player", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}