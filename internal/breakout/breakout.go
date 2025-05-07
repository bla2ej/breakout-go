// Package breakout implements the core logic for a breakout-style game.
// It includes the game state, mechanics, and interactions between game elements
// such as the ball, paddle, and bricks.

//
// Constants:
// - AREA_WIDTH: The width of the game area.
// - AREA_HEIGHT: The height of the game area.
// - BRICKS_PER_ROW: The number of bricks in each row.
// - BRICK_ROWS: The number of rows of bricks.
// - TOP_OFFSET: The vertical offset from the top of the game area.
// - BRICK_HEIGHT: The height of each brick.
//
// Variables:
// - ErrGameOver: An error indicating that the game is over.
//
// Types:
// - Breakout: Represents the main game state and logic.
// - BreakoutState: Represents the current state of the game, used for rendering or external interaction.
//
// Functions:
// - NewBreakout: Creates and initializes a new Breakout game instance.
// - (*Breakout) GetState: Returns the current state of the game as a BreakoutState.
// - (*Breakout) MoveBall: Updates the ball's position and handles collisions with bricks, the paddle, and the game area.
// - (*Breakout) PaddleRight: Moves the paddle to the right.
// - (*Breakout) PaddleLeft: Moves the paddle to the left.
// - (*Breakout) PaddleShrink: Shrinks the paddle's width.
// - (*Breakout) PaddleUnShrink: Restores the paddle's width to its original size.
// - (*Breakout) CheckColision: Checks for collisions between the ball and a brick, returning whether the ball should reverse its x or y velocity.
// - (*Breakout) CheckPaddleColision: Checks for collisions between the ball and the paddle, adjusting the ball's direction if necessary.
package breakout

import (
	"errors"
)

const (
	AREA_WIDTH     = 182
	AREA_HEIGHT    = 240
	BRICKS_PER_ROW = 14
	BRICK_ROWS     = 8
	TOP_OFFSET     = 32
	BRICK_HEIGHT   = 7
)

var ErrGameOver = errors.New("game over")

type Breakout struct {
	ball        *Ball
	bricks      [][]*Brick
	score       int
	level       int
	live        int
	paddle      *Paddle // Paddle is a struct that represents the paddle in the game
	frameReward int     // reward for the current frame

	// Game state
	gameOver bool
}

type BreakoutState struct {
	BallX, BallY  int          // current ball coordinates
	BallRadius    int          // ball radius
	Width, Height int          // screen dimensions
	PaddleX       int          // paddle x coordinate
	PaddleWidth   int          // paddle width
	PaddleHeight  int          // paddle height
	Bricks        []BrickState // bricks state
	Level         int          // current level
	Score         int          // current score
	Live          int          // current live
	FrameReward   int          // reward for the current frame
	Done          bool         // game over
}

func NewBreakout() *Breakout {
	bricks := make([][]*Brick, BRICK_ROWS)
	for i := 0; i < BRICK_ROWS; i++ {
		bricks[i] = make([]*Brick, BRICKS_PER_ROW)
		for j := 0; j < BRICKS_PER_ROW; j++ {
			bricks[i][j] = NewBrick(i, j)
		}
	}
	// Initialize the paddle
	paddle := NewPaddle()

	return &Breakout{
		ball:     NewBall(),
		bricks:   bricks,
		paddle:   paddle,
		score:    0,
		level:    1,
		live:     1,
		gameOver: false,
	}
}

func (b *Breakout) GetState() BreakoutState {
	state := BreakoutState{
		BallX:        b.ball.GetX(),
		BallY:        b.ball.GetY(),
		BallRadius:   b.ball.GetRadius(),
		Width:        AREA_WIDTH,
		Height:       AREA_HEIGHT,
		PaddleX:      b.paddle.GetX(),
		PaddleWidth:  b.paddle.GetWidth(),
		PaddleHeight: b.paddle.GetHeight(),
		Level:        b.level,
		Score:        b.score,
		Live:         b.live,
		Done:         b.gameOver,
		FrameReward:  b.frameReward,
	}
	for i := range BRICK_ROWS {
		for j := range BRICKS_PER_ROW {
			if b.bricks[i][j] != nil {
				if !b.bricks[i][j].IsCleared() {
					state.Bricks = append(state.Bricks, b.bricks[i][j].GetState())
				}
			}
		}
	}
	return state
}

