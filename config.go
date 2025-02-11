package main

import (
  "bufio"
  "encoding/json"
  "fmt"
  "os"
  "path/filepath"
  "strings"
)

// Config represents the structure of the configuration file
type Config struct {
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
	AccessToken    string `json:"access_token"`
	AccessSecret   string `json:"access_secret"`
}

const configFileName = "w2x.json"

func getConfigFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		os.Exit(1)
	}
	return filepath.Join(homeDir, ".config", configFileName)
}

func loadOrCreateConfig() (*Config, error) {
	configFilePath := getConfigFilePath()
	configDir := filepath.Dir(configFilePath)
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create config directory: %w", err)
		}
	}

	config := &Config{}
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		fmt.Println("Configuration file not found. Creating a new one...")
		if err := promptForConfigValues(config); err != nil {
			return nil, err
		}
		if err := saveConfig(config, configFilePath); err != nil {
			return nil, err
		}
	} else {
		file, err := os.Open(configFilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to open config file: %w", err)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		if err := decoder.Decode(config); err != nil {
			return nil, fmt.Errorf("failed to parse config file: %w", err)
		}

		if config.ConsumerKey == "" || config.ConsumerSecret == "" || config.AccessToken == "" || config.AccessSecret == "" {
			fmt.Println("Configuration file is incomplete. Prompting for missing values...")
			if err := promptForConfigValues(config); err != nil {
				return nil, err
			}
			if err := saveConfig(config, configFilePath); err != nil {
				return nil, err
			}
		}
	}
	return config, nil
}

func promptForConfigValues(config *Config) error {
	reader := bufio.NewReader(os.Stdin)
	if config.ConsumerKey == "" {
		fmt.Print("Enter Consumer Key: ")
		key, _ := reader.ReadString('\n')
		config.ConsumerKey = strings.TrimSpace(key)
	}
	if config.ConsumerSecret == "" {
		fmt.Print("Enter Consumer Secret: ")
		secret, _ := reader.ReadString('\n')
		config.ConsumerSecret = strings.TrimSpace(secret)
	}
	if config.AccessToken == "" {
		fmt.Print("Enter Access Token: ")
		token, _ := reader.ReadString('\n')
		config.AccessToken = strings.TrimSpace(token)
	}
	if config.AccessSecret == "" {
		fmt.Print("Enter Access Secret: ")
		secret, _ := reader.ReadString('\n')
		config.AccessSecret = strings.TrimSpace(secret)
	}
	return nil
}

func saveConfig(config *Config, configFilePath string) error {
	file, err := os.Create(configFilePath)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	return nil
}
