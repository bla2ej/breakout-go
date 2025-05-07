package breakout

import "testing"

func TestNewBrick(t *testing.T) {
	brick := NewBrick(2, 3)
	if brick.GetRow() != 2 {
		t.Errorf("expected row to be 2, got %d", brick.GetRow())
	}
	if brick.GetCol() != 3 {
		t.Errorf("expected col to be 3, got %d", brick.GetCol())
	}
}

func TestBrickSetRow(t *testing.T) {
	brick := NewBrick(0, 0)
	brick.SetRow(5)
	if brick.GetRow() != 5 {
		t.Errorf("expected row to be 5, got %d", brick.GetRow())
	}
}

func TestBrickSetCol(t *testing.T) {
	brick := NewBrick(0, 0)
	brick.SetCol(7)
	if brick.GetCol() != 7 {
		t.Errorf("expected col to be 7, got %d", brick.GetCol())
	}
}

func TestBrickGetWidth(t *testing.T) {
	firstBrick := NewBrick(0, 0)
	lastBrick := NewBrick(0, BRICKS_PER_ROW-1)
	middleBrick := NewBrick(0, 1)

	firstWidth := firstBrick.GetWidth()
	lastWidth := lastBrick.GetWidth()
	middleWidth := middleBrick.GetWidth()

	expectedFirstWidth := AREA_WIDTH/BRICKS_PER_ROW 
	expectedLastWidth := AREA_WIDTH/BRICKS_PER_ROW 
	expectedMiddleWidth := AREA_WIDTH / BRICKS_PER_ROW

	if firstWidth != expectedFirstWidth {
		t.Errorf("expected first brick width to be %d, got %d", expectedFirstWidth, firstWidth)
	}
	if lastWidth != expectedLastWidth {
		t.Errorf("expected last brick width to be %d, got %d", expectedLastWidth, lastWidth)
	}
	if middleWidth != expectedMiddleWidth {
		t.Errorf("expected middle brick width to be %d, got %d", expectedMiddleWidth, middleWidth)
	}
}

func TestBrickGetHeight(t *testing.T) {
	brick := NewBrick(0, 0)
	height := brick.GetHeight()
	expectedHeight := 7

	if height != expectedHeight {
		t.Errorf("expected height to be %d, got %d", expectedHeight, height)
	}
}
func TestBrickGetColor(t *testing.T) {
	brick := NewBrick(0, 0)
	color := brick.GetColor()
	expectedColor := "yellow"

	if color != expectedColor {
		t.Errorf("expected color to be %s, got %s", expectedColor, color)
	}

	brick.SetRow(2)
	color = brick.GetColor()
	expectedColor = "green"

	if color != expectedColor {
		t.Errorf("expected color to be %s, got %s", expectedColor, color)
	}
	brick.SetRow(4)
	color = brick.GetColor()
	expectedColor = "orange"

	if color != expectedColor {
		t.Errorf("expected color to be %s, got %s", expectedColor, color)
	}

	brick.SetRow(6)
	color = brick.GetColor()
	expectedColor = "red"

	if color != expectedColor {
		t.Errorf("expected color to be %s, got %s", expectedColor, color)
	}

	brick.SetRow(8)
	color = brick.GetColor()
	expectedColor = "red"

	if color != expectedColor {
		t.Errorf("expected color to be %s, got %s", expectedColor, color)
	}
}

func TestBrickGetPoints(t *testing.T) {
	brick := NewBrick(0, 1)
	points := brick.GetPoints()
	expectedPoints := 1

	if points != expectedPoints {
		t.Errorf("expected points to be %d, got %d", expectedPoints, points)
	}

	brick.SetRow(3)
	points = brick.GetPoints()
	expectedPoints = 3

	if points != expectedPoints {
		t.Errorf("expected points to be %d, got %d", expectedPoints, points)
	}
	brick.SetRow(5)
	points = brick.GetPoints()
	expectedPoints = 5

	if points != expectedPoints {
		t.Errorf("expected points to be %d, got %d", expectedPoints, points)
	}

	brick.SetRow(7)
	points = brick.GetPoints()
	expectedPoints = 7

	if points != expectedPoints {
		t.Errorf("expected points to be %d, got %d", expectedPoints, points)
	}
}
func TestBrickGetPointsWithInvalidRow(t *testing.T) {
	brick := NewBrick(0, 1)
	brick.SetRow(8) // Invalid row
	points := brick.GetPoints()
	expectedPoints := 7 // Default to last row points

	if points != expectedPoints {
		t.Errorf("expected points to be %d, got %d", expectedPoints, points)
	}
}
func TestBrickIsCleared(t *testing.T) {
	brick := NewBrick(0, 0)
	if brick.IsCleared() {
		t.Errorf("expected brick to not be cleared, but it is")
	}

	brick.SetCleared(true)
	if !brick.IsCleared() {
		t.Errorf("expected brick to be cleared, but it is not")
	}
}

func TestBrickGetState(t *testing.T) {
	brick := NewBrick(2, 3)
	state := brick.GetState()

	if state.X != brick.GetX() {
		t.Errorf("expected state X to be %d, got %d", brick.GetX(), state.X)
	}
	if state.Y != brick.GetY() {
		t.Errorf("expected state Y to be %d, got %d", brick.GetY(), state.Y)
	}
	if state.Width != brick.GetWidth() {
		t.Errorf("expected state Width to be %d, got %d", brick.GetWidth(), state.Width)
	}
	if state.Height != brick.GetHeight() {
		t.Errorf("expected state Height to be %d, got %d", brick.GetHeight(), state.Height)
	}
	if state.Color != brick.GetColor() {
		t.Errorf("expected state Color to be %s, got %s", brick.GetColor(), state.Color)
	}
}

func TestBrickSetCleared(t *testing.T) {
	brick := NewBrick(0, 0)
	brick.SetCleared(true)
	if !brick.IsCleared() {
		t.Errorf("expected brick to be cleared, but it is not")
	}

	brick.SetCleared(false)
	if brick.IsCleared() {
		t.Errorf("expected brick to not be cleared, but it is")
	}
}
