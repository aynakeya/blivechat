package blivechat

import (
	"fmt"
	"github.com/aynakeya/blivedm"
	"github.com/aynakeya/gocui"
	"time"
)

func viewPrintln(v *gocui.View, a ...interface{}) {
	fmt.Fprintln(v, a...)
}

func viewPrint(v *gocui.View, a interface{}) {
	fmt.Fprint(v, a)
}

func viewPrintWithTime(v *gocui.View, a interface{}) {
	fmt.Fprintf(v, "%s >\n%s\n", time.Now().Format("2006/01/02 15:04:05"), a)
}

func PrintToDebug(g *gocui.Gui, a interface{}) {
	view, err := g.View(ViewDebug)
	if err != nil {
		return
	}
	viewPrintWithTime(view, a)
}

func printDanmuColor(v *gocui.View, msg blivedm.DanmakuMessage) {
	name := SetForegroundColor(HexToRGB(msg.UnameColor), msg.Uname)
	medal := "[Unknown](0)"
	if len(msg.MedalName) > 0 {
		medal = SetForegroundColor(IntToRGB(int(msg.MedalColor)),
			fmt.Sprintf("[%s](%d)", msg.MedalName, msg.MedalLevel))
	}
	viewPrintln(v,
		fmt.Sprintf("%s %s: %s",
			medal, name, SetForegroundColor(IntToRGB(int(msg.Color)), msg.Msg)))
}

func printDanmuNoColor(v *gocui.View, msg blivedm.DanmakuMessage) {
	name := msg.Uname
	medal := "[Unknown](0)"
	if len(msg.MedalName) > 0 {
		medal = fmt.Sprintf("[%s](%d)", msg.MedalName, msg.MedalLevel)
	}
	viewPrintln(v,
		fmt.Sprintf("%s %s: %s",
			medal, name, msg.Msg))
}

func PrintDanmu(v *gocui.View, msg blivedm.DanmakuMessage) {
	if Config.VisualColorMode {
		printDanmuColor(v, msg)
	} else {
		printDanmuNoColor(v, msg)
	}
}
