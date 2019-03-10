package board

import (
	"cli-mine-game/bdio"
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
	SetGameEnd()
	Listen()
	DisplayPending()
	DisplayEnd()
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

	//输入器
	inputer   bdio.InputReceiver
	//输出器
	outputter bdio.OutputReceiver
}

func NewBoard(inputer bdio.InputReceiver, outputter bdio.OutputReceiver) *Board {
	b := &Board{
		length:    config.Bconfig.Length,
		height:    config.Bconfig.Height,
		mines:     config.Bconfig.Mines,
		inputer:   inputer,
		outputter: outputter,
	}
	b.build()
	b.DisplayEnd()
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

func (b *Board) Listen() {
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
}

func (b *Board) DisplayPending() {
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

	b.outputter.Output(bvalue, 0, 0)
}

func (b *Board) DisplayEnd() {
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
	b.outputter.Output(bvalue, 0, 0)
}

func (b *Board) GameEnded() bool {
	end := b.steppedOnMine || b.unToggledSpaces == 0
	if end {
		if b.unToggledSpaces == 0 {
			fmt.Println("grats, you have solve it")
		}else {
			fmt.Println("sorry, you have stepped on the mine")
		}
	}
	return end
}

func (b *Board) SetGameEnd() {
	b.steppedOnMine = true
}