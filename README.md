# PortOps

---

PortOps is a command-line interface (CLI) tool for managing processes and ports across different operating systems. It provides functionality to kill processes by their process IDs (PIDs) and release ports that are in use.

## Installation

To install PortOps, clone this repo and cd into the repo

```bash
go build
```

## Usage

PortOps provides three main commands:

### Kill

The kill command is used to terminate a process by its PID.

```bash
portops kill [pid]
```

Replace `[pid]` with the PID of the process you want to terminate.

Release
The release command is used to release a port that is curre

```bash
portops release [port]
```

Replace `[port]` with the port number you want to release.

### Supported Platforms

PortOps supports Windows, macOS, and Linux operating systems.

### Contributing

Contributions to PortOps are welcome! If you encounter any issues or have suggestions for improvements, feel free to open an issue or submit a pull request on GitHub.
