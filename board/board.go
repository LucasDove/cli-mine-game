package board

import (
	"cli-mine-game/input"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type BoardReactor interface {
	build()
	IsValidCell(x, y int32) bool

	Listen()
	Display()
	OnCellChange()
}

type Board struct {
	length int32
	height int32

	cells [][]int8

	inputer input.Trigger
}

func NewBoard(l, w int32) *Board {
	b := &Board{
		length: l,
		height: w,
	}
	b.build()
	return b
}

func (b *Board) build() {
	for i := int32(0); i < b.height; i++ {
		tmp := make([]int8, b.length)
		b.cells = append(b.cells, tmp)
	}
}

func (b *Board) IsValidCell(x, y int32) bool {
	return y  > b.height || x > b.length || y < 1 || x <= 0
}

func (b *Board) IsValidCellInObj(cell *Cell) bool {
	return cell.y  > b.height || cell.x > b.length ||
		cell.y < 1 || cell.x <= 0
}

func (b *Board) Listen() {
	for ; ; {
		cell, err := b.inputer.RecvInput()
	}
}

func (b *Board) Display() {

}

func (b *Board) OnCellChange() {

}
