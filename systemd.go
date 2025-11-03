package main

import (
	"fmt"
	"os"
	"os/exec"
)

// DaemonReload executes systemctl daemon-reload
func DaemonReload() error {
	cmd := exec.Command("systemctl", "daemon-reload")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("systemctl daemon-reload failed: %w", err)
	}
	return nil
}

// EnableTimer enables and starts the timer
func EnableTimer(timerName string) error {
	cmd := exec.Command("systemctl", "enable", "--now", timerName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("systemctl enable --now %s failed: %w", timerName, err)
	}
	return nil
}

// DisableTimer disables and stops the timer
func DisableTimer(timerName string) error {
	cmd := exec.Command("systemctl", "disable", "--now", timerName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("systemctl disable --now %s failed: %w", timerName, err)
	}
	return nil
}
