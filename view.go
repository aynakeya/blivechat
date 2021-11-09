package blivechat

import (
	"github.com/aynakeya/gocui"
)

const (
	ViewRoom   = "room"
	ViewDanmu  = "danmu"
	ViewDebug  = "debug"
	ViewSend   = "send"
	ViewConfig = "config"
)

var ViewSequence = [...]string{ViewRoom, ViewDanmu, ViewDebug, ViewSend, ViewConfig}

func getNextView(viewName string) string {
	index := -1
	for i, v := range ViewSequence {
		if v == viewName {
			index = i
		}
	}
	index = (index + 1) % len(ViewSequence)
	return ViewSequence[index]
}

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func NextView(g *gocui.Gui, v *gocui.View) error {
	var vName string
	if v == nil {
		vName = ""
	} else {
		vName = v.Name()
	}
	if _, err := setCurrentViewOnTop(g, getNextView(vName)); err != nil {
		return err
	}
	if g.CurrentView().Editable {
		g.Cursor = true
	} else {
		g.Cursor = false
	}
	return nil
}

func InitEditView(g *gocui.Gui) error {
	if _, err := setCurrentViewOnTop(g, ViewSend); err != nil {
		return err
	}
	g.Cursor = true
	return nil
}
