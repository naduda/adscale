package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handlers) createRemoveImageFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var m map[string]bool
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	var err error

	if m["status"] {
		if err = h.Docker.CreateImage(h.Settings); err == nil {
			err = h.Docker.CreateAndRunContainer(h.Settings.AppPort)
		}
	} else {
		if err = h.Docker.StopAndRemoveContainer(); err == nil {
			err = h.Docker.RemoveImage()
		}
	}

	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *Handlers) createRemoveContainerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var m map[string]bool
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	var err error

	if m["status"] {
		err = h.Docker.CreateAndRunContainer(h.Settings.AppPort)
	} else {
		err = h.Docker.StopAndRemoveContainer()
	}

	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}
