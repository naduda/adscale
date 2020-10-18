package fileutils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
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

func RemoveLineInFile(filename string, line int) error {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")

	for i := range lines {
		if line-1 == i {
			lines[i] = ""
		}
	}

	output := strings.Join(lines, "\n")
	return ioutil.WriteFile(filename, []byte(output), 0644)
}

func ReplaceInFile(filename string, old string, text string) error {
	if old == text {
		return nil
	}

	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")

	for i := range lines {
		line := lines[i]
		if strings.Contains(line, old) {
			lines[i] = strings.Replace(line, old, text, -1)
		}
	}

	output := strings.Join(lines, "\n")
	return ioutil.WriteFile(filename, []byte(output), 0644)
}
