package handlers

import (
	"adscale-tools/config"
	"adscale-tools/fileutils"
	"adscale-tools/model"
	"adscale-tools/repo"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func easyleadsPropertiesFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var settings model.Settings
	var err error

	if r.Method == http.MethodGet {
		if err = fileutils.GetStructFromJsonFile(&settings, model.SettingsFilePath); err == nil {
			config := config.Config{}
			config.Init(settings.Easyleads)
			config.Format()

			repo := repo.Repository{}
			repo.Init(settings.Repo)

			repo.CheckProperties(&config)

			err = json.NewEncoder(w).Encode(config.Props)
		}
	}

	if r.Method == http.MethodPost {
		var properties map[string]config.ConfigurationProperty
		if err = json.NewDecoder(r.Body).Decode(&properties); err == nil {
			err = changeProperties(properties)
		}
	}

	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func changeProperties(properties map[string]config.ConfigurationProperty) error {
	var settings model.Settings
	if err := fileutils.GetStructFromJsonFile(&settings, model.SettingsFilePath); err != nil {
		return err
	}

	input, err := ioutil.ReadFile(settings.Easyleads)
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")

	for i, _ := range lines {
		for k, v := range properties {
			if v.Line == i+1 {
				lines[i] = fmt.Sprintf("%s=%s", k, v.Value)
			}
		}
	}

	output := strings.Join(lines, "\n")
	return ioutil.WriteFile(settings.Easyleads, []byte(output), 0644)
}
