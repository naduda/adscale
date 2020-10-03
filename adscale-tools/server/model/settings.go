package model

const SettingsFilePath = "./settings.json"
const DockerEasyleadsConf = "./docker/easyleads.conf"
const DockerDataFolder = "/adscale"

type Settings struct {
	Easyleads string `json:"easyleads"`
	Repo      string `json:"repo"`
}
