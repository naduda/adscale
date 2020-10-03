package fileutils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func GetStructFromJsonFile(data interface{}, filename string) error {
	if content, err := ioutil.ReadFile(filename); err != nil {
		return err
	} else {
		return json.Unmarshal(content, &data)
	}
}

func SaveStructToJsonFile(data interface{}, filename string) error {
	if content, err := json.MarshalIndent(data, "", " "); err != nil {
		return err
	} else {
		return ioutil.WriteFile(filename, content, 0644)
	}
}

func MakeDirIfNotExist(name string) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return os.Mkdir(name, os.ModePerm)
	}
	return nil
}
