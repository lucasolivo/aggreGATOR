package config

import (
    "encoding/json"
    "os"
    "path/filepath"
)
const configFileName = ".gatorconfig.json"

type Config struct {
    DBURL          string `json:"db_url"`
    CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
	return "", err
	}
	return filepath.Join(homeDir, configFileName), nil
}
		
		

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
    if err != nil {
        return Config{}, err
    }
	// Read the file
    data, err := os.ReadFile(filePath)
    if err != nil {
        return Config{}, err
    }
    
    // Parse JSON into the Config struct
    var cfg Config
    err = json.Unmarshal(data, &cfg)
    if err != nil {
        return Config{}, err
    }
    
    return cfg, nil
}

// SetUser updates the current user name and writes the updated config to file
func (cfg *Config) SetUser(username string) error {
    // Set the current user name
    cfg.CurrentUserName = username
    
    // Write the updated config to file
    return write(*cfg)
}

// write writes the config to the config file
func write(cfg Config) error {
    // Get the config file path
    filePath, err := getConfigFilePath()
    if err != nil {
        return err
    }
    
    // Marshal the config to JSON
    data, err := json.Marshal(cfg)
    if err != nil {
        return err
    }
    
    // Write the JSON data to the file
    return os.WriteFile(filePath, data, 0644)
}