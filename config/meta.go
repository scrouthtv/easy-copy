package config

import (
	"os"
	"path/filepath"
)

var configLocs []string

func init() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	configLocs = []string{
		filepath.Join(configDir, "ec.conf"),
		filepath.Join(configDir, "ec/", "ec.conf"),
		filepath.Join(homeDir, "ec.conf"),
	}
}
