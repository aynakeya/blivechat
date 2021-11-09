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
	g.Cursor = true

	g.SelFgColor = gocui.ColorWhite
	g.SelFrameColor = gocui.ColorBlue
	g.SetManagerFunc(Layout)
	g.Update(func(gui *gocui.Gui) error {
		return nil
	})
	return g
}
