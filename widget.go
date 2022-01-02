package blivechat

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
)

type DynamicSizeFunc func(g *gocui.Gui) (x0, y0, x1, y1 int)

type BaseWidget struct {
	ViewName string
	GetSize  DynamicSizeFunc
}

type LinkedWidget struct {
	PrevViewName string
	NextViewName string
}

type ConfigOptionPanel struct {
	BaseWidget
	LinkedWidget
	DisplayName string
	Option      ConfigOption
	SetConfig   func(value string)
}

func (w *ConfigOptionPanel) Layout(g *gocui.Gui) error {
	x0, y0, x1, y1 := w.GetSize(g)
	if v, err := g.SetView(w.ViewName, x0, y0, x1, y1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = w.DisplayName
		v.Wrap = true
		v.Editable = false
		viewPrint(v, w.Option.Current())
		return w.setKeybindings(g, v)
	}
	return nil
}

func (w *ConfigOptionPanel) setKeybindings(g *gocui.Gui, v *gocui.View) error {
	if err := g.SetKeybinding(w.ViewName, gocui.KeyArrowRight, gocui.ModNone, func(gui *gocui.Gui, view *gocui.View) error {
		if w.NextViewName == "" {
			return nil
		}
		if _, err := setCurrentViewOnTop(g, w.NextViewName); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	if err := g.SetKeybinding(w.ViewName, gocui.KeyArrowLeft, gocui.ModNone, func(gui *gocui.Gui, view *gocui.View) error {
		if w.PrevViewName == "" {
			return nil
		}
		if _, err := setCurrentViewOnTop(g, w.PrevViewName); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	if err := g.SetKeybinding(w.ViewName, gocui.KeyArrowUp, gocui.ModNone, func(gui *gocui.Gui, view *gocui.View) error {
		v.Clear()
		viewPrint(v, w.Option.Prev())
		w.SetConfig(w.Option.Value())
		PrintToDebug(g, fmt.Sprintf("Set Config %s to %s(%s)", w.DisplayName, w.Option.Current(), w.Option.Value()))
		return nil
	}); err != nil {
		return err
	}
	if err := g.SetKeybinding(w.ViewName, gocui.KeyArrowDown, gocui.ModNone, func(gui *gocui.Gui, view *gocui.View) error {
		v.Clear()
		viewPrint(v, w.Option.Next())
		w.SetConfig(w.Option.Value())
		PrintToDebug(g, fmt.Sprintf("Set Config %s to %s(%s)", w.DisplayName, w.Option.Current(), w.Option.Value()))
		return nil
	}); err != nil {
		return err
	}
	return nil
}
