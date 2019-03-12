package bdio

import (
	"cli-mine-game/board"
	"cli-mine-game/config"
	"errors"
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

var (
	Gui *gocui.Gui
	bboard board.BoardReactor
	output = &GuiOutput{}

	exitCountDown = 1
)

func InitGui(b board.BoardReactor) {
	var err error
	Gui, err = gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}
	defer Gui.Close()

	Gui.Cursor = true
	Gui.Mouse = true
	bboard = b

	Gui.SetManagerFunc(layout)

	if err := Gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := Gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func getNameFromCoordinate(x, y int) string {
	return fmt.Sprintf("cell %d %d", x, y)
}

func extCoordinateFromName(name string) (x, y int32) {
	var prefix string
	var xzb int32
	var yzb int32
	_, err := fmt.Sscanf(name, "%s %d %d", &prefix, &xzb, &yzb)
	if err != nil {
		log.Panicf("%s onclick, err:%+v", name, err)
	}
	return xzb, yzb
}

func layout(g *gocui.Gui) error {
	for i := 0; i < int(config.Bconfig.Height); i++ {
		var cellHeight int = 2
		var cellLength int = 4
		for j := 0; j < int(config.Bconfig.Length); j++ {
			name := getNameFromCoordinate(i, j)
			v, err := g.SetView(name, j * cellLength, i * cellHeight, (j+1) * cellLength, (i+1) * cellHeight)
			if err != nil && err != gocui.ErrUnknownView {
				return err
			}
			_, err = fmt.Fprintln(v, " ")
			if err != nil {
				fmt.Printf("fill err:%+v\n", err)
			}

			if err := Gui.SetKeybinding(v.Name(), gocui.MouseLeft, gocui.ModNone, onclick); err != nil {
				log.Panicln(err)
			}
		}
	}

	return nil
}

func onclick(g *gocui.Gui, v *gocui.View) error {
	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}
	xzb, yzb := extCoordinateFromName(v.Name())

	if bboard.GameEnded() {
		return quit(g, v)
	}
	cell := bboard.GetCell(xzb, yzb)
	if cell.Toggle(bboard) {
		output.Output(bboard.DisplayPending())
	}else {
		bboard.SetGameEnd()
		output.Output(bboard.DisplayEnd())
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	if exitCountDown == 0 {
		return gocui.ErrQuit
	}
	exitCountDown--

	maxX, maxY := g.Size()
	v, err := g.SetView("script", -1, maxY - 10, maxX, maxY)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	text := ""
	if bboard.ProblemSolved() {
		text = "congrats, you solved the problem"
	}else {
		text = "sorry, you stepped on the mine"
	}
	text += ". exit in next time"

	_, err = fmt.Fprintln(v, text)
	if err != nil {
		fmt.Printf("fill err:%+v\n", err)
	}

	return nil
}

type GuiOutput struct {

}

func (g *GuiOutput) Output(bvalue [][]int8) {
	for i := 0; i < int(config.Bconfig.Height); i++ {
		for j := 0; j < int(config.Bconfig.Length); j++ {
			cell := bvalue[i][j]
			name := getNameFromCoordinate(i, j)
			v, err := Gui.View(name)
			if err != nil && err != gocui.ErrUnknownView {
				log.Panicf("wrong cell name:%s", name)
			}
			v.Clear()
			if cell == config.DispMine || cell == config.DispSpace || cell == config.DispUndigged {
				_, err = fmt.Fprintf(v, " %c", cell)
				if cell == config.DispMine {
					v.SelBgColor = gocui.ColorGreen
				}
			}else {
				_, err = fmt.Fprintf(v, " %d", cell)
			}
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