package handlers

import (
	"adscale-tools/docker"
	"adscale-tools/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handlers) buildWarFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var m map[string]bool
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := docker.CreateWar(m["installModule"], m["installCbfsms"], h.Settings); err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *Handlers) updateFrontendFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var m map[string]bool
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	isDev := m["isDev"]
	installNpm := m["installNpm"]

	folderName := "policy"
	if isDev {
		folderName += "-dev"
	}

	npmCmd := ""
	if installNpm {
		npmCmd = "npm install && "
	}

	ngCommand := "build"
	if isDev {
		ngCommand = "build-dev"
	}
	cmd := fmt.Sprintf("%snpm run %s --prefix %s", npmCmd, ngCommand, h.Settings.UiFolder)

	cmd += fmt.Sprintf(" && docker exec %s rm -rf /usr/local/tomcat/webapps/ROOT/%s", model.DockerContainerName, folderName)

	cmd += fmt.Sprintf(" && docker cp %s %s:/usr/local/tomcat/webapps/ROOT/%s",
		h.Settings.UiFolder+"/dist/adscale", model.DockerContainerName, folderName)

	if err := docker.RunCommand(cmd, h.Settings.UiFolder); err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
}
