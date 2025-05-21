package main

import (
	"fmt"
	"go_server/internal/api/routes"
	"go_server/internal/utils"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  utils.InitDatabase()

  port := os.Getenv("PORT")

  if port == "" {
    log.Fatal("Error: PORT env is not set")
  }

  r := chi.NewRouter()

  r.Mount("/api/logs", routes.LogRouter())

  fmt.Printf("Server go is listening on http://localhost:%s", port)
  http.ListenAndServe(":" + port, r)
}
