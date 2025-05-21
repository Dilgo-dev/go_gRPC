package handlers

import (
	"encoding/json"
	"go_server/internal/utils"
	"net/http"
)

func GetLogs(w http.ResponseWriter, r *http.Request) {
  db := utils.Database()
  var logs utils.Log

  db.Find(&logs)
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(logs)
}
