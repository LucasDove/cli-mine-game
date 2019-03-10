package bdio

import (
	"cli-mine-game/config"
	"errors"
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

var (
	Gui *gocui.Gui
)

func InitGui() {
	var err error
	Gui, err = gocui.NewGui(gocui.Output256)
	Gui.ASCII = false
	if err != nil {
		log.Panicln(err)
	}
	defer Gui.Close()

	Gui.SetManagerFunc(layout)

	if err := Gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := Gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	for i := 0; i < int(config.Bconfig.Height); i++ {
		var cellHeight int = 2
		var cellLength int = 4
		for j := 0; j < int(config.Bconfig.Length); j++ {
			name := fmt.Sprintf("cell%d%d", i, j)
			_, err := g.SetView(name, j * cellLength, i * cellHeight, (j+1) * cellLength, (i+1) * cellHeight)
			if err != nil && err != gocui.ErrUnknownView {
				return err
			}
		}
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

type GuiOutput struct {

}

func (g *GuiOutput) Output(bvalue [][]int8, x, y int32) {
	log.Println("runs here....")
	for i := 0; i < int(config.Bconfig.Height); i++ {
		for j := 0; j < int(config.Bconfig.Length); j++ {
			cell := bvalue[i][j]
			name := fmt.Sprintf("cell%d%d", i, j)
			v, err := Gui.View(name)
			if err != nil && err != gocui.ErrUnknownView {
				log.Panicf("wrong cell name:%s", name)
			}
			//v.Clear()
			var value string
			if cell == config.DispMine || cell == config.DispSpace || cell == config.DispUndigged {
				value = fmt.Sprintf("%c", cell)
			}else {
				value = fmt.Sprintf("%d", cell)
			}
			fmt.Printf("fill %s with value: %s\n", name, value)
			_, err = fmt.Fprint(v, []byte(value))
			if err != nil {
				fmt.Printf("fill err:%+v\n", err)
			}
		}
	}
}

type GuiInput struct {

}

func (g *GuiInput) Input() (x, y int32, err error) {
	return 0, 0, errors.New("战术失败")
}