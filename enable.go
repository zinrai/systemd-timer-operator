package main

import (
	"fmt"
)

// Enable reads YAML, generates unit files, deploys them, and enables the timer
func Enable(yamlPath string) error {
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

	// Get unit file paths
	servicePath, timerPath := GetUnitPaths(taskName)
	_, timerName := GetUnitNames(taskName)

	// Write .service file
	if err := WriteUnitFile(servicePath, serviceContent); err != nil {
		return err
	}
	fmt.Printf("Created: %s\n", servicePath)

	// Write .timer file
	if err := WriteUnitFile(timerPath, timerContent); err != nil {
		return err
	}
	fmt.Printf("Created: %s\n", timerPath)

	// Reload systemd daemon
	if err := DaemonReload(); err != nil {
		return err
	}
	fmt.Println("Executed: systemctl daemon-reload")

	// Enable and start the timer
	if err := EnableTimer(timerName); err != nil {
		return err
	}
	fmt.Printf("Executed: systemctl enable --now %s\n", timerName)

	return nil
}
