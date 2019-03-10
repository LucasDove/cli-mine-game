package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type BoardConfig struct {
	Length int32
	Height int32

	Mines  int32
}

var (
	Bconfig *BoardConfig
)

func BuildFromStdin() (cerr error) {
	height := int32(8)
	length := int32(8)
	mines := int32(10)
	if len(os.Args) > 2 {
		height64, err := strconv.ParseInt(os.Args[1], 10, 32)
		if err != nil {
			fmt.Println("bdio valid length")
		}
		length64, err := strconv.ParseInt(os.Args[2], 10, 32)
		if err != nil {
			fmt.Println("bdio valid width")
		}
		mines64, err := strconv.ParseFloat(os.Args[3], 10)
		if err != nil {
			fmt.Println("bdio valid mineRate")
		}
		height = int32(height64)
		length = int32(length64)
		mines = int32(mines64)

		if height > 16 || length > 16 {
			cerr = errors.New("invalid param")
		}
		Bconfig = &BoardConfig{height, length, mines}
	}
	return
}

func IsValidCell(x, y int32) bool {
	return y < Bconfig.Length && x < Bconfig.Height && y >= 0 && x >= 0
}