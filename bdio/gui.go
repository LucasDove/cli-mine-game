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
	_, maxY := g.Size()
	for i := 0; i < int(config.Bconfig.Height); i++ {
		var cellWidth int = maxY / int(config.Bconfig.Height)
		for j := 0; j < int(config.Bconfig.Length); j++ {
			name := fmt.Sprintf("cell%d%d", i, j)
			v, err := g.SetView(name, j * cellWidth, i * cellWidth, (j+1) * cellWidth, (i) * cellWidth + 3)
			if err != nil && err != gocui.ErrUnknownView {
				return err
			}
			x, y := v.Size()
			_, _ = fmt.Fprintf(v, "%d, %d", x, y)
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
	view, err := Gui.View("cell00")
	if err == nil {
		var buf []byte
		buf = append(buf, byte(bvalue[0][0]))
		_, _ = view.Write(buf)
	}
}

type GuiInput struct {

}

func (g *GuiInput) Input() (x, y int32, err error) {
	return 0, 0, errors.New("战术失败")
}