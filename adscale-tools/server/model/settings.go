package model

const SettingsFilePath = "./settings.json"

type Settings struct {
	Easyleads string `json:"easyleads"`
	Repo      string `json:"repo"`
}
