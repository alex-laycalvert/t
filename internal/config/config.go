package config

import (
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	DatabasePath string
}

func New() (*Config, error) {
	dataDirPath, err := expandHomeDir("~/.local/share/t")
	if err != nil {
		return nil, err
	}
	dbPath := dataDirPath + "/t.db"

	err = os.MkdirAll(dataDirPath, 0755)
	if err != nil {
		return nil, err
	}

	return &Config{
		// TODO: make configurable at ~/.config/t/config
		DatabasePath: dbPath,
	}, nil
}

func expandHomeDir(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, path[1:]), nil
	}
	return path, nil
}
