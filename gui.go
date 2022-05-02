package blivechat

import (
	"github.com/awesome-gocui/gocui"
	"log"
)

var MainGui *gocui.Gui
var MainManager []gocui.Manager

func initializeGUI() error {
	view, err := MainGui.View(ViewSend)
	if err != nil {
		return err
	}
	view.Editable = true
	return nil
}

func CreateGUI() (err error) {
	MainGui, err = gocui.NewGui(gocui.OutputTrue, false)
	MainManager = make([]gocui.Manager, 0)
	if err != nil {
		log.Panicln(err)
	}
	MainGui.Highlight = true
	MainGui.SelFgColor = gocui.ColorWhite
	MainGui.SelFrameColor = gocui.ColorBlue
	createLayoutManager()
	MainGui.SetManager(MainManager...)
	setKeyBindings()

	MainGui.Update(func(gui *gocui.Gui) error {
		if err := initializeGUI(); err != nil {
			return err
		}
		return nil
	})
	return
}
