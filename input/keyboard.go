package input

import (
	"errors"
	"fmt"
	"strconv"
)

type Trigger interface {
	RecvInput() (x, y int32, err error)
}

type KeyBoard struct {
}

func (k *KeyBoard) RecvInput() (x, y int32, err error) {
	fmt.Println("take a guess, input coordinate:")
	var lenth, height string
	cnt, cerr := fmt.Scanln(&height, &lenth)
	if cerr != nil || cnt < 2 {
		return 0, 0, errors.New("input error, retry again")
	}

	l, cerr := strconv.ParseInt(lenth, 10, 32)
	if cerr != nil {
		return 0, 0, errors.New("length error, retry again")
	}
	h, cerr := strconv.ParseInt(height, 10, 32)
	if cerr != nil {
		return 0, 0, errors.New("height error, retry again")
	}
	y = int32(l)
	x = int32(h)
	err = nil
	return
}