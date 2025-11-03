# systemd-timer-operator

A command-line tool to manage systemd timers with YAML configuration files. This tool consolidates the `.service` and `.timer` file pair into a single YAML file and reduces multiple systemctl operations to a single command.

## Motivation

Managing `systemd` timers requires the following operations:

1. Create a `.service` file to define the task execution
2. Create a `.timer` file to define the schedule
3. Place both files in the appropriate directory
4. Run `systemctl daemon-reload`
5. Run `systemctl enable --now <timer-name>.timer`

This tool reduces this workflow to a single YAML file and a single command, while leveraging native systemd capabilities.

## Features

- **Single YAML file**: Define both service and timer configuration in one place
- **Direct systemd mapping**: YAML structure mirrors native systemd unit file format
- **Full systemd access**: All systemd features available via official documentation
- **Single command operation**: One command to enable or disable timers

## Installation

```bash
$ go install github.com/zinrai/systemd-timer-operator@latest
```

## Usage

Enable a timer:

```bash
$ sudo systemd-timer-operator enable examples/basic.yml
```

Disable a timer:

```bash
$ sudo systemd-timer-operator disable examples/basic.yml
```

For YAML configuration examples, see the `examples/` directory.

## YAML Configuration Reference

### Basic Configuration

| Field          | Description                                          | Required     |
|----------------|------------------------------------------------------|--------------|
| `Unit`         | Common `[Unit]` section for both .service and .timer | Optional     |
| `Service`      | `[Service]` section content                          | **Required** |
| `Timer`        | `[Timer]` section content                            | **Required** |
| `TimerInstall` | `[Install]` section for .timer file                  | Optional     |

### Dedicated Configuration

For use cases requiring different `[Unit]` sections for .service and .timer files:

| Field            | Description                                                     |
|------------------|-----------------------------------------------------------------|
| `ServiceUnit`    | Dedicated `[Unit]` section for .service file (overrides `Unit`) |
| `TimerUnit`      | Dedicated `[Unit]` section for .timer file (overrides `Unit`)   |
| `ServiceInstall` | `[Install]` section for .service file                           |

## How It Works

1. **YAML to Unit File Mapping**: The tool reads the YAML configuration and generates two systemd unit files (.service and .timer) by directly mapping YAML fields to INI-format sections
2. **Fallback Logic**: If `ServiceUnit` or `TimerUnit` fields are not specified, the common `Unit` field value is used for the respective `[Unit]` sections
3. **File Placement**: Generated unit files are placed in `/etc/systemd/system/` with the prefix `systemd-timer-operator-` to avoid naming conflicts
4. **systemd Registration**: The tool executes `systemctl daemon-reload` to register new units, then `systemctl enable --now` to activate the timer
5. **Transparent Output**: All systemctl command output is displayed as-is, without modification or filtering
6. **Validation Delegation**: Configuration syntax and semantic validation are delegated to systemd. The tool only validates the presence of required YAML fields (`Service` and `Timer`)

## License

This project is licensed under the [MIT License](./LICENSE).
