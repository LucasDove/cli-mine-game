package board

import "cli-mine-game/config"

type CellReactor interface {
	Toggle(b BoardReactor) bool
	SweepArea(b BoardReactor)

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

	defer func() {c.isToggled = true}()

	if c.value == config.Mine {
		return false
	}else if c.value == config.Space {
		c.SweepArea(b)
	}else {
		b.MinusUntoggledMines()
		return true
	}

	return true
}

//保证调用处的cell一定是space
func (c *Cell) SweepArea(b BoardReactor) {
	if c.isToggled {
		return
	}
	if c.value == config.Space {
		//whitespaces
		b.MinusUntoggledMines()
		c.isToggled = true
		cells := c.aroundCells(b)
		for _, cell := range cells {
			cell.SweepArea(b)
		}
	}else if c.value != config.Mine {
		b.MinusUntoggledMines()
		c.isToggled = true
	}else {
		//mines
		return
	}
}

func (c *Cell) aroundCells(b BoardReactor) []CellReactor {
	var arrounds []CellReactor
	if config.IsValidCell(c.x+1, c.y) {
		arrounds = append(arrounds, b.GetCell(c.x+1, c.y))
	}
	if config.IsValidCell(c.x+1, c.y+1) {
		arrounds = append(arrounds, b.GetCell(c.x+1, c.y+1))
	}
	if config.IsValidCell(c.x+1, c.y-1) {
		arrounds = append(arrounds, b.GetCell(c.x+1, c.y-1))
	}

	if config.IsValidCell(c.x-1, c.y) {
		arrounds = append(arrounds, b.GetCell(c.x-1, c.y))
	}
	if config.IsValidCell(c.x-1, c.y+1) {
		arrounds = append(arrounds, b.GetCell(c.x-1, c.y+1))
	}
	if config.IsValidCell(c.x-1, c.y-1) {
		arrounds = append(arrounds, b.GetCell(c.x-1, c.y-1))
	}

	if config.IsValidCell(c.x, c.y+1) {
		arrounds = append(arrounds, b.GetCell(c.x, c.y+1))
	}
	if config.IsValidCell(c.x, c.y-1) {
		arrounds = append(arrounds, b.GetCell(c.x, c.y-1))
	}
	return arrounds
}