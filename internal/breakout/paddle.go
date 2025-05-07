// Package breakout provides the core components and functionality for the
// Breakout game, including the Paddle struct and its associated methods.
//
// The Paddle struct represents the player's paddle in the game, which can
// move horizontally and change its size. The package includes methods for
// creating a new paddle, retrieving its properties, and modifying its
// position and dimensions.
//
// Constants:
// - AREA_WIDTH: Represents the width of the game area. This constant is
//   used to constrain the paddle's movement within the game boundaries.
//
// Types:
// - Paddle: A struct that defines the paddle's position (x-coordinate),
//   width, and height.
//
// Functions:
// - NewPaddle: Creates and returns a new Paddle instance with default
//   dimensions and position.
//
// Methods:
// - GetX: Returns the current x-coordinate of the paddle.
// - GetWidth: Returns the current width of the paddle.
// - GetHeight: Returns the current height of the paddle.
// - Shrink: Reduces the paddle's width to a smaller size (minimum 12).
// - UnShrink: Ensures the paddle's width is at least the default minimum
//   width of 24.
// - MoveLeft: Moves the paddle to the left by 3 units, constrained by the
//   left boundary of the game area.
// - MoveRight: Moves the paddle to the right by 3 units, constrained by
//   the right boundary of the game area.
package breakout

type Paddle struct {
	x      int
	width  int
	height int
}

// NewPaddle creates and returns a new Paddle instance with default
// dimensions and position. The paddle is initialized at the center
// of the game area horizontally, with a predefined width and height.
func NewPaddle() *Paddle {
	return &Paddle{
		x:      AREA_WIDTH/2 - 12,
		width:  24,
		height: 4,
	}
}

// GetX returns the current x-coordinate of the paddle.
// This represents the horizontal position of the paddle in the game.
func (p *Paddle) GetX() int {
	return p.x
}

// GetWidth returns the current width of the paddle.
// This represents the horizontal size of the paddle in the game.
func (p *Paddle) GetWidth() int {
	return p.width
}

// GetHeight returns the current height of the paddle.
// This represents the vertical size of the paddle in the game.
func (p *Paddle) GetHeight() int {
	return p.height
}

// Shrink reduces the width of the paddle to a smaller size.
// The width is set to 12 if it is currently greater than 12.
func (p *Paddle) Shrink() {
	if p.width > 12 {
		p.width = 12
	}
}

// UnShrink ensures that the paddle's width is at least the default minimum width of 24.
// If the current width is less than 24, it resets the width to 24.
func (p *Paddle) UnShrink() {
	if p.width < 24 {
		p.width = 24
	}
}

// MoveLeft moves the paddle to the left by 2 units.
// The paddle's position is constrained to ensure it does not move
// beyond the left boundary of the game area.
func (p *Paddle) MoveLeft() {
	if p.x >= 0 {
		p.x -= 3
	}
	if p.x < 0 {
		p.x = 0
	}
}

// MoveRight moves the paddle to the right by 2 units.
// The paddle's position is constrained to ensure it does not move
// beyond the right boundary of the game area.
func (p *Paddle) MoveRight() {
	if p.x <= AREA_WIDTH-p.width {
		p.x += 3
	}
	if p.x > AREA_WIDTH-p.width {
		p.x = AREA_WIDTH - p.width
	}
}
