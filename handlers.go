package blivechat

import (
	"github.com/aynakeya/gocui"
)

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
