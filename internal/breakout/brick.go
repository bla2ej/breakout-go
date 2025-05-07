// Package breakout provides the core functionality for the breakout game,
// including the representation and behavior of bricks in the game grid.
//
// The Brick struct represents a single brick in the game, with properties
// such as its position in the grid, dimensions, and state (e.g., whether it
// has been cleared). The BrickState struct is used to encapsulate the
// visual and positional state of a brick for rendering purposes.
//
// Constants such as AREA_WIDTH, BRICKS_PER_ROW, BRICK_HEIGHT, and TOP_OFFSET
// are assumed to define the layout and dimensions of the game area.
//
// Functions and methods provided:
// - NewBrick: Creates a new Brick instance with calculated dimensions and position.
// - GetRow, SetRow: Get or set the row index of the brick.
// - GetCol, SetCol: Get or set the column index of the brick.
// - IsCleared, SetCleared: Check or set whether the brick has been cleared.
// - GetX, GetY: Get the x and y coordinates of the brick.
// - CalcWidth: Calculate the width of the brick based on its column and layout.
// - GetWidth, GetHeight: Get the width and height of the brick.
// - GetColor: Determine the color of the brick based on its row.
// - GetPoints: Determine the points awarded for clearing the brick based on its row.
// - GetState: Retrieve the current state of the brick as a BrickState struct.
//
// The Brick struct uses the row index to determine the brick's color and
// point value, with higher rows corresponding to higher point values and
// different colors. The layout of bricks is calculated dynamically based
// on the game area's dimensions and the number of bricks per row.
package breakout

type Brick struct {
	row, col      int  // row and column of the brick
	cleared       bool // cleared is true if the brick has been hit
	x, y          int  // x and y coordinates of the brick
	width, height int  // width and height of the brick
}

type BrickState struct {
	X, Y          int    // coordinates of the brick
	Width, Height int    // dimensions of the brick
	Color         string // color of the brick
}

// NewBrick creates a new Brick with the given row and column
func NewBrick(row, col int) *Brick {
	br := &Brick{row: row, col: col}
	br.width = br.CalcWidth()
	br.height = BRICK_HEIGHT
	br.y = TOP_OFFSET + (BRICK_ROWS-row-1)*BRICK_HEIGHT
	br.x = col * (AREA_WIDTH / BRICKS_PER_ROW)
	if col > 0 {
		realWidth := float64(AREA_WIDTH) / float64(BRICKS_PER_ROW)
		intWidth := int(realWidth)
		rest := AREA_WIDTH - intWidth*BRICKS_PER_ROW
		br.x += rest / 2
	}
	return br
}

// GetRow returns the row of the Brick
func (b *Brick) GetRow() int {
	return b.row
}

// SetRow sets the row of the Brick
func (b *Brick) SetRow(row int) {
	b.row = row
}

// GetCol returns the column of the Brick
func (b *Brick) GetCol() int {
	return b.col
}

// SetCol sets the column of the Brick
func (b *Brick) SetCol(col int) {
	b.col = col
}

// IsCleared returns true if the Brick is cleared
func (b *Brick) IsCleared() bool {
	return b.cleared
}

// SetCleared sets the cleared status of the Brick
func (b *Brick) SetCleared(cleared bool) {
	b.cleared = cleared
}

// GetX returns the x coordinate of the Brick
func (b *Brick) GetX() int {
	return b.x
}

// GetY returns the y coordinate of the Brick
func (b *Brick) GetY() int {
	return b.y
}

// GetWidth returns the width of the Brick
func (b *Brick) CalcWidth() int {
	realWidth := float64(AREA_WIDTH) / float64(BRICKS_PER_ROW)
	intWidth := int(realWidth)
	rest := AREA_WIDTH - intWidth*BRICKS_PER_ROW
	if b.col == 0 {
		return intWidth + rest/2
	}
	if b.col == BRICKS_PER_ROW-1 {
		return AREA_WIDTH - (BRICKS_PER_ROW-1)*intWidth - rest/2
	}
	return intWidth
}

// GetWidth returns the width of the Brick
func (b *Brick) GetWidth() int {
	return b.width
}

// GetHeight returns the height of the Brick
func (b *Brick) GetHeight() int {
	return b.height
}

// GetColor returns the color of the Brick
func (b *Brick) GetColor() string {
	switch b.row {
	case 0, 1:
		return "yellow"
	case 2, 3:
		return "green"
	case 4, 5:
		return "orange"
	case 6, 7:
		return "red"
	default: // should never happen
		return "red"
	}
}

// GetPoints returns the points of the Brick
func (b *Brick) GetPoints() int {
	switch b.row {
	case 0, 1:
		return 1
	case 2, 3:
		return 3
	case 4, 5:
		return 5
	case 6, 7:
		return 7
	default: // should never happen
		return 7
	}
}

// GetState of the Brick
func (b *Brick) GetState() BrickState {
	return BrickState{
		X:      b.x,
		Y:      b.y,
		Width:  b.width,
		Height: b.height,
		Color:  b.GetColor(),
	}
}
