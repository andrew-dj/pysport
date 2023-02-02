package api

import (
	"encoding/json"
	"github.com/sportorg/online/database"
	"github.com/sportorg/online/database/model"
	"net/http"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode([]User{})
}

func getRaces(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(model.GetRaces(database.DB))
}