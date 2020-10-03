package handlers

import (
	"adscale-tools/docker"
	"net/http"
)

func prepareDockerFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	docker.PrepareEasyleadsConf()
}
