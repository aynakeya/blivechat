package widget

import (
	"blivechat/util"
	"fmt"
	"github.com/awesome-gocui/gocui"
)

type ConfigOption struct {
	index        int
	Options      []string
	OptionValues []interface{}
}

func NewConfigOption(values map[string]interface{}) ConfigOption {
	options, opvalues := make([]string, 0), make([]interface{}, 0)
	for key, val := range values {
		options = append(options, key)
		opvalues = append(opvalues, val)
	}
	return ConfigOption{
		index:        0,
		Options:      options,
		OptionValues: opvalues,
	}
}

func (c *ConfigOption) Prev() string {
	c.index--
	if c.index < 0 {
		c.index = len(c.Options) - 1
	}
	return c.Current()
}

func (c *ConfigOption) Current() string {
	return c.Options[c.index]
}

func (c *ConfigOption) Next() string {
	c.index = (c.index + 1) % len(c.Options)
	return c.Current()
}

func (c *ConfigOption) Value() interface{} {
	return c.OptionValues[c.index]
}

func (c *ConfigOption) SetByValue(value interface{}) {
	for index, val := range c.OptionValues {
		if val == value {
			c.index = index
		}
	}
}

type ConfigOptionPanel struct {
	BaseWidget
	DoubleLinkedWidget
	Option    ConfigOption
	SetConfig func(value interface{}, origin *ConfigOptionPanel)
}

func (w *ConfigOptionPanel) Layout(g *gocui.Gui) error {
	x0, y0, x1, y1 := w.GetSize(g.Size())
	if v, err := g.SetView(w.ViewName, x0, y0, x1, y1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = w.DisplayName
		v.Wrap = true
		v.Editable = false
		fmt.Fprintf(v, w.Option.Current())
		return w.setKeybindings(g, v)
	}
	return nil
}

func (w *ConfigOptionPanel) setKeybindings(g *gocui.Gui, v *gocui.View) error {
	if err := g.SetKeybinding(w.ViewName, gocui.KeyArrowRight, gocui.ModNone, func(gui *gocui.Gui, view *gocui.View) error {
		if w.NextViewName == "" {
			return nil
		}
		if _, err := util.SetCurrentViewOnTop(g, w.NextViewName); err != nil {
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
		if _, err := util.SetCurrentViewOnTop(g, w.PrevViewName); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	if err := g.SetKeybinding(w.ViewName, gocui.KeyArrowUp, gocui.ModNone, func(gui *gocui.Gui, view *gocui.View) error {
		v.Clear()
		fmt.Fprintf(view, w.Option.Prev())
		w.SetConfig(w.Option.Value(), w)
		return nil
	}); err != nil {
		return err
	}
	if err := g.SetKeybinding(w.ViewName, gocui.KeyArrowDown, gocui.ModNone, func(gui *gocui.Gui, view *gocui.View) error {
		v.Clear()
		fmt.Fprintf(view, w.Option.Next())
		w.SetConfig(w.Option.Value(), w)
		return nil
	}); err != nil {
		return err
	}
	if err := g.SetKeybinding(w.ViewName, gocui.KeyEnter, gocui.ModNone, func(gui *gocui.Gui, view *gocui.View) error {
		_, err := util.SetCurrentViewOnTop(g, w.ParentViewName)
		return err
	}); err != nil {
		return err
	}
	return nil
}
