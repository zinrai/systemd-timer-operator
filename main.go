package main

import (
	"fmt"
	"os"
)

const version = "0.2.0"

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 2 {
		printUsage()
		return fmt.Errorf("no subcommand specified")
	}

	subcommand := os.Args[1]

	switch subcommand {
	case "enable":
		if len(os.Args) < 3 {
			return fmt.Errorf("usage: systemd-timer-operator enable <yaml-file>")
		}
		yamlPath := os.Args[2]
		return Enable(yamlPath)

	case "disable":
		if len(os.Args) < 3 {
			return fmt.Errorf("usage: systemd-timer-operator disable <yaml-file>")
		}
		yamlPath := os.Args[2]
		return Disable(yamlPath)

	case "generate":
		if len(os.Args) < 3 {
			return fmt.Errorf("usage: systemd-timer-operator generate <yaml-file>")
		}
		yamlPath := os.Args[2]
		return Generate(yamlPath)

	case "version":
		fmt.Printf("systemd-timer-operator version %s\n", version)
		return nil

	case "help":
		printUsage()
		return nil

	default:
		printUsage()
		return fmt.Errorf("unknown subcommand: %s", subcommand)
	}
}

func printUsage() {
	usage := `systemd-timer-operator - Manage systemd timers with YAML configuration

Usage:
  systemd-timer-operator <subcommand> [arguments]

Subcommands:
  enable <yaml-file>    Generate unit files, deploy them, and enable the timer
  disable <yaml-file>   Disable the timer and remove unit files
  generate <yaml-file>  Generate unit files in the current directory
  version               Show version information
  help                  Show this help message
`
	fmt.Print(usage)
}
