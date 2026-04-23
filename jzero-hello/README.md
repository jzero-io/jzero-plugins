# jzero-hello Plugin

A simple example plugin for jzero that demonstrates basic plugin functionality.

## Features

- Display plugin descriptor information (API, Proto, Model files)
- Modular command structure using cobra

## Installation

```bash
# Build the plugin
go build -o jzero-hello main.go
mv jzero-hello $GOBIN
```

## Usage

```bash
# Show help
jzero-hello -h

# Display descriptor information
jzero-hello desc
```

## Development

This plugin uses:
- **cobra** for CLI command structure
- **Standard Go libraries** for functionality

## License

Same as jzero project
