package main

import (
	"fmt"
	"strings"
)

// GenerateServiceUnit generates the content of .service file
func GenerateServiceUnit(config *TimerConfig) string {
	var sections []string

	// [Unit] section
	if unitContent := config.GetServiceUnit(); unitContent != "" {
		sections = append(sections, formatSection("Unit", unitContent))
	}

	// [Service] section
	sections = append(sections, formatSection("Service", config.Service))

	// [Install] section (optional)
	if config.ServiceInstall != "" {
		sections = append(sections, formatSection("Install", config.ServiceInstall))
	}

	return strings.Join(sections, "\n")
}

// GenerateTimerUnit generates the content of .timer file
func GenerateTimerUnit(config *TimerConfig) string {
	var sections []string

	// [Unit] section
	if unitContent := config.GetTimerUnit(); unitContent != "" {
		sections = append(sections, formatSection("Unit", unitContent))
	}

	// [Timer] section
	sections = append(sections, formatSection("Timer", config.Timer))

	// [Install] section
	if config.TimerInstall != "" {
		sections = append(sections, formatSection("Install", config.TimerInstall))
	}

	return strings.Join(sections, "\n")
}

// formatSection formats a systemd unit section
func formatSection(name, content string) string {
	// Trim leading/trailing whitespace
	content = strings.TrimSpace(content)

	// If content is empty, return empty string
	if content == "" {
		return ""
	}

	// Ensure content doesn't end with multiple newlines
	content = strings.TrimRight(content, "\n")

	return fmt.Sprintf("[%s]\n%s\n", name, content)
}
