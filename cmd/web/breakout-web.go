package main

import (
	"breakout-go/internal/breakout"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
)

//go:embed index.html
var indexHTML []byte

// main is the entry point of the breakout game server. It initializes the game,
// parses command-line flags, and sets up HTTP handlers for serving the game
// and managing its state.
//
// The server supports two modes:
// - Human player mode: Allows a human player to control the paddle.
// - AI player mode: Allows an AI to control the paddle.
//
// Command-line flags:
// - -aibot: A boolean flag to enable AI player mode. Defaults to false (human player mode).
//
// The following HTTP endpoints are provided:
//   - "/" (GET): Serves the static HTML file for the game interface.
//   - "/reset" (POST): Resets the game state to its initial configuration.
//   - "/game-state" (POST): Updates the game state based on player input and
//     returns the current game state as JSON.
//   - "/ai-state" (POST): Updates the game state based on AI input and returns
//     the AI-specific game state, including action, reward, and game status.
//
// The server listens on a port specified by the PORT environment variable.
// If the PORT variable is not set, it defaults to port 8080.
func main() {
	port := "8080"
	aibot := flag.Bool("aibot", false, "Run as AI player. Defaults to human player.")
	flag.Parse()
	humanPlayer := !*aibot
	if humanPlayer {
		fmt.Println("Running in human player mode")
	} else {
		fmt.Println("Running in AI player mode")
	}

	game := breakout.NewBreakout()

	// Serve the static HTML file
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// replace within data file string 8080 with port
		s := string(indexHTML)
		s = strings.ReplaceAll(s, "8080", port)
		if !humanPlayer {
			s = strings.ReplaceAll(s, "humanplay=1", "humanplay=0")
		}
		data := []byte(s)

		w.Header().Set("Content-Type", "text/html")
		w.Write(data)
	})

	// reset the game state
	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		game = breakout.NewBreakout()
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Game reset"})
	})

	// Handle game state updates via POST requests
	http.HandleFunc("/game-state", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Parse the form data
			var input struct {
				Left  bool `json:"left"`
				Right bool `json:"right"`
			}

			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
				return
			}

			// Log the action received from the client
			if humanPlayer {
				if input.Left && !input.Right {
					game.PaddleLeft()
				} else if input.Right && !input.Left {
					game.PaddleRight()
				}
				game.MoveBall()
			}
			// time.Sleep(20 * time.Millisecond)
		}
		// Serve the game state as JSON
		state := game.GetState()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(state)
	})

	// serve AI player
	// add AI handle at /ai-state
	http.HandleFunc("/ai-state", func(w http.ResponseWriter, r *http.Request) {
		action := 0
		score := game.GetState().Score
		if r.Method == http.MethodPost {
			// Parse the form data
			var input struct {
				Action int `json:"action"`
			}
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
				return
			}
			// Log the action received from the client
			// only action 1 and 2 is valid to move paddle
			// action 0 is no action
			if !humanPlayer {
				action = input.Action
				if action == 1 {
					game.PaddleLeft()
				} else if action == 2 {
					game.PaddleRight()
				}
				game.MoveBall()
			}
		}

		var aiState struct {
			Action int     `json:"action"`
			Reward float64 `json:"reward"`
			State  [][]int `json:"state"`
			Done   bool    `json:"done"`
			Lives  int     `json:"lives"`
		}
		// Serve the game state as JSON
		state := game.GetState()
		aiState.State = BreakoutState2Bitmap(&state)
		aiState.Action = action
		if state.Score > score {
			aiState.Reward = 1.0
		} else {
			aiState.Reward = 0.0
		}
		aiState.Done = state.Done
		aiState.Lives = 5 - state.Live + 1
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(aiState)
	})

	// Start the server
	// read port from env var or use default port 8080
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		os.Exit(1)
	}
}

// BreakoutState2Bitmap converts the state of a Breakout game into a 2D bitmap representation.
// Each element in the bitmap corresponds to a specific part of the game:
// - 0: Empty space
// - 1: Yellow brick
// - 2: Green brick
// - 3: Orange brick
// - 4: Red brick
// - 5: Paddle
// - 6: Ball
//
// The function scales down the game state by a factor to fit into a fixed bitmap size.
// It processes the ball, paddle, and bricks, mapping their positions and dimensions
// to the bitmap grid.
//
// Parameters:
// - state: A pointer to a BreakoutState struct containing the current game state.
//
// Returns:
// - A 2D slice of integers representing the bitmap of the game state.
//   size of the slice is 80x60 (height x width) - compression factor around 3

func BreakoutState2Bitmap(state *breakout.BreakoutState) [][]int {
	colormap := map[string]int{
		"red":    4,
		"orange": 3,
		"green":  2,
		"yellow": 1,
	}
	bitH := 80
	bitW := 60
	factor := 3
	// Initialize the bitmap with empty spaces
	bitmap := make([][]int, bitH)
	for i := range bitmap {
		bitmap[i] = make([]int, bitW)
		for j := range bitmap[i] {
			bitmap[i][j] = 0 // Empty space
		}
	}
	// Draw the ball
	ballX := state.BallX / factor
	ballY := state.BallY / factor
	ballRadius := state.BallRadius / factor
	for i := -ballRadius; i <= ballRadius; i++ {
		for j := -ballRadius; j <= ballRadius; j++ {
			if i*i+j*j <= ballRadius*ballRadius {
				x := ballX + i
				y := ballY + j
				if x >= 0 && x < bitW && y >= 0 && y < bitH {
					bitmap[y][x] = 6 // Ball
				}
			}
		}
	}
	// Draw the paddle
	paddleX := state.PaddleX / factor
	paddleY := bitH - state.PaddleHeight/factor
	paddleWidth := state.PaddleWidth / factor
	for i := 0; i < paddleWidth; i++ {
		for j := 0; j < state.PaddleHeight/factor; j++ {
			x := paddleX + i
			y := paddleY + j
			if x >= 0 && x < bitW && y >= 0 && y < bitH {
				bitmap[y][x] = 5 // Paddle
			}
		}
	}

	// Draw the bricks
	for _, brick := range state.Bricks {
		brickX := float64(brick.X) / float64(factor)
		brickY := float64(brick.Y) / float64(factor)
		brickWidth := float64(brick.Width) / float64(factor)
		brickHeight := float64(brick.Height) / float64(factor)
		color := brick.Color
		colorValue := 0
		if val, ok := colormap[color]; ok {
			colorValue = val
		}
		for i := 0.0; i < brickWidth; i++ {
			for j := 0.0; j < brickHeight; j++ {
				x := int(math.Round(brickX + i))
				y := int(math.Round(brickY + j))
				if x >= 0 && x < bitW && y >= 0 && y < bitH {
					bitmap[y][x] = colorValue // Brick
				}
			}
		}
	}

	return bitmap
}
