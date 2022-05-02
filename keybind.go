package blivechat

import (
	"blivechat/util"
	"fmt"
	"github.com/awesome-gocui/gocui"
	"log"
)

func setKeyBindings() {
	setMainKeyBindings()
	setConfigKeyBindings()
	setSendKeyBindings()
}

func setMainKeyBindings() {
	if err := MainGui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := MainGui.SetKeybinding("", gocui.KeyTab, gocui.ModNone, func(gui *gocui.Gui, view *gocui.View) error {
		_, err := gui.SetCurrentView(ViewRoom)
		return err
	}); err != nil {
		log.Panicln(err)
	}
}

func setConfigKeyBindings() {
	if err := MainGui.SetKeybinding(ViewConfig, gocui.KeyEnter, gocui.ModNone, func(gui *gocui.Gui, view *gocui.View) error {
		util.SetCurrentViewOnTop(gui, ViewConfigVisualColorMode)
		return nil
	}); err != nil {
		return
	}
}

func setSendKeyBindings() {
	if err := MainGui.SetKeybinding(ViewSend, gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		msg := v.Buffer()
		if len(msg) == 0 {
			return nil
		}
		v.Clear()
		util.ViewPrintWithTime(MainGui, ViewDebug, fmt.Sprintf("try send msg: %s", msg))
		if Client.Account.UID == 0 {
			util.ViewPrintWithTime(MainGui, ViewDebug, "Send Msg fail, not login")
			return v.SetCursor(0, 0)
		}
		resp, err := sendDanmuMsg(msg)
		if err != nil {
			util.ViewPrintWithTime(MainGui, ViewDebug, fmt.Sprintf("Send Msg fail, %s", err))
			return v.SetCursor(0, 0)
		}
		util.ViewPrintWithTime(MainGui, ViewDebug, fmt.Sprintf("send result - code:%d msg:%s", resp.Code, resp.Message))
		return v.SetCursor(0, 0)
	}); err != nil {
		return
	}
}
