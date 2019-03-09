package input

import "cli-mine-game/board"

type Trigger interface {
	RecvInput() (board.CellReactor, error)
}
