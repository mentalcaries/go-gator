package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)


type Config struct {
  DbUrl string `json:"db_url"`
}

func Read() Config {
  filePath, err := getConfigFilePath()

  if err != nil {
    fmt.Println("womp womp")
  }
  data, err := os.ReadFile(filePath)

  if err != nil {
    fmt.Println("womp womp")
  }

  var config Config

  err = json.Unmarshal(data, &config)
  if err != nil {
    fmt.Println("could not get config object", err)
  }

  return config
}

func getConfigFilePath()(string, error){
  homeLocation, err := os.UserHomeDir()
  const configFileName = ".gatorconfig.json"


  if err != nil {
    return "", errors.New("invalid path")
  }
  configPath := homeLocation + "/" + configFileName

  return configPath, nil
}