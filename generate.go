package main

import (
	"fmt"
	"os"
)

// Generate reads YAML, generates unit files, and writes them to the current directory
func Generate(yamlPath string) error {
	// Load and parse YAML configuration
	config, err := LoadConfig(yamlPath)
	if err != nil {
		return err
	}

	// Extract task name from YAML file name
	taskName := GetTaskName(yamlPath)

	// Generate unit file contents
	serviceContent := GenerateServiceUnit(config)
	timerContent := GenerateTimerUnit(config)

	// Get unit file names with prefix
	serviceName, timerName := GetUnitNames(taskName)

	// Write .service file to current directory
	if err := os.WriteFile(serviceName, []byte(serviceContent), 0644); err != nil {
		return fmt.Errorf("failed to write %s: %w", serviceName, err)
	}
	fmt.Printf("Generated: %s\n", serviceName)

	// Write .timer file to current directory
	if err := os.WriteFile(timerName, []byte(timerContent), 0644); err != nil {
		return fmt.Errorf("failed to write %s: %w", timerName, err)
	}
	fmt.Printf("Generated: %s\n", timerName)

	return nil
}
