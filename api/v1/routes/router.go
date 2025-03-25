package routes

import (
	"database/sql"
	"net/http"

	"historyHunters/api/v1/controllers/stage_controller"
	"historyHunters/api/v1/controllers/player_controller"
	"historyHunters/api/v1/controllers/player_session_controller"
	"github.com/go-chi/chi/v5"
)

func NewRouter(db *sql.DB) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to History Hunters!"))
	})

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is healthy"))
	})

	router.Route("/players", func(r chi.Router) {
		r.Post("/", player_controller.CreatePlayer(db))
		r.Get("/", player_controller.GetAllPlayers(db))
		r.Get("/{id}", player_controller.GetPlayerByID(db))
		r.Patch("/{id}", player_controller.UpdatePlayer(db))
		r.Delete("/{id}", player_controller.DeletePlayer(db))

		r.Route("/{id}/player_sessions", func(r chi.Router) {
			r.Post("/", player_session_controller.CreatePlayerSession(db))
			r.Get("/{session_id}", player_session_controller.GetPlayerSessionByID(db))
			r.Patch("/{session_id}", player_session_controller.UpdatePlayerSession(db))
		})
	})

	router.Route("/stages", func(r chi.Router) {
		r.Post("/", stage_controller.CreateStage(db))
		r.Get("/", stage_controller.GetAllStages(db))
		r.Get("/{id}", stage_controller.GetStageByID(db))
		r.Patch("/{id}", stage_controller.UpdateStage(db))
		r.Delete("/{id}", stage_controller.DeleteStage(db))
	})

	return router
}