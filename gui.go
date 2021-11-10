package blivechat

import (
	"github.com/aynakeya/gocui"
	"log"
)

func CreateGUI() *gocui.Gui {
	g, err := gocui.NewGui(gocui.OutputTrue, false)
	if err != nil {
		log.Panicln(err)
	}
	g.Highlight = true
	g.SelFgColor = gocui.ColorWhite
	g.SelFrameColor = gocui.ColorBlue
	addWidgets(g)
	setKeybindings(g)
	return g
}

func addWidgets(g *gocui.Gui) {
	managers := ConfigLayouts(g)
	managers = append(managers, gocui.ManagerFunc(MainLayout))
	g.SetManager(managers...)
}
