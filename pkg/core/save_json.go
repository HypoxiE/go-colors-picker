package gocolorspicker

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

func FormConfigPath(path string) string {
	return filepath.Join(filepath.Dir(path), strings.TrimSuffix(filepath.Base(path), filepath.Ext(filepath.Base(path)))+".conf")
}

func SaveConfig(path string, config Configuration) error {
	conf, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	SaveJson(path, conf)
	return nil
}

func SaveJson(path string, config []byte) {
	if !strings.HasSuffix(path, ".conf") {
		path = FormConfigPath(path)
	}

	os.WriteFile(path, config, 0666)
}
