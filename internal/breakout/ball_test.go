package breakout

import (
	"math"
	"testing"
)

func TestBallNew(t *testing.T) {
	ball := NewBall()
	ball.SetDir(45)

	if ball.GetY() != AREA_HEIGHT/2 {
		t.Errorf("Expected Y to be %d, got %d", AREA_HEIGHT/2, ball.GetY())
	}

	if ball.GetRadius() != 2 {
		t.Errorf("Expected radius to be 2, got %d", ball.GetRadius())
	}

	if ball.GetDir() != 45 {
		t.Errorf("Expected direction to be 45, got %f", ball.GetDir())
	}
}

func TestBallSetDir(t *testing.T) {
	ball := NewBall()
	ball.SetDir(90)

	if ball.GetDir() != 90 {
		t.Errorf("Expected direction to be 90, got %f", ball.GetDir())
	}

	expectedVx := ball.speed * math.Cos(90*math.Pi/180)
	expectedVy := ball.speed * math.Sin(90*math.Pi/180)

	if ball.v_x != expectedVx {
		t.Errorf("Expected v_x to be %f, got %f", expectedVx, ball.v_x)
	}

	if ball.v_y != expectedVy {
		t.Errorf("Expected v_y to be %f, got %f", expectedVy, ball.v_y)
	}
}

func TestBallMove(t *testing.T) {
	ball := NewBall()
	ball.x = AREA_WIDTH / 2
	ball.SetDir(45) // Set direction to 45 degrees

	for i := 0; i < 1000; i++ {
		ball.Move()
		// fmt.Printf("Ball position: (%d, %d)\n", ball.GetX(), ball.GetY())
	}

	if ball.GetX() != 74 || ball.GetY() != 129 {
		t.Errorf("Ball did not move as expected")
	}
}

func TestBallReverseVX(t *testing.T) {
	ball := NewBall()
	initialVX := ball.v_x
	ball.ReverseVX()

	if ball.v_x != -initialVX {
		t.Errorf("Expected v_x to be %f, got %f", -initialVX, ball.v_x)
	}
}

func TestBallReverseVY(t *testing.T) {
	ball := NewBall()
	initialVY := ball.v_y
	ball.ReverseVY()

	if ball.v_y != -initialVY {
		t.Errorf("Expected v_y to be %f, got %f", -initialVY, ball.v_y)
	}
}

func TestBallCollisionWithWalls(t *testing.T) {
	ball := NewBall()

	// Test collision with left wall
	ball.x = float64(ball.radius)
	ball.v_x = -1
	ball.Move()
	if ball.v_x < 0 {
		t.Errorf("Expected v_x to be reversed after hitting left wall, got %f", ball.v_x)
	}

	// Test collision with right wall
	ball.x = float64(AREA_WIDTH - ball.radius)
	ball.v_x = 1
	ball.Move()
	if ball.v_x > 0 {
		t.Errorf("Expected v_x to be reversed after hitting right wall, got %f", ball.v_x)
	}

	// Test collision with top wall
	ball.y = float64(ball.radius)
	ball.v_y = -1
	ball.Move()
	if ball.v_y < 0 {
		t.Errorf("Expected v_y to be reversed after hitting top wall, got %f", ball.v_y)
	}
}

func TestBallSetDirLimits(t *testing.T) {
	ball := NewBall()
	ball.SetDir(45)

	if ball.v_x > 10 {
		t.Errorf("Expected v_x to be <= 10, got %f", ball.v_x)
	}

	if ball.v_y > BRICK_HEIGHT {
		t.Errorf("Expected v_y to be <= BRICK_HEIGHT, got %f", ball.v_y)
	}
}

func TestBallTooHighSpeed(t *testing.T) {
	ball := NewBall()
	ball.speed = 20
	ball.SetDir(0)

	if ball.v_x > 10 {
		t.Errorf("Expected v_x to be <= 10, got %f", ball.v_x)
	}

	ball.SetDir(90)
	if ball.v_y > BRICK_HEIGHT {
		t.Errorf("Expected v_y to be <= BRICK_HEIGHT, got %f", ball.v_y)
	}

	ball.SetDir(180)
	if ball.v_x < -10 {
		t.Errorf("Expected v_x to be >= -10, got %f", ball.v_x)
	}

	ball.SetDir(270)
	if ball.v_y < -BRICK_HEIGHT {
		t.Errorf("Expected v_y to be >= -BRICK_HEIGHT, got %f", ball.v_y)
	}

}
