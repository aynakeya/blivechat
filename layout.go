package blivechat

import (
	"github.com/aynakeya/gocui"
	"log"
)

func Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(ViewRoom, 0, 0, maxX-1, maxY/8-1,0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "RoomInfo"
		v.Wrap = true
		v.Editable = false

	}

	if v, err := g.SetView(ViewDanmu, 0, maxY/8, maxX*5/8-1, maxY*6/8-1,0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Danmu"
		v.Wrap = true
		v.Editable = false
		v.Autoscroll = true
	}
	if v, err := g.SetView(ViewDebug, maxX*5/8, maxY/8, maxX-1, maxY*6/8-1,0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Debug"
		v.Wrap = true
		v.Editable = false
		v.Autoscroll = true
		log.SetOutput(v)
	}

	if v, err := g.SetView(ViewSend, 0, maxY*6/8, maxX*5/8-1, maxY-1,0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Input"
		v.Wrap = true
		v.Editable = true
		v.Autoscroll = true
	}
	if v, err := g.SetView(ViewConfig, maxX*5/8, maxY*6/8, maxX-1, maxY-1,0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Config"
		v.Wrap = true
		v.Editable = false
		v.Autoscroll = true
	}
	return nil
}
