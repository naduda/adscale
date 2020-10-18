package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handlers) settingsHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var err error

	if r.Method == http.MethodGet {
		err = json.NewEncoder(w).Encode(h.Settings)
	}

	if r.Method == http.MethodPost {
		if err = json.NewDecoder(r.Body).Decode(&h.Settings); err == nil {
			err = h.Settings.Save()
		}
	}

	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}
