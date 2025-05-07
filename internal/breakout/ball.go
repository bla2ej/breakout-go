// Package breakout provides the implementation of the Ball struct and its associated methods
// for simulating the behavior of a ball in a breakout game.
//
// The Ball struct represents a ball with properties such as position, radius, speed, and direction.
// It includes methods for creating a new ball, retrieving its properties, updating its direction,
// moving it within the game area, and reversing its velocity components.
package breakout

import (
	"fmt"
	"math"
	"math/rand"
)

type Ball struct {
	x, y     float64 // x and y coordinates of the ball
	radius   int     // radius of the ball
	dir      float64 // degree of speed vectore of the ball
	speed    float64 // speed of the ball
	v_x, v_y float64 // x and y speed of the ball
}

// NewBall creates a new Ball with the given x, y coordinates, radius, and direction
func NewBall() *Ball {
	b := &Ball{
		x:      float64(rand.Intn(AREA_WIDTH)-6) + 3,
		y:      AREA_HEIGHT / 2,
		radius: 2,
		speed:  3,
	}
	if rand.Intn(2) == 0 {
		b.SetDir(45)
	} else {
		b.SetDir(135)
	}
	return b
}

// GetX returns the x coordinate of the Ball
func (b *Ball) GetX() int {
	return int(b.x)
}

// GetY returns the y coordinate of the Ball
func (b *Ball) GetY() int {
	return int(b.y)
}

// GetRadius returns the radius of the Ball
func (b *Ball) GetRadius() int {
	return b.radius
}

// GetDir returns the direction of the Ball
func (b *Ball) GetDir() float64 {
	return b.dir
}

// SetDir sets the direction of the Ball
func (b *Ball) SetDir(dir float64) {
	b.dir = dir
	b.v_x = b.speed * math.Cos(dir*math.Pi/180)
	b.v_y = b.speed * math.Sin(dir*math.Pi/180)
	if b.v_x > 10 {
		b.v_x = 10
	}
	if b.v_x < -10 {
		b.v_x = -10
	}
	if b.v_y > BRICK_HEIGHT {
		b.v_y = BRICK_HEIGHT
	}
	if b.v_y < -BRICK_HEIGHT {
		b.v_y = -BRICK_HEIGHT
	}
}

// Move moves the ball in the direction of the dir value
// with the given speed by updating the x and y coordinates
func (b *Ball) Move() error {
	b.x += b.v_x
	b.y += b.v_y
	if int(b.x) <= b.radius {
		if b.v_x < 0 {
			b.x = float64(b.radius)
			b.v_x = -b.v_x
		}
	}
	if int(b.x) >= AREA_WIDTH-b.radius {
		if b.v_x > 0 {
			b.x = float64(AREA_WIDTH - b.radius)
			b.v_x = -b.v_x
		}
	}
	if int(b.y) <= b.radius {
		if b.v_y < 0 {
			b.y = float64(b.radius)
			b.v_y = -b.v_y
		}
	}
	if int(b.y) > AREA_HEIGHT-b.radius {
		if b.v_y > 0 {
			b.y = float64(AREA_HEIGHT - b.radius)
			b.v_y = -b.v_y
		}
		return fmt.Errorf("Ball is out of bounds")
	}
	return nil
}

// ReverseVX reverses the x direction of the Ball
func (b *Ball) ReverseVX() {
	b.v_x = -b.v_x
}

// ReverseVY reverses the y direction of the Ball
func (b *Ball) ReverseVY() {
	b.v_y = -b.v_y
}
