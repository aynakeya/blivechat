package blivechat

import (
	"github.com/aynakeya/gocui"
	"log"
)

func SetKeybindings(gui *gocui.Gui)  {
	if err := gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := gui.SetKeybinding("", gocui.KeyTab, gocui.ModNone, NextView); err != nil {
		log.Panicln(err)
	}
}
