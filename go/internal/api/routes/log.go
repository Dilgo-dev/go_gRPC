package routes

import (
	"go_server/internal/api/handlers"

	"github.com/go-chi/chi"
)

func LogRouter() *chi.Mux {
  r := chi.NewRouter()

  r.Get("/", handlers.GetLogs)

  return r
}
