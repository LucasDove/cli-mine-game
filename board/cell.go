package board

const (
	Mine = int8(100)
)

type CellReactor interface {
	Toggle(b BoardReactor) bool
	ToggleSpaces(b BoardReactor)

	GetValue() int8
	SetValue(value int8)
	IsToggled() bool
}

type Cell struct {
	x int32
	y int32

	value     int8
	isToggled bool
}

func (c *Cell) SetValue(value int8) {
	c.value = value
}

func (c *Cell) GetValue() int8 {
	return c.value
}

func (c *Cell) IsToggled() bool {
	return c.isToggled
}

func (c *Cell) Toggle(b BoardReactor) bool {
	if c.isToggled {
		return true
	}
	c.isToggled = true
	if c.value != Mine {
		return false
	}else if c.value > 0 {
		return true
	}else {
		c.ToggleSpaces(b)
	}

	return true
}

func (c *Cell) ToggleSpaces(b BoardReactor) {
	if c.isToggled {
		return
	}
	if c.value != Mine && c.value > 0 {
		//whitespaces
		c.isToggled = true
		cells := c.aroundCells(b)
		for _, cell := range cells {
			cell.ToggleSpaces(b)
		}
	}
}

func (c *Cell) aroundCells(b BoardReactor) []CellReactor {
	var arrounds []CellReactor
	if b.IsValidCell(c.x+1, c.y) {
		arrounds = append(arrounds, b.GetCell(c.x+1, c.y))
	}
	if b.IsValidCell(c.x-1, c.y) {
		arrounds = append(arrounds, b.GetCell(c.x-1, c.y))
	}
	if b.IsValidCell(c.x, c.y+1) {
		arrounds = append(arrounds, b.GetCell(c.x, c.y+1))
	}
	if b.IsValidCell(c.x, c.y-1) {
		arrounds = append(arrounds, b.GetCell(c.x, c.y-1))
	}
	return arrounds
}