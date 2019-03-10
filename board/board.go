package board

import (
	"cli-mine-game/config"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type BoardReactor interface {
	MinusUntoggledMines()
	GetCell(x, y int32) CellReactor
	GameEnded() bool
	ProblemSolved() bool
	SetGameEnd()
	DisplayPending() [][]int8
	DisplayEnd() [][]int8
}

type Board struct {
	//长度
	length          int32
	//高度
	height          int32
	//格子对象
	cells           [][]CellReactor
	//未翻开的好格子，空白和数字
	unToggledSpaces int32
	//地雷的数量
	mines           int32
	//是否踩中了地雷
	steppedOnMine   bool
}

func NewBoard() *Board {
	b := &Board{
		length:    config.Bconfig.Length,
		height:    config.Bconfig.Height,
		mines:     config.Bconfig.Mines,
	}
	b.build()
	return b
}

func (b *Board) build() {
	b.unToggledSpaces = b.length*b.height - b.mines

	for i := int32(0); i < b.height; i++ {
		var row []CellReactor
		for j := int32(0); j < b.length ; j++ {
			cell := &Cell{
				x: i,
				y: j,
				value: config.Space,
				isToggled: false,
			}
			row = append(row, cell)
		}
		b.cells = append(b.cells, row)
	}

	fillsMap := make(map[string]bool)
	for ; int32(len(fillsMap)) < b.mines; {
		y := rand.Int31n(b.length)
		x := rand.Int31n(b.height)
		key := fmt.Sprintf("%d+%d", x, y)
		if _, ok := fillsMap[key]; !ok {
			b.cells[x][y].SetValue(config.Mine)
			fillsMap[key] = true
		}
	}

	for i := int32(0); i < b.height; i++ {
		for j := int32(0); j < b.length ; j++ {
			cell := b.cells[i][j]
			if cell.GetValue() != config.Mine {
				mines := int8(0)
				if j + 1 < b.length && b.cells[i][j+1].GetValue() == config.Mine {
					mines++
				}
				if j - 1 >= 0 && b.cells[i][j-1].GetValue() == config.Mine {
					mines++
				}
				if i + 1 < b.height && b.cells[i+1][j].GetValue() == config.Mine {
					mines++
				}
				if i - 1 >= 0 && b.cells[i-1][j].GetValue() == config.Mine {
					mines++
				}


				if j + 1 < b.length && i + 1 < b.height && b.cells[i+1][j+1].GetValue() == config.Mine {
					mines++
				}
				if j - 1 >= 0 && i + 1 < b.height && b.cells[i+1][j-1].GetValue() == config.Mine {
					mines++
				}
				if j + 1 < b.length && i - 1 >= 0 && b.cells[i-1][j+1].GetValue() == config.Mine {
					mines++
				}
				if j - 1 >= 0 && i - 1 >= 0 && b.cells[i-1][j-1].GetValue() == config.Mine {
					mines++
				}

				cell.SetValue(mines)
			}
		}
	}
}

func (b *Board) MinusUntoggledMines() {
	b.unToggledSpaces--
}

func (b *Board) GetCell(x, y int32) CellReactor {
	return b.cells[x][y]
}

/*func (b *Board) Listen() {
	for ; !b.GameEnded(); {
		x, y, err := b.inputer.Input()
		if err != nil {
			break
		}
		if config.IsValidCell(x, y) {
			if ! b.GetCell(x, y).Toggle(b) {
				b.SetGameEnd()
			}else {
				b.DisplayPending()
			}
		}else {
			fmt.Println("wrong cell, try again")
		}
	}
	b.DisplayEnd()
}*/

func (b *Board) DisplayPending() [][]int8 {
	var bvalue [][]int8

	for i := int32(0); i < b.height; i++ {
		var row []int8
		for j := int32(0); j < b.length ; j++ {
			cell := b.cells[i][j]
			if cell.IsToggled() {
				if cell.GetValue() == config.Mine {
					row = append(row, config.DispMine)
				}else if cell.GetValue() == config.Space {
					row = append(row, config.DispSpace)
				}else {
					row = append(row, cell.GetValue())
				}
			}else {
				row = append(row, config.DispUndigged)
			}
		}
		bvalue = append(bvalue, row)
	}

	return bvalue
}

func (b *Board) DisplayEnd() [][]int8 {
	var bvalue [][]int8

	for i := int32(0); i < b.height; i++ {
		var row []int8
		for j := int32(0); j < b.length ; j++ {
			cell := b.cells[i][j]
			if cell.GetValue() == config.Mine {
				row = append(row, config.DispMine)
			}else if cell.GetValue() == config.Space {
				row = append(row, config.DispSpace)
			}else {
				row = append(row, cell.GetValue())
			}
		}
		bvalue = append(bvalue, row)
	}
	return bvalue
}

func (b *Board) GameEnded() bool {
	return b.steppedOnMine || b.unToggledSpaces == 0
}

func (b *Board) SetGameEnd() {
	b.steppedOnMine = true
}

func (b *Board) ProblemSolved() bool {
	return b.unToggledSpaces == 0
}