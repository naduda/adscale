package handlers

import (
	"adscale-tools/docker"
	"adscale-tools/model"
	"encoding/json"
	"fmt"
	"net/http"
)

type Handlers struct {
	Settings model.Settings
	Docker   docker.AdscaleContainer
}

func (h *Handlers) GetHandlers(api string, disable bool) map[string]http.HandlerFunc {
	handlers := map[string]http.HandlerFunc{
		"/": http.FileServer(http.Dir("/webapp")).ServeHTTP,
	}

	handlers[fmt.Sprintf("/%s/file-path-autocomplete", api)] = filePathAutocompleteHandleFunc
	handlers[fmt.Sprintf("/%s/settings", api)] = h.settingsHandleFunc
	handlers[fmt.Sprintf("/%s/properties", api)] = h.easyleadsPropertiesFunc
	handlers[fmt.Sprintf("/%s/add-property", api)] = h.addPropertyFunc
	handlers[fmt.Sprintf("/%s/remove-property", api)] = h.removePropertyFunc
	handlers[fmt.Sprintf("/%s/copy-properties-to-container", api)] = copyPropertiesToContainerFunc
	handlers[fmt.Sprintf("/%s/remove-extra-empty-lines", api)] = h.removeExtraEmptyLinesFunc
	handlers[fmt.Sprintf("/%s/docker-state", api)] = h.dockerStateFunc
	handlers[fmt.Sprintf("/%s/toggle-container", api)] = h.toggleContainerFunc
	handlers[fmt.Sprintf("/%s/create-remove-container", api)] = h.createRemoveContainerFunc
	handlers[fmt.Sprintf("/%s/create-remove-image", api)] = h.createRemoveImageFunc
	handlers[fmt.Sprintf("/%s/build-war", api)] = buildWarFunc
	handlers[fmt.Sprintf("/%s/update-frontend", api)] = h.updateFrontendFunc

	if disable {
		for k, hf := range handlers {
			handlers[k] = disableCors(hf, true)
		}
	}

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
