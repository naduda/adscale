package fileutils

import (
	"encoding/json"
	"io/ioutil"
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

func ChangeLineInFileByNumber(filename string, linenumber int, content string) error {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")

	for i, _ := range lines {
		if i == linenumber {
			lines[i] = content
		}
	}

	output := strings.Join(lines, "\n")
	return ioutil.WriteFile(filename, []byte(output), 0644)
}
