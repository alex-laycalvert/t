package config

import (
	"alex-laycalvert/t/internal/utils"
	"bufio"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"strings"
)

const defaultConfigPath = "~/.config/t/config"

const (
	DBPathKey ConfigKey = "database_path"
)

var defaultConfig = map[ConfigKey]string{
	"database_path": "~/.local/share/t/t.db",
}

type Config struct {
	FilePath string
	config   map[ConfigKey]string
}

func New() (*Config, error) {
	configFilePath := FilePath()

	fileValues, err := readConfigFile(configFilePath)
	if err != nil {
		return nil, err
	}

	config := &Config{
		FilePath: configFilePath,
		config:   fileValues,
	}
	dbPath := config.Get(DBPathKey)
	dbPath = utils.ExpandHomeDir(dbPath)

	err = os.MkdirAll(filepath.Dir(dbPath), 0755)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) Get(key ConfigKey) string {
	v, ok := c.config[key]
	if ok {
		return v
	}

	return ""
}

func (c *Config) Set(key ConfigKey, value string) error {
	c.config[key] = value
	return writeConfigFile(c.FilePath, c.config)
}

// FilePath returns the resolved path to the config file.
// Uses T_CONFIG_FILE env var if set, otherwise defaults to ~/.config/t/config.
func FilePath() string {
	if envPath := os.Getenv("T_CONFIG_FILE"); envPath != "" {
		return utils.ExpandHomeDir(envPath)
	}
	return utils.ExpandHomeDir(defaultConfigPath)
}

func readConfigFile(path string) (map[ConfigKey]string, error) {
	values := make(map[ConfigKey]string)
	maps.Copy(values, defaultConfig)

	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return values, nil
		}
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, found := strings.Cut(line, "=")
		if !found {
			continue
		}
		key = strings.TrimSpace(key)
		if !IsConfigKey(key) {
			return values, fmt.Errorf("invalid key '%s'", key)
		}
		confKey := ConfigKey(key)

		value = strings.TrimSpace(value)
		values[confKey] = value
	}

	return values, scanner.Err()
}

func writeConfigFile(path string, values map[ConfigKey]string) error {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for key, value := range values {
		_, err := fmt.Fprintf(w, "%s = %s\n", key, value)
		if err != nil {
			return err
		}
	}

	return w.Flush()
}

func IsConfigKey(key string) bool {
	return key == DBPathKey.String()
}

func ValidKeys() []string {
	return []string{DBPathKey.String()}
}
