package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3" // Using yaml for config, could also use JSON or TOML
)

// AppConfig holds the application's configuration settings.
type AppConfig struct {
	ReportDir string `yaml:"reportDir"` // Directory where reports should be saved
	SchemaDir string `yaml:"schemaDir"` // Directory for custom risk/control schemas
	// Add other global configuration parameters here
	// Example:
	// DefaultControlCatalog string `yaml:"defaultControlCatalog"`
}

// LoadAppConfig loads the application configuration from a specified file path.
// If filePath is empty, it attempts to load from default locations (e.g., ~/.opencrace.yaml).
func LoadAppConfig(filePath string) (*AppConfig, error) {
	var cfg AppConfig
	var configData []byte
	var err error

	if filePath != "" {
		configData, err = os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read config file '%s': %w", filePath, err)
		}
	} else {
		// Try default locations
		homeDir, err := os.UserHomeDir()
		if err == nil {
			defaultPath := filepath.Join(homeDir, ".opencrace.yaml")
			if _, err := os.Stat(defaultPath); err == nil {
				configData, err = os.ReadFile(defaultPath)
				if err != nil {
					return nil, fmt.Errorf("failed to read default config file '%s': %w", defaultPath, err)
				}
				fmt.Printf("Loaded default config from: %s\n", defaultPath)
			}
		}
	}

	if len(configData) > 0 {
		err = yaml.Unmarshal(configData, &cfg)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal config: %w", err)
		}
	} else {
		fmt.Println("No config file loaded, using default settings.")
		// Set sensible defaults if no config file is found
		cfg.ReportDir = "reports"
		cfg.SchemaDir = "static/schemas" // Assuming schemas are relative to executable or module root
	}

	// Ensure report directory exists
	if _, err := os.Stat(cfg.ReportDir); os.IsNotExist(err) {
		err = os.MkdirAll(cfg.ReportDir, 0755) // Create with read/write for owner, read for others
		if err != nil {
			return nil, fmt.Errorf("failed to create report directory '%s': %w", cfg.ReportDir, err)
		}
	}


	return &cfg, nil
}