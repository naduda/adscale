package handlers

import (
	"adscale-tools/model"
	"encoding/json"
	"net/http"
)

func getTechnologies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	technologies := []model.Technology{
		{"Tech1", "Details1"},
		{"Tech2", "Details2"},
	}
	err := json.NewEncoder(w).Encode(technologies)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}
