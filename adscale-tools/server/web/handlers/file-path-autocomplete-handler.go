package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type IPath struct {
	Path string `json:"path"`
}

func filePathAutocompleteHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var path IPath

	if r.Method == http.MethodPost {
		var err error
		if err = json.NewDecoder(r.Body).Decode(&path); err == nil {
			if list, err := listOfFiles(path.Path); err == nil {
				err = json.NewEncoder(w).Encode(list)
			}
		}

		if err != nil {
			sendErr(w, http.StatusInternalServerError, err.Error())
		}
	}
}

func listOfFiles(path string) ([]string, error) {
	suffix := ""
	if _, err := os.Stat(path); err != nil {
		idx := strings.LastIndex(path, "/") + 1
		suffix = strings.ToLower(path[idx:])
		path = path[:idx]
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var res []string
	for _, f := range files {
		fName := strings.ToLower(f.Name())
		if strings.HasPrefix(fName, suffix) {
			res = append(res, f.Name())
		}
	}

	return res, nil
}
