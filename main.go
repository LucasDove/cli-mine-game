package main

import (
	"cli-mine-game/bdio"
	"cli-mine-game/board"
	"cli-mine-game/config"
	"log"
)

func main() {

	err := config.BuildFromStdin()
	if err != nil {
		log.Panicf("bad config, err:%+v", err)
	}

	b := board.NewBoard()
	bdio.InitGui(b)

	return
}
