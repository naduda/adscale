package model

import (
	"adscale-tools/fileutils"
	"os"
)

const settingsFilePath = "./settings.json"
const SettingsFilePath = "./settings.json"

type Settings struct {
	Easyleads string `json:"easyleads"`
	Repo      string `json:"repo"`
	UiFolder  string `json:"ui"`
	DbIP      string `json:"dbIP"`
	AppPort   int    `json:"appPort"`
}

func (s *Settings) Read() error {
	return fileutils.GetStructFromJsonFile(&s, settingsFilePath)
}

func (s *Settings) Save() error {
	if err := s.isValid(); err != nil {
		return err
	}
	return fileutils.SaveStructToJsonFile(s, settingsFilePath)
}

func (s *Settings) isValid() error {
	if _, err := os.Stat(s.Easyleads); err != nil {
		return err
	}

	if _, err := os.Stat(s.Repo); err != nil {
		return err
	}

	if _, err := os.Stat(s.UiFolder); err != nil {
		return err
	}

	return nil
}
