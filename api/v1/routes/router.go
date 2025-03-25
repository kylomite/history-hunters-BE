package routes

import (
	"database/sql"
	"net/http"

	"historyHunters/api/v1/controllers/player_controller"
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
		r.Get("/", player_controller.GetAllPlayers(db))  
		r.Get("/{id}", player_controller.GetPlayerByID(db))
		r.Patch("/{id}", player_controller.UpdatePlayer(db))
		r.Delete("/{id}", player_controller.DeletePlayer(db))
	})

	return router
}