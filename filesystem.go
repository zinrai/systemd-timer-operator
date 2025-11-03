package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	SystemdUnitDir = "/etc/systemd/system"
	UnitPrefix     = "systemd-timer-operator"
)

// GetTaskName extracts the task name from YAML file path
func GetTaskName(yamlPath string) string {
	base := filepath.Base(yamlPath)
	return strings.TrimSuffix(base, filepath.Ext(base))
}

// GetUnitNames generates the .service and .timer unit names
func GetUnitNames(taskName string) (string, string) {
	serviceName := fmt.Sprintf("%s-%s.service", UnitPrefix, taskName)
	timerName := fmt.Sprintf("%s-%s.timer", UnitPrefix, taskName)
	return serviceName, timerName
}

// GetUnitPaths generates the full paths for .service and .timer files
func GetUnitPaths(taskName string) (string, string) {
	serviceName, timerName := GetUnitNames(taskName)
	servicePath := filepath.Join(SystemdUnitDir, serviceName)
	timerPath := filepath.Join(SystemdUnitDir, timerName)
	return servicePath, timerPath
}

// WriteUnitFile writes content to a unit file
func WriteUnitFile(path, content string) error {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write %s: %w", path, err)
	}
	return nil
}

// RemoveUnitFile removes a unit file
func RemoveUnitFile(path string) error {
	if err := os.Remove(path); err != nil {
		return fmt.Errorf("failed to remove %s: %w", path, err)
	}
	return nil
}
