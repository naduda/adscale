package docker

import (
	"adscale-tools/config"
	"adscale-tools/fileutils"
	"adscale-tools/model"
	"fmt"
	"os"
	"strings"
)

func PrepareEasyleadsConf() error {
	if err := fileutils.MakeDirIfNotExist("./docker"); err != nil {
		return err
	}

	var s model.Settings
	if err := fileutils.GetStructFromJsonFile(&s, model.SettingsFilePath); err != nil {
		return err
	}

	var config config.Config
	if err := config.Init(s.Easyleads); err != nil {
		return err
	}

	os.Remove(model.DockerEasyleadsConf)
	file, err := os.OpenFile(model.DockerEasyleadsConf, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	for k, v := range config.Props {
		if isPath(v.Value) {
			file.WriteString(fmt.Sprintf("%s=%s/%s\n", k, model.DockerDataFolder, k))
		} else {
			file.WriteString(fmt.Sprintf("%s=%s\n", k, v.Value))
		}
	}

	return nil
}

func isPath(t string) bool {
	t = strings.ToLower(t)
	return strings.HasPrefix(t, "c:\\") ||
		strings.HasPrefix(t, "c:/") ||
		strings.HasPrefix(t, "/")
}