func (b *Breakout) MoveBall() {
	if b.gameOver {
		return
	}
	err := b.ball.Move()
	if err != nil {
		b.live++
		// b.score--
		if b.live > 5 {
			b.gameOver = true
		}
		b.frameReward = -10
		b.ball = NewBall()
		b.paddle = NewPaddle()
		return
	}
	// check ball collisions with bricks
	cleared := 0
	xrev := false
	yrev := false
	for i := range b.bricks {
		for j := range b.bricks[i] {
			if b.bricks[i][j] != nil {
				if !b.bricks[i][j].IsCleared() {
					if !xrev && !yrev {
						xr, yr := b.CheckColision(b.bricks[i][j], b.ball)
						if xr || yr {
							b.bricks[i][j].SetCleared(true)
							b.score += b.bricks[i][j].GetPoints() * b.level
							cleared++
						}
						xrev = xrev || xr
						yrev = yrev || yr
					}
				} else {
					cleared++
				}
			}
		}
	}
	if xrev {
		b.ball.ReverseVX()
	}
	if yrev {
		b.ball.ReverseVY()
	}
	// check ball collisions with paddle
	if b.CheckPaddleColision(b.paddle, b.ball) {
		b.frameReward = 10
	} else {
		b.frameReward = 0
	}

	// check if there is no more bricks left
	// all bricks are cleared
	if cleared == BRICKS_PER_ROW*BRICK_ROWS {
		b.level++
		b.bricks = make([][]*Brick, BRICK_ROWS)
		for i := range BRICK_ROWS {
			b.bricks[i] = make([]*Brick, BRICKS_PER_ROW)
			for j := range BRICKS_PER_ROW {
				b.bricks[i][j] = NewBrick(i, j)
			}
		}
		b.ball = NewBall()
		b.paddle = NewPaddle()
	}

}

func (b *Breakout) PaddleRight() {
	b.paddle.MoveRight()
}

func (b *Breakout) PaddleLeft() {
	b.paddle.MoveLeft()
}
func (b *Breakout) PaddleShrink() {
	b.paddle.Shrink()
}
func (b *Breakout) PaddleUnShrink() {
	b.paddle.UnShrink()
}

func (b *Breakout) CheckColision(br *Brick, bl *Ball) (bool, bool) {
	// check if ball should bounce off the brick
	// consider ball radius
	if br.IsCleared() {
		return false, false
	}
	blX := bl.GetX()
	blY := bl.GetY()
	blR := bl.GetRadius()
	brX := br.GetX()
	brY := br.GetY()
	brW := br.GetWidth()
	brH := br.GetHeight()
	xrev := false
	yrev := false
	if blX+blR > brX && blX-blR <= brX { //might be hiting the left side
		if blY+blR > brY && blY+blR <= brY+brH { // right hight so colision is there and ball bounces
			if b.ball.v_x > 0 {
				xrev = true
			}
		}
		if blY-blR > brY && blY-blR <= brY+brH {
			if b.ball.v_x > 0 {
				xrev = true
			}
		}
	}
	if blX-blR < brX+brW && blX+blR >= brX+brW { //might be hiting the right side
		if blY+blR > brY && blY+blR <= brY+brH { // right hight so colision is there and ball bounces
			if b.ball.v_x < 0 {
				xrev = true
			}
		}
		if blY-blR > brY && blY-blR <= brY+brH {
			if b.ball.v_x < 0 {
				xrev = true
			}
		}
	}
	if blY+blR > brY && blY-blR <= brY { //might be hiting the top side
		if blX+blR > brX && blX+blR <= brX+brW { // right x position so colision is there and ball bounces
			if b.ball.v_y > 0 {
				yrev = true
			}
		}
		if blX-blR > brX && blX-blR <= brX+brW {
			if b.ball.v_y > 0 {
				yrev = true
			}
		}
	}
	if blY-blR < brY+brH && blY+blR >= brY+brH { //might be hiting the bottom side
		if blX+blR > brX && blX+blR <= brX+brW { // right x position so colision is there and ball bounces
			if b.ball.v_y < 0 {
				yrev = true
			}
		}
		if blX-blR > brX && blX-blR <= brX+brW {
			if b.ball.v_y < 0 {
				yrev = true
			}
		}
	}
	return xrev, yrev
}

func (b *Breakout) CheckPaddleColision(pa *Paddle, bl *Ball) bool {
	// check if ball should bounce off the paddle
	// consider ball radius
	blX := bl.GetX()
	blY := bl.GetY()
	blR := bl.GetRadius()
	paX := pa.GetX()
	paY := AREA_HEIGHT - pa.GetHeight()
	paW := pa.GetWidth()
	paH := pa.GetHeight()
	col := false
	if blY+blR >= paY && blY < paY+paH { //might be hiting the top side
		if blX+blR >= paX && blX+blR <= paX+paW { // right x position so colision is there and ball bounces
			col = true
		}
		if blX-blR >= paX && blX-blR <= paX+paW {
			col = true
		}
	}
	if col {
		xpaddle := paX + paW/2
		xball := blX
		h := float64(xball-xpaddle) / float64(paW/2)
		bl.SetDir(270 + h*60 + 1) // plus one to avoid 0 degree
	}
	return col
}
