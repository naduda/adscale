package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetHandlers(api string, disable bool) map[string]http.HandlerFunc {
	handlers := map[string]http.HandlerFunc{}

	handlers["/"] = http.FileServer(http.Dir("/webapp")).ServeHTTP
	handlers[fmt.Sprintf("/%s/settings", api)] = settingsHandleFunc
	handlers[fmt.Sprintf("/%s/file-path-autocomplete", api)] = filePathAutocompleteHandleFunc
	handlers[fmt.Sprintf("/%s/properties", api)] = easyleadsPropertiesFunc
	handlers[fmt.Sprintf("/%s/technologies", api)] = getTechnologies
	handlers[fmt.Sprintf("/%s/prepare-docker-files", api)] = prepareDockerFiles

	if disable {
		for k, hf := range handlers {
			handlers[k] = disableCors(hf, true)
		}
	}

	fmt.Println(handlers)
	return handlers
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

// Needed in order to disable CORS for local development
func disableCors(h http.HandlerFunc, disable bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if disable {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.Header().Set("Access-Control-Allow-Headers", "*")
		}
		h(w, r)
	}
}
