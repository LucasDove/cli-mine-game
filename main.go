package main

import (
	"cli-mine-game/board"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var length, width int64 = 8,8
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
	}

	b := board.NewBoard(int32(length), int32(width))
	b.Listen()

	return
}
