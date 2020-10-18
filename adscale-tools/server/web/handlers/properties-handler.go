package handlers

import (
	"adscale-tools/config"
	"adscale-tools/docker"
	"adscale-tools/fileutils"
	"adscale-tools/model"
	"adscale-tools/repo"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func (h *Handlers) easyleadsPropertiesFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error

	if r.Method == http.MethodGet {
		if properties, err := getProperties(h.Settings); err == nil {
			err = json.NewEncoder(w).Encode(properties)
		}
	}

	if r.Method == http.MethodPost {
		var properties map[string]config.ConfigurationProperty
		if err = json.NewDecoder(r.Body).Decode(&properties); err == nil {
			err = h.changeProperties(properties)
		}
	}

	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *Handlers) addPropertyFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error
	var data map[string]string

	if r.Method == http.MethodPost {
		if err = json.NewDecoder(r.Body).Decode(&data); err == nil {
			file, err := os.OpenFile(h.Settings.Easyleads, os.O_APPEND|os.O_WRONLY, 0666)
			if err != nil {
				sendErr(w, http.StatusInternalServerError, err.Error())
				return
			}
			defer file.Close()

			_, err = file.WriteString(fmt.Sprintf("%s=%s\n", data["name"], data["value"]))
		}
	}

	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *Handlers) removePropertyFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error
	var data map[string]int

	if r.Method == http.MethodPost {
		if err = json.NewDecoder(r.Body).Decode(&data); err == nil {
			err = fileutils.RemoveLineInFile(h.Settings.Easyleads, data["line"])
		}
	}

	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *Handlers) removeExtraEmptyLinesFunc(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error
	var conf config.Config

	if err = conf.Init(h.Settings.Easyleads); err == nil {
		if err = conf.Format(); err == nil {
			if properties, err := getProperties(h.Settings); err == nil {
				err = json.NewEncoder(w).Encode(properties)
			}
		}
	}

	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *Handlers) copyPropertiesToContainerFunc(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error
	if err = docker.PrepareEasyleadsConf(h.Settings); err == nil {
		cmd := fmt.Sprintf("docker cp ./%s %s:/adscale/configuration/easyleads.conf",
			model.DockerEasyleadsConf, model.DockerContainerName)
		err = docker.RunCommand(cmd, "./")
	}

	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *Handlers) changeProperties(properties map[string]config.ConfigurationProperty) error {
	input, err := ioutil.ReadFile(h.Settings.Easyleads)
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")

	for i := range lines {
		for k, v := range properties {
			if v.Line == i+1 {
				pref := ""
				if !v.Enabled {
					pref = "# "
				}
				lines[i] = fmt.Sprintf("%s%s=%s", pref, k, v.Value)
			}
		}
	}

	output := strings.Join(lines, "\n")
	return ioutil.WriteFile(h.Settings.Easyleads, []byte(output), 0644)
}

func getProperties(settings model.Settings) (map[string]*config.ConfigurationProperty, error) {
	var conf config.Config
	var err error

	if err = conf.Init(settings.Easyleads); err == nil {
		var r repo.Repository
		if err = r.Init(settings.Repo); err == nil {
			err = r.CheckProperties(&conf)
		}
	}

	return conf.Props, err
}
