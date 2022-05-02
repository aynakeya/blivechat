package util

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"github.com/aynakeya/blivedm"
	"time"
)

func ViewPrintln(g *gocui.Gui, viewName string, a ...interface{}) {
	view, err := g.View(viewName)
	if err != nil {
		return
	}
	fmt.Fprintln(view, a...)
}

func ViewPrint(g *gocui.Gui, viewName string, a interface{}) {
	view, err := g.View(viewName)
	if err != nil {
		return
	}
	fmt.Fprint(view, a)
}

func ViewPrintWithTime(g *gocui.Gui, viewName string, a interface{}) {
	view, err := g.View(viewName)
	if err != nil {
		return
	}
	fmt.Fprintf(view, "%s >\n%s\n", time.Now().Format("2006/01/02 15:04:05"), a)
}

func PrintDanmu(g *gocui.Gui, viewName string, color bool, showMedal bool, msg blivedm.DanmakuMessage) {
	view, err := g.View(viewName)
	if err != nil {
		return
	}

	var medal, name, danmu string

	medal = ""
	name = msg.Uname
	danmu = msg.Msg

	if showMedal {
		medal = "[Unknown](0)"
		if len(msg.MedalName) > 0 {
			medal = fmt.Sprintf("[%s](%d)", msg.MedalName, msg.MedalLevel)
		}
	}

	if color {
		medal = SetForegroundColor(IntToRGB(int(msg.MedalColor)), medal)
		name = SetForegroundColor(HexToRGB(msg.UnameColor), name)
		danmu = SetForegroundColor(IntToRGB(int(msg.Color)), danmu)
	}

	fmt.Fprintln(view, fmt.Sprintf("%s %s: %s", medal, name, danmu))
}
