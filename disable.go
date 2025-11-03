package main

import (
	"fmt"
)

// Disable reads YAML, disables the timer, and removes unit files
func Disable(yamlPath string) error {
	// Load YAML to get task name (validation not strictly required for disable)
	config, err := LoadConfig(yamlPath)
	if err != nil {
		return err
	}

	// Suppress unused variable warning (we load config for future extensibility)
	_ = config

	// Extract task name from YAML file name
	taskName := GetTaskName(yamlPath)

	// Get unit file paths and names
	servicePath, timerPath := GetUnitPaths(taskName)
	_, timerName := GetUnitNames(taskName)

	// Disable and stop the timer
	if err := DisableTimer(timerName); err != nil {
		return err
	}
	fmt.Printf("Executed: systemctl disable --now %s\n", timerName)

	// Remove .timer file
	if err := RemoveUnitFile(timerPath); err != nil {
		return err
	}
	fmt.Printf("Removed: %s\n", timerPath)

	// Remove .service file
	if err := RemoveUnitFile(servicePath); err != nil {
		return err
	}
	fmt.Printf("Removed: %s\n", servicePath)

	// Reload systemd daemon
	if err := DaemonReload(); err != nil {
		return err
	}
	fmt.Println("Executed: systemctl daemon-reload")

	return nil
}
