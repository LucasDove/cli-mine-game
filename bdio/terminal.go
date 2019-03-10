package bdio

import (
	"cli-mine-game/config"
	"fmt"
)

type TerminalOutput struct {

}

func (t *TerminalOutput) Output(bvalue [][]int8, x, y int32) {
	for i := 0; i < len(bvalue); i++ {
		row := ""
		for j := 0; j < len(bvalue[i]) ; j++ {
			cell := bvalue[i][j]
			if cell == config.DispMine || cell == config.DispSpace || cell == config.DispUndigged {
				row += fmt.Sprintf("%c  ", cell)
			}else {
				row += fmt.Sprintf("%d  ", cell)
			}
		}
		fmt.Println(row)
	}
}