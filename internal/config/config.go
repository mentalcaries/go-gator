package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	filePath, err := getConfigFilePath()

	if err != nil {
		fmt.Println("womp womp")
		return Config{}, err
	}
	data, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("womp womp")
		return Config{}, err
	}

	config := Config{}

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("could not get config object", err)
	}

	return config, nil
}

func (c *Config) SetUser(currentUser string) error {
	c.CurrentUserName = currentUser
	return write(*c)
}

func getConfigFilePath() (string, error) {
	homeLocation, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}
	configPath := filepath.Join(homeLocation, configFileName)

	return configPath, nil
}

func write(config Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return errors.New("invalid path")
	}

	configJson, _ := json.MarshalIndent(config, "", "  ")

	err = os.WriteFile(configPath, configJson, 0644)
	if err != nil {
		return err
	}

	return nil
}
