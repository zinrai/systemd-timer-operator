package main

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

// TimerConfig represents the YAML configuration structure
type TimerConfig struct {
	// Common Unit section for both .service and .timer
	Unit string `yaml:"Unit,omitempty"`

	// Service-specific sections
	ServiceUnit    string `yaml:"ServiceUnit,omitempty"`
	Service        string `yaml:"Service"`
	ServiceInstall string `yaml:"ServiceInstall,omitempty"`

	// Timer-specific sections
	TimerUnit    string `yaml:"TimerUnit,omitempty"`
	Timer        string `yaml:"Timer"`
	TimerInstall string `yaml:"TimerInstall,omitempty"`
}

// LoadConfig loads and parses a YAML configuration file
func LoadConfig(path string) (*TimerConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var config TimerConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &config, nil
}

// Validate checks if required fields are present
func (c *TimerConfig) Validate() error {
	if c.Service == "" {
		return fmt.Errorf("required field 'Service' is missing")
	}
	if c.Timer == "" {
		return fmt.Errorf("required field 'Timer' is missing")
	}
	return nil
}

// GetServiceUnit returns the Unit section content for .service file
func (c *TimerConfig) GetServiceUnit() string {
	if c.ServiceUnit != "" {
		return c.ServiceUnit
	}
	return c.Unit
}

// GetTimerUnit returns the Unit section content for .timer file
func (c *TimerConfig) GetTimerUnit() string {
	if c.TimerUnit != "" {
		return c.TimerUnit
	}
	return c.Unit
}
