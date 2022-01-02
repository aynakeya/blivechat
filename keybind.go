package blivechat

import (
	"github.com/awesome-gocui/gocui"
	"log"
)

func setKeybindings(gui *gocui.Gui) {
	if err := gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := gui.SetKeybinding("", gocui.KeyTab, gocui.ModNone, NextView); err != nil {
		log.Panicln(err)
	}
	if err := gui.SetKeybinding(ViewConfig, gocui.KeyArrowRight, gocui.ModNone, NextConfigView); err != nil {
		log.Panicln(err)
	}
}
