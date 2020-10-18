package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handlers) dockerStateFunc(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(h.Docker.State); err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *Handlers) toggleContainerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var m map[string]bool
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.Docker.ToggleContainer(m["status"]); err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}
