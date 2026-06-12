# Opal

Opal is a high-performance command-line interface (CLI) developed in Go, designed for static repository analysis, automated technology detection, and Node.js dependency auditing. Built upon the Cobra framework, it provides a robust and scalable architecture for developer productivity.

## Features

- **Language Detection**: Automatic identification of project stacks based on file extensions, supporting over 20 languages and frameworks.
- **Frontend Ecosystem Analysis**: Specialized detection for modern frameworks (Next.js, Vite, Angular, Vue) and tooling (Tailwind, PostCSS, ESLint).
- **Dependency Auditing**: Deep scanning of Node.js projects to identify outdated packages and manage updates via npm.
- **Project Documentation**: Extensible command structure designed for automated README and technical documentation generation.

## Project Structure

```text
.
├── cmd/
│   └── opal/
│       └── main.go           # Application entry point
├── internal/
│   ├── cli/                  # Cobra command implementations
│   │   ├── deps.go           # 'dep' command logic
│   │   ├── init.go           # 'init' command logic
│   │   ├── readme.go         # 'readme' command logic
│   │   ├── root.go           # Root command configuration
│   │   └── version.go        # 'version' command logic
│   ├── deps/                 # Dependency analysis engine
│   └── scanner/              # File system scanning and language detection
├── testdata/                 # Mock data for automated testing
├── go.mod                    # Go module definition
└── README.md
```

## Requirements

- Go 1.21 or above.
- Node.js and npm.

## Installation

To build the binary locally:

```bash
go build -o opal ./cmd/main.go
```

This builds the CLI binary from the repository root.

If you prefer to install it into your Go bin directory, you can use:

```bash
go install ./cmd
```

## Usage

Run the binary from the project root or point it at another directory by changing into that directory first:

```bash
./cmd
./cmd init
./cmd dep
./cmd readme
./cmd version
```

Because the current implementation scans from `.` by default, the command should be launched from the directory you want to inspect.

## Example Project Layout

The repository includes a small frontend sample under `testdata/frontend/` that can be used to exercise the Node.js dependency flow.

## Current Limitations

- README generation is not implemented yet; the `readme` command is only a placeholder.
- Node dependency analysis shells out to `npm install`, `npm outdated`, and optionally `npm update`.
- Language detection is extension-based and does not parse source contents.

## License

See [LICENSE](LICENSE) for details.
