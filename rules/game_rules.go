package rules

import (
	"cli-mine-game/board"
)

func OnCellTrigger(board board.BoardReactor, cell board.CellReactor) {
	if !cell.ActionAccepted() {
		return
	}

	cell.ChangeState(board)

	board.Display()
}