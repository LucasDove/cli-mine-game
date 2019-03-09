package input

import (
	"errors"
	"fmt"
	"strconv"
)

type KeyBoard struct {
}

func (k *KeyBoard) RecvInput() (x, y int32, err error) {
	var lenthStr, widthStr string
	cnt, cerr := fmt.Scanln(&lenthStr, &widthStr)
	if cerr != nil || cnt < 2 {
		return 0, 0, errors.New("input error, retry again")
	}

	l, cerr := strconv.ParseInt(lenthStr, 10, 32)
	if cerr != nil {
		return 0, 0, errors.New("length error, retry again")
	}
	h, cerr := strconv.ParseInt(lenthStr, 10, 32)
	if cerr != nil {
		return 0, 0, errors.New("height error, retry again")
	}
	x = int32(l)
	y = int32(h)
	err = nil
	return
}