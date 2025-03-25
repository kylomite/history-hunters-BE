package stage_controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"historyHunters/internal/models/stage"
	"github.com/go-chi/chi/v5"
)

func CreateStage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Title         string `json:"title"`
			BackgroundImg string `json:"background_img"`
			Difficulty    int    `json:"difficulty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		newStage := stage.NewStage(req.Title, req.BackgroundImg, req.Difficulty)

		if err := newStage.Save(db); err != nil {
			http.Error(w, "Failed to create stage: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newStage)
	}
}

func GetAllStages(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stages, err := stage.GetAllStages(db)
		if err != nil {
			http.Error(w, "Failed to fetch stages", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(stages)
	}
}

func GetStageByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid stage ID", http.StatusBadRequest)
			return
		}

		stage, err := stage.FindStageByID(db, id)
		if err != nil {
			http.Error(w, "Stage not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(stage)
	}
}

func UpdateStage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		var req struct {
			Title         string `json:"title"`
			BackgroundImg string `json:"background_img"`
			Difficulty    int    `json:"difficulty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		stage, err := stage.FindStageByID(db, id)
		if err != nil {
			http.Error(w, "Stage not found", http.StatusNotFound)
			return
		}

		stage.Title = req.Title
		stage.BackgroundImg = req.BackgroundImg
		stage.Difficulty = req.Difficulty

		if err := stage.Update(db); err != nil {
			http.Error(w, "Failed to update stage", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(stage)
	}
}

func DeleteStage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		if err := stage.DeleteStage(db, id); err != nil {
			http.Error(w, "Failed to delete stage", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}