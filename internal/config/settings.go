package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// SettingFromFile will read the configuration from a JSON file.
func SettingFromFile(filename string) (*Settings, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("config setting file error: %w", err)
	}
	defer f.Close()

	settings := &Settings{}
	if err := json.NewDecoder(f).Decode(settings); err != nil {
		return nil, fmt.Errorf("config setting file decode error: %w", err)
	}
	return settings, nil
}
