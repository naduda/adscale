package handlers

import (
	"adscale-tools/fileutils"
	"adscale-tools/model"
	"encoding/json"
	"net/http"
	"os"
)

func settingsHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s model.Settings
	var err error

	if r.Method == http.MethodGet {
		if err = fileutils.GetStructFromJsonFile(&s, model.SettingsFilePath); err == nil {
			err = json.NewEncoder(w).Encode(s)
		}
	}

	if r.Method == http.MethodPost {
		if err = json.NewDecoder(r.Body).Decode(&s); err == nil {
			if err = isSettingsValid(s); err == nil {
				err = fileutils.SaveStructToJsonFile(s, model.SettingsFilePath)
			}
		}
	}

	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func isSettingsValid(s model.Settings) error {
	var err error
	if _, err = os.Stat(s.Easyleads); err == nil {
		_, err = os.Stat(s.Repo)
	}
	return err
}
