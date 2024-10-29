package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (*Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}

	dat, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	cfg := Config{}
	err = json.Unmarshal(dat, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username

	return write(*cfg)
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configFilePath := filepath.Join(homeDir, configFileName)

	return configFilePath, nil
}

func write(cfg Config) error {
	dat, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	return os.WriteFile(configFilePath, dat, 0644)
}
