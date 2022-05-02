package widget

import (
	"blivechat/util"
	"github.com/awesome-gocui/gocui"
)

type CommonPanel struct {
	BaseWidget
	LinkedWidget
	SwitchKey KeyCombo
}

func (w *CommonPanel) Layout(g *gocui.Gui) error {
	x0, y0, x1, y1 := w.GetSize(g.Size())
	if v, err := g.SetView(w.ViewName, x0, y0, x1, y1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = w.DisplayName
		v.Wrap = true
		v.Editable = false
		v.Autoscroll = true
		return w.setKeybindings(g, v)
	}
	return nil
}

func (w *CommonPanel) setKeybindings(g *gocui.Gui, v *gocui.View) (err error) {
	err = g.SetKeybinding(w.ViewName, w.SwitchKey.Key, w.SwitchKey.Modifier, func(gui *gocui.Gui, view *gocui.View) error {
		if w.NextViewName == "" {
			return nil
		}
		if _, err := util.SetCurrentViewOnTop(g, w.NextViewName); err != nil {
			return err
		}
		if g.CurrentView().Editable {
			g.Cursor = true
		} else {
			g.Cursor = false
		}
		return nil
	})
	if err != nil {
		return
	}
	return
}
