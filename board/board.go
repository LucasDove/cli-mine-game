package board

import (
	"cli-mine-game/input"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type BoardReactor interface {
	IsValidCell(x, y int32) bool
	GetCell(x, y int32) CellReactor
	GameEnded() bool
	SetGameEnd()
	Listen()
	DisplayPending()
	DisplayEnd()
}

type Board struct {
	length int32
	height int32

	cells           [][]CellReactor
	unToggledCelled int32
	mineNum         int32
	steppedOnMine   bool

	inputer input.Trigger
}

func NewBoard(l, w int32, inputer input.Trigger) *Board {
	b := &Board{
		length: l,
		height: w,
		inputer: inputer,
	}
	b.build()
	return b
}

func (b *Board) build() {
	for i := int32(0); i < b.height; i++ {
		var row []CellReactor
		for j := int32(0); j < b.length ; j++ {
			cell := &Cell{
				x: i,
				y: j,
				value: 0,
				isToggled: false,
			}
			row = append(row, cell)
		}
		b.cells = append(b.cells, row)
	}
	b.unToggledCelled = b.length * b.height
	b.mineNum = b.unToggledCelled / 4

	fillsMap := make(map[string]bool)
	for ; int32(len(fillsMap)) < b.mineNum; {
		x := rand.Int31n(b.length)
		y := rand.Int31n(b.height)
		key := fmt.Sprintf("%d+%d", x, y)
		if _, ok := fillsMap[key]; !ok {
			b.cells[x][y].SetValue(Mine)
			fillsMap[key] = true
		}
	}

	for i := int32(0); i < b.height; i++ {
		for j := int32(0); j < b.length ; j++ {
			cell := b.cells[i][j]
			if cell.GetValue() != Mine {
				mines := int8(0)
				if j + 1 < b.length && b.cells[i][j+1].GetValue() == Mine {
					mines++
				}
				if j - 1 < b.length && j - 1 >= 0 && b.cells[i][j-1].GetValue() == Mine {
					mines++
				}
				if i + 1 < b.height && b.cells[i+1][j].GetValue() == Mine {
					mines++
				}
				if i - 1 < b.height && i - i >- 0 && b.cells[i-1][j].GetValue() == Mine {
					mines++
				}
				cell.SetValue(mines)
			}
		}
	}
}

func (b *Board) IsValidCell(x, y int32) bool {
	return y > b.height || x > b.length || y < 1 || x <= 0
}

func (b *Board) GetCell(x, y int32) CellReactor {
	return b.cells[x][y]
}

func (b *Board) Listen() {
	for ; !b.GameEnded(); {
		x, y, err := b.inputer.RecvInput()
		if err != nil {
			continue
		}
		if b.IsValidCell(x, y) {
			if ! b.GetCell(x, y).Toggle(b) {
				b.SetGameEnd()
			}
			b.DisplayPending()
		}else {
			fmt.Println("wrong cell, try again")
		}
	}
	b.DisplayEnd()
}

func (b *Board) DisplayPending() {
	row := ""
	for i := int32(0); i < b.height; i++ {
		for j := int32(0); j < b.length ; j++ {
			cell := b.cells[i][j]
			if cell.IsToggled() {
				if cell.GetValue() == Mine {
					row += "# "
				}else if cell.GetValue() == 0 {
					row += "_ "
				}else {
					row += fmt.Sprintf("%d", cell.GetValue())
				}
			}else {
				row += "* "
			}
		}
		row += `\n`
	}
	fmt.Print(row)
}

func (b *Board) DisplayEnd() {
	row := ""
	for i := int32(0); i < b.height; i++ {
		for j := int32(0); j < b.length ; j++ {
			cell := b.cells[i][j]
			if cell.GetValue() == Mine {
				row += "# "
			}else if cell.GetValue() == 0 {
				row += "_ "
			}else {
				row += fmt.Sprintf("%d", cell.GetValue())
			}
		}
		row += `\n`
	}
	fmt.Print(row)
}

func (b *Board) GameEnded() bool {
	return b.steppedOnMine
}

func (b *Board) SetGameEnd() {
	b.steppedOnMine = true

	if b.unToggledCelled == b.mineNum {
		fmt.Println("grats, you have solve it")
	}else {
		fmt.Println("sorry, you have stepped on the mine")
	}
}