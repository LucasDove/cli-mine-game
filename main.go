package main

import (
	"cli-mine-game/board"
	"cli-mine-game/input"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var length, width int64 = 8,8
	var mineRate float64 = 0.3
	if len(os.Args) > 2 {
		var err error
		length, err = strconv.ParseInt(os.Args[1], 10, 32)
		if err != nil {
			fmt.Println("input valid length")
		}
		width, err = strconv.ParseInt(os.Args[2], 10, 32)
		if err != nil {
			fmt.Println("input valid width")
		}
		mineRate, err = strconv.ParseFloat(os.Args[3], 10)
		if err != nil {
			fmt.Println("input valid mineRate")
		}
		mineRate *= 10
	}

	b := board.NewBoard(int32(length), int32(width), mineRate, &input.KeyBoard{})
	b.Listen()

	return
}
