package input

import (
	"cli-mine-game/board"
	"errors"
	"fmt"
	"strconv"
)

type KeyBoard struct {
	b board.BoardReactor
}

func (k *KeyBoard) RecvInput() (board.CellReactor, error) {
	var lenthStr, widthStr string
	cnt, err := fmt.Scanln(&lenthStr, &widthStr)
	if err != nil || cnt < 2 {
		return nil, errors.New("input error, retry again")
	}

	l, err := strconv.ParseInt(lenthStr, 10, 32)
	if err != nil {
		return nil, errors.New("length error, retry again")
	}
	h, err := strconv.ParseInt(lenthStr, 10, 32)
	if err != nil {
		return nil, errors.New("height error, retry again")
	}

	if k.b.IsValidCell(int32(l), int32(h)) {
		return nil, errors.New("invalid cell, retry again")
	}

	return board.NewCell(int32(l), int32(h)), nil
}