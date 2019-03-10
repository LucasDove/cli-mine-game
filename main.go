package main

import (
	"cli-mine-game/bdio"
	"cli-mine-game/board"
	"cli-mine-game/config"
	"fmt"
	"log"
)

func main() {

	err := config.BuildFromStdin()
	if err != nil {
		log.Panicf("bad config, err:%+v", err)
	}

	b := board.NewBoard()
	bdio.InitGui(b)
	if b.ProblemSolved() {
		fmt.Println("congrats, you have solved this problem")
	}else {
		fmt.Println("sorry, you have stepped on the mine")
	}

	/*b := board.NewBoard(&bdio.KeyBoardInput{}, &bdio.TerminalOutput{})
	b.Listen()*/

	return
}
