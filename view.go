package blivechat

import (
	"github.com/awesome-gocui/gocui"
	"strings"
)

const (
	ViewRoom   = "room"
	ViewDanmu  = "danmu"
	ViewDebug  = "debug"
	ViewSend   = "send"
	ViewConfig = "config"
)

const (
	ViewConfigVisualColorMode = "config.color"
	ViewConfigDanmuColor      = "config.danmu.color"
	ViewConfigDanmuMode       = "config.danmu.mode"
)

var ViewSequence = [...]string{ViewRoom, ViewDanmu, ViewDebug, ViewSend, ViewConfig}
var ViewSequenceConfig = [...]string{ViewConfigVisualColorMode, ViewConfigDanmuColor, ViewConfigDanmuMode}

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func getNextViewName(viewName string, targetViewSequence []string) string {
	index := -1
	for i, v := range targetViewSequence {
		if v == viewName {
			index = i
		}
	}
	index = (index + 1) % len(targetViewSequence)
	return targetViewSequence[index]
}

func nextView(g *gocui.Gui, v *gocui.View, targetViewSequence []string) error {
	var vName string
	if v == nil {
		vName = ""
	} else {
		vName = v.Name()
	}
	if _, err := setCurrentViewOnTop(g, getNextViewName(vName, targetViewSequence)); err != nil {
		return err
	}
	if g.CurrentView().Editable {
		g.Cursor = true
	} else {
		g.Cursor = false
	}
	return nil
}

func NextView(g *gocui.Gui, v *gocui.View) error {
	if v == nil || strings.HasPrefix(v.Name(), "config") {
		if _, err := g.SetViewOnTop(ViewConfig); err != nil {
			return err
		}
	}
	if err := nextView(g, v, ViewSequence[:]); err != nil {
		return err
	}
	if g.CurrentView().Name() == ViewConfig {
		if _, err := g.SetViewOnTop(ViewSequenceConfig[0]); err != nil {
			return err
		}
	}
	return nil

}

func NextConfigView(g *gocui.Gui, v *gocui.View) error {
	return nextView(g, v, ViewSequenceConfig[:])
}
