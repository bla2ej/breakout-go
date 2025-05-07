package breakout

import "testing"

func TestPaddleNew(t *testing.T) {
	paddle := NewPaddle()
	if paddle.x != AREA_WIDTH/2-12 {
		t.Errorf("Expected x to be %d, got %d", AREA_WIDTH/2-12, paddle.x)
	}
	if paddle.width != 24 {
		t.Errorf("Expected width to be 24, got %d", paddle.width)
	}
	if paddle.height != 4 {
		t.Errorf("Expected height to be 4, got %d", paddle.height)
	}
}

func TestPaddleGetX(t *testing.T) {
	paddle := NewPaddle()
	if paddle.GetX() != paddle.x {
		t.Errorf("Expected GetX to return %d, got %d", paddle.x, paddle.GetX())
	}
}

func TestPaddleGetWidth(t *testing.T) {
	paddle := NewPaddle()
	if paddle.GetWidth() != paddle.width {
		t.Errorf("Expected GetWidth to return %d, got %d", paddle.width, paddle.GetWidth())
	}
}

func TestPaddleGetHeight(t *testing.T) {
	paddle := NewPaddle()
	if paddle.GetHeight() != paddle.height {
		t.Errorf("Expected GetHeight to return %d, got %d", paddle.height, paddle.GetHeight())
	}
}

func TestPaddleShrink(t *testing.T) {
	paddle := NewPaddle()
	paddle.Shrink()
	if paddle.width != 12 {
		t.Errorf("Expected width to be 12 after Shrink, got %d", paddle.width)
	}
}

func TestPaddleUnShrink(t *testing.T) {
	paddle := NewPaddle()
	paddle.Shrink()
	if paddle.width != 12 {
		t.Errorf("Expected width to be 12 after Shrink, got %d", paddle.width)
	}
	paddle.UnShrink()
	if paddle.width != 24 {
		t.Errorf("Expected width to be 24 after UnShrink, got %d", paddle.width)
	}
}

func TestPaddleMoveLeft(t *testing.T) {
	paddle := NewPaddle()
	initialX := paddle.x
	paddle.MoveLeft()
	if paddle.x != initialX-3 && paddle.x != 0 {
		t.Errorf("Expected x to be %d after MoveLeft, got %d", initialX-2, paddle.x)
	}
	for i := 0; i < 1000; i++ {
		paddle.MoveLeft()
	}	
	if paddle.x != 0 {
		t.Errorf("Expected x to be 0 after moving left multiple times, got %d", paddle.x)
	}

}

func TestPaddleMoveRight(t *testing.T) {
	paddle := NewPaddle()
	initialX := paddle.x
	paddle.MoveRight()
	if paddle.x != initialX+3 && paddle.x != AREA_WIDTH-paddle.width {
		t.Errorf("Expected x to be %d after MoveRight, got %d", initialX+2, paddle.x)
	}
	for i := 0; i < 1000; i++ {
		paddle.MoveRight()
	}
	if paddle.x != AREA_WIDTH-paddle.width {
		t.Errorf("Expected x to be %d after moving right multiple times, got %d", AREA_WIDTH-paddle.width, paddle.x)
	}
}

