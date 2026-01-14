# Perimeter

A CLI tool for analyzing JavaScript and TypeScript projects.

## Overview

Perimeter scans and analyzes JavaScript/TypeScript codebases, providing insights and reporting capabilities. It supports projects using Node.js (requires `package.json`).

## Installation

```bash
go build -o perimeter ./cmd/perimeter
```

## Usage

```bash
perimeter --path /path/to/project
```

### Options

- `--path`: Path to project root (default: current directory)

## Requirements

- Go 1.21 or later
- Target project must contain a `package.json` file in the root directory

## Supported File Types

Perimeter recognizes the following source file extensions:
- `.js` - JavaScript
- `.jsx` - JavaScript React
- `.ts` - TypeScript
- `.tsx` - TypeScript React

## Project Structure

```
perimeter/
├── cmd/perimeter/    # CLI entry point
├── internal/
│   ├── agents/       # Agent implementations
│   ├── cli/          # CLI flag parsing
│   ├── helpers/      # Utility functions
│   ├── index/        # File indexing and scanning
│   ├── llm/          # LLM integration
│   ├── logx/         # Logging utilities
│   ├── reporter/     # Reporting functionality
│   └── types/        # Type definitions
└── go.mod
```

## Development

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build ./cmd/perimeter
```

## License

[Add your license here]
