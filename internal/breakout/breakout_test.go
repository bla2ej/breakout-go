package breakout

import (
	"testing"
)

func TestNewBreakout(t *testing.T) {
	breakout := NewBreakout()

	if breakout == nil {
		t.Fatal("Expected Breakout instance, got nil")
	}

	if breakout.ball == nil {
		t.Error("Expected ball to be initialized")
	}

	if breakout.paddle == nil {
		t.Error("Expected paddle to be initialized")
	}

	if len(breakout.bricks) != BRICK_ROWS {
		t.Errorf("Expected %d rows of bricks, got %d", BRICK_ROWS, len(breakout.bricks))
	}

	for i := range breakout.bricks {
		if len(breakout.bricks[i]) != BRICKS_PER_ROW {
			t.Errorf("Expected %d bricks in row %d, got %d", BRICKS_PER_ROW, i, len(breakout.bricks[i]))
		}
	}
}

func TestMoveBall_GameOver(t *testing.T) {
	breakout := NewBreakout()
	breakout.gameOver = true

	breakout.MoveBall()

	if breakout.ball == nil {
		t.Error("Expected ball to remain unchanged when game is over")
	}
}

func TestMoveBall_LifeLost(t *testing.T) {
	breakout := NewBreakout()
	breakout.ball.y = AREA_HEIGHT + 1 // Simulate ball falling out of bounds

	breakout.MoveBall()

	if breakout.live != 2 {
		t.Errorf("Expected live to increase to 2, got %d", breakout.live)
	}

	if breakout.frameReward != -10 {
		t.Errorf("Expected frameReward to be -10, got %d", breakout.frameReward)
	}
}

func TestPaddleMovement(t *testing.T) {
	breakout := NewBreakout()
	initialX := breakout.paddle.GetX()

	breakout.PaddleRight()
	if breakout.paddle.GetX() <= initialX {
		t.Error("Expected paddle to move right")
	}

	breakout.PaddleLeft()
	breakout.PaddleLeft()
	if breakout.paddle.GetX() >= initialX {
		t.Error("Expected paddle to move left")
	}
}

func TestPaddleShrinkUnshrink(t *testing.T) {
	breakout := NewBreakout()
	initialWidth := breakout.paddle.GetWidth()

	breakout.PaddleShrink()
	if breakout.paddle.GetWidth() >= initialWidth {
		t.Error("Expected paddle to shrink")
	}

	breakout.PaddleUnShrink()
	if breakout.paddle.GetWidth() != initialWidth {
		t.Error("Expected paddle to return to original width")
	}
}

func TestCheckPaddleCollision(t *testing.T) {
	breakout := NewBreakout()
	ball := NewBall()
	paddle := breakout.paddle

	ball.x = float64(paddle.x + paddle.width/2)
	ball.y = float64(AREA_HEIGHT - paddle.height - 1)
	collision := breakout.CheckPaddleColision(paddle, ball)

	if !collision {
		t.Error("Expected collision with paddle")
	}
}

func TestGetState(t *testing.T) {
	breakout := NewBreakout()
	state := breakout.GetState()

	if state.Width != AREA_WIDTH || state.Height != AREA_HEIGHT {
		t.Error("Expected state dimensions to match game area")
	}

	if state.Level != breakout.level {
		t.Errorf("Expected level %d, got %d", breakout.level, state.Level)
	}

	if state.Score != breakout.score {
		t.Errorf("Expected score %d, got %d", breakout.score, state.Score)
	}

	if state.Live != breakout.live {
		t.Errorf("Expected live %d, got %d", breakout.live, state.Live)
	}
}
func TestCheckCollision_BrickHit_FromRight(t *testing.T) {
	breakout := NewBreakout()
	ball := NewBall()
	breakout.ball = ball
	brick := NewBrick(1, 1)

	// Simulate ball hitting the brick from right
	ball.SetDir(180)
	ball.x = float64(brick.x + brick.width + 1)
	ball.y = float64(brick.y + brick.height/2)

	xrev, yrev := breakout.CheckColision(brick, ball)

	if !xrev || yrev {
		t.Error("Expected collision with brick")
	}

	if brick.IsCleared() {
		t.Error("Expected brick to be cleared after collision")
	}
}

func TestCheckCollision_BrickHit_FromLeft(t *testing.T) {
	breakout := NewBreakout()
	ball := NewBall()
	breakout.ball = ball
	brick := NewBrick(1, 1)

	// Simulate ball hitting the brick from left
	ball.SetDir(0)
	ball.x = float64(brick.x - 1)
	ball.y = float64(brick.y + brick.height/2)

	xrev, yrev := breakout.CheckColision(brick, ball)

	if !xrev || yrev {
		t.Error("Expected collision with brick")
	}

	if brick.IsCleared() {
		t.Error("Expected brick to be cleared after collision")
	}
}

func TestCheckCollision_BrickHit_FromTop(t *testing.T) {
	breakout := NewBreakout()
	ball := NewBall()
	breakout.ball = ball
	brick := NewBrick(1, 1)

	// Simulate ball hitting the brick from top
	ball.SetDir(90)
	ball.x = float64(brick.x + brick.width/2)
	ball.y = float64(brick.y - 1)

	xrev, yrev := breakout.CheckColision(brick, ball)

	if xrev || !yrev {
		t.Error("Expected collision with brick")
	}

	if brick.IsCleared() {
		t.Error("Expected brick to be cleared after collision")
	}
}

func TestCheckCollision_BrickHit_FromBottom(t *testing.T) {
	breakout := NewBreakout()
	ball := NewBall()
	breakout.ball = ball
	brick := NewBrick(1, 1)

	// Simulate ball hitting the brick from bottom
	ball.SetDir(270)
	ball.x = float64(brick.x + brick.width/2)
	ball.y = float64(brick.y + brick.height + 1)

	xrev, yrev := breakout.CheckColision(brick, ball)

	if xrev || !yrev {
		t.Error("Expected collision with brick")
	}

	if brick.IsCleared() {
		t.Error("Expected brick to be cleared after collision")
	}
}

func TestCheckCollision_BrickHit_AlreadyCleared(t *testing.T) {
	breakout := NewBreakout()
	ball := NewBall()
	breakout.ball = ball
	brick := NewBrick(1, 1)
	brick.SetCleared(true)

	// Simulate ball hitting the already cleared brick
	ball.SetDir(0)
	ball.x = float64(brick.x - 1)
	ball.y = float64(brick.y + brick.height/2)

	xrev, yrev := breakout.CheckColision(brick, ball)

	if xrev || yrev {
		t.Error("Expected no collision with cleared brick")
	}

	if !brick.IsCleared() {
		t.Error("Expected brick to remain cleared")
	}
}

func TestMoveBallAndClearBrick(t *testing.T) {
	breakout := NewBreakout()
	breakout.ball.x = float64(0)
	breakout.ball.SetDir(45)
	breakout.paddle.x = 110
	for i := 0; i < 150; i++ {
		breakout.MoveBall()
	}
	clearedBricks := 0
	for _, row := range breakout.bricks {
		for _, brick := range row {
			if brick.IsCleared() {
				clearedBricks++
			}
		}
	}

	if clearedBricks != 1 {
		t.Error("Expected one brick to be cleared")
	}
}
