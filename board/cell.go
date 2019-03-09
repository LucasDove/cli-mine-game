package board

type CellState int32

const (
	Untoggled   CellState = 1
	Toggled     CellState = 2
)

type CellReactor interface {
	ActionAccepted() bool
	SetValue()
	GetValue()
	ChangeState(board BoardReactor)
}

type Cell struct {
	x int32
	y int32

	value int8
	state CellState
}

func (c *Cell) ActionAccepted() bool {

}

func (c *Cell) SetValue() {

}

func (c *Cell) GetValue() {

}

func (c *Cell) ChangeState(board BoardReactor) {

}



func NewCell(x, y int32) *Cell {
	return &Cell{
		x: x,
		y: y,
	}
}

func Right(cell *Cell) *Cell {
	return &Cell{
		x: cell.x + 1,
		y: cell.y,
	}
}

func Bottom(cell *Cell) *Cell {
	return &Cell{
		x: cell.x,
		y: cell.y + 1,
	}
}

func Left(cell *Cell) *Cell {
	return &Cell{
		x: cell.x - 1,
		y: cell.y,
	}
}

func Top(cell *Cell) *Cell {
	return &Cell{
		x: cell.x,
		y: cell.y - 1,
	}
}