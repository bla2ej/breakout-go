# Project breakout-go

A simple and fun breakout game implemented in Go. This project demonstrates game development concepts and serves as a learning resource for Go enthusiasts. The intention of this project is to provide a hands-on example of how to build a game using Go, showcasing techniques such as rendering, collision detection, and game loop management.

## Getting Started

This section will guide you through setting up the project on your local machine for development and testing. The project uses a `Makefile` to simplify common tasks such as building, running, and testing the application.

### Prerequisites

Ensure you have the following installed on your system:
- [Go](https://golang.org/dl/) (latest stable version recommended)
- [Make](https://www.gnu.org/software/make/) (for running Makefile commands)

### Installation

1. Clone the repository:
  ```bash
  git clone https://github.com/bla2ej/breakout-go.git
  cd breakout-go
  ```

2. Install dependencies (if any):
  ```bash
  go mod tidy
  ```

## Makefile Commands

The `Makefile` provides several commands to streamline development. Below is a detailed explanation of each command:

### Build and Test
Run the build process and execute all tests to ensure the application is functioning correctly:
```bash
make all
```

### Build
Compile the application into a binary. The output binary will be located in the project directory:
```bash
make build
```

### Run
Execute the compiled application. This will launch the breakout game:
```bash
make run
```

### Live Reload
Enable live reloading during development. This is useful for making changes to the code and seeing the results immediately:
```bash
make watch
```

### Test
Run the test suite to verify the functionality of the application. This ensures that all components are working as expected:
```bash
make test
```

### Clean
Remove the binary and other build artifacts from the previous build. Use this command to clean up the project directory:
```bash
make clean
```

## Usage

Package breakout-web.go is the breakout game server. This server
is designed to host a breakout game that can be played either by a human player
or controlled by an AI. The primary intention of this project is to serve as a
learning resource for Go developers interested in game development and server-side
programming.

Usage:
- The server can be started by running the compiled binary or using the `make run`
  command if the Makefile is used.
- By default, the game operates in human player mode. To enable AI player mode,
  use the `-aibot` command-line flag when starting the server.
- The game interface is accessible via a web browser by navigating to the server's
  root URL (e.g., http://localhost:8080/).

Features:
- Human player mode: Allows a human player to control the paddle using keyboard input.
- AI player mode: Enables an AI to control the paddle, useful for testing or experimentation.
- HTTP API: Provides endpoints for interacting with the game state, resetting the game,
  and retrieving AI-specific game data.

Command-Line Flags:
- `-aibot`: A boolean flag to enable AI player mode. Defaults to `false` (human player mode).

HTTP Endpoints:
- `GET /`: Serves the static HTML file for the game interface.
- `POST /reset`: Resets the game state to its initial configuration.
- `POST /game-state`: Updates the game state based on player input and returns the
  current game state as JSON.
- `POST /ai-state`: Updates the game state based on AI input and returns the AI-specific
  game state, including action, reward, and game status.

Environment Variables:
- `PORT`: Specifies the port on which the server listens. Defaults to `8080` if not set.

This project is ideal for:
- Learning Go by exploring a practical example of game development.
- Understanding server-side programming concepts, including HTTP handlers and state management.
- Experimenting with AI integration in games and extending the game with new features.
