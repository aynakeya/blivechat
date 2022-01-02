package blivechat

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"github.com/aynakeya/blivedm"
	"github.com/spf13/cast"
	"log"
)

func MainLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(ViewRoom, 0, 0, maxX-1, maxY/8-1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "RoomInfo"
		v.Wrap = true
		v.Editable = false
	}

	if v, err := g.SetView(ViewDanmu, 0, maxY/8, maxX*5/8-1, maxY*6/8-1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Danmu"
		v.Wrap = true
		v.Editable = false
		v.Autoscroll = true
	}
	if v, err := g.SetView(ViewDebug, maxX*5/8, maxY/8, maxX-1, maxY*6/8-1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Debug"
		v.Wrap = true
		v.Editable = false
		v.Autoscroll = true
		log.SetOutput(v)
	}

	if v, err := g.SetView(ViewSend, 0, maxY*6/8, maxX*5/8-1, maxY-1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Input"
		v.Wrap = true
		v.Editable = true
		v.Autoscroll = true
		g.Update(func(gui *gocui.Gui) error {
			if _, err := g.SetCurrentView(ViewSend); err != nil {
				return err
			}
			g.Cursor = true
			return nil
		})
	}

	if v, err := g.SetView(ViewConfig, maxX*5/8, maxY*6/8, maxX-1, maxY-1, 0); err != nil {
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

func ConfigLayouts(g *gocui.Gui) []gocui.Manager {
	enableColor := &ConfigOptionPanel{
		BaseWidget: BaseWidget{
			ViewConfigVisualColorMode,
			func(g *gocui.Gui) (x0, y0, x1, y1 int) {
				maxX, maxY := g.Size()
				xa, ya, xb, yb := maxX*5/8, maxY*6/8, maxX-1, maxY-1
				dx, dy := xb-xa, yb-ya
				return xa + dx/8, ya + dy*2/8, xa + dx*7/8, ya + dy*7/8
			},
		},
		LinkedWidget: LinkedWidget{
			ViewConfigDanmuMode,
			ViewConfigDanmuColor,
		},
		DisplayName: "VisualMode",
		Option: ConfigOption{
			index:        0,
			Options:      []string{"On", "Off"},
			OptionValues: []string{"1", "0"},
		},
		SetConfig: func(value string) {
			Config.VisualColorMode = value == "1"
		},
	}
	enableColor.Option.SetIndexToValue("0")

	danmuColor := &ConfigOptionPanel{
		BaseWidget: BaseWidget{
			ViewConfigDanmuColor,
			func(g *gocui.Gui) (x0, y0, x1, y1 int) {
				maxX, maxY := g.Size()
				xa, ya, xb, yb := maxX*5/8, maxY*6/8, maxX-1, maxY-1
				dx, dy := xb-xa, yb-ya
				return xa + dx/8, ya + dy*2/8, xa + dx*7/8, ya + dy*7/8
			},
		},
		LinkedWidget: LinkedWidget{
			ViewConfigVisualColorMode,
			ViewConfigDanmuMode,
		},
		DisplayName: "DanmuColor",
		Option: ConfigOption{
			index:        0,
			Options:      []string{"白色"},
			OptionValues: []string{"16777215"},
		},
		SetConfig: func(value string) {
			SendFormConfig.Color = value
			go func() {
				resp, err := blivedm.ApiSetDanmuConfig(Client.Account, Client.RoomId,
					"color", "0x"+IntToRGB(cast.ToInt(value)).ToHex())
				if err != nil {
					PrintToDebug(g, "Fail to set color")
					return
				}
				PrintToDebug(g, fmt.Sprintf("set color result - result:%d msg: %s", resp.Code, resp.Message))
				g.Update(func(gui *gocui.Gui) error {
					return nil
				})
			}()
		},
	}
	danmuColor.Option.SetIndexToValue(SendFormConfig.Color)

	danmuMode := &ConfigOptionPanel{
		BaseWidget: BaseWidget{
			ViewConfigDanmuMode,
			func(g *gocui.Gui) (x0, y0, x1, y1 int) {
				maxX, maxY := g.Size()
				xa, ya, xb, yb := maxX*5/8, maxY*6/8, maxX-1, maxY-1
				dx, dy := xb-xa, yb-ya
				return xa + dx/8, ya + dy*2/8, xa + dx*7/8, ya + dy*7/8
			},
		},
		LinkedWidget: LinkedWidget{
			ViewConfigDanmuColor,
			ViewConfigVisualColorMode,
		},
		DisplayName: "DanmuMode",
		Option: ConfigOption{
			index:        0,
			Options:      []string{"滚动"},
			OptionValues: []string{"1"},
		},
		SetConfig: func(value string) {
			SendFormConfig.Mode = cast.ToInt(value)
			go func() {
				resp, err := blivedm.ApiSetDanmuConfig(Client.Account, Client.RoomId,
					"mode", value)
				if err != nil {
					PrintToDebug(g, "Fail to set mode")
					return
				}
				PrintToDebug(g, fmt.Sprintf("set mode result - result:%d msg: %s", resp.Code, resp.Message))
				g.Update(func(gui *gocui.Gui) error {
					return nil
				})
			}()
		},
	}
	danmuMode.Option.SetIndexToValue(cast.ToString(SendFormConfig.Mode))
	go func() {
		_ = <-ClientSet
		defer g.Update(func(gui *gocui.Gui) error {
			return nil
		})
		config, err := blivedm.ApiGetRoomDanmuConfig(Client.Account, Client.RoomId)
		if err != nil || config.Code != 0 {
			PrintToDebug(g, "Load room danmu config fail")
			return
		}
		colors := make([]string, 0)
		colorvals := make([]string, 0)
		for _, group := range config.Data.Group {
			for _, color1 := range group.Color {
				if color1.Status == 1 {
					colors = append(colors, color1.Name)
					colorvals = append(colorvals, color1.Color)
				}
			}
		}
		danmuColor.Option.Options = colors
		danmuColor.Option.OptionValues = colorvals
		danmuColor.Option.SetIndexToValue(SendFormConfig.Color)
		modes := make([]string, 0)
		modevals := make([]string, 0)
		for _, mode := range config.Data.Mode {
			if mode.Status == 1 {
				modes = append(modes, mode.Name)
				modevals = append(modevals, cast.ToString(mode.Mode))
			}
		}
		danmuMode.Option.Options = modes
		danmuMode.Option.OptionValues = modevals
		danmuMode.Option.SetIndexToValue(cast.ToString(SendFormConfig.Mode))

		PrintToDebug(g, "Load room danmu config success")
	}()
	return []gocui.Manager{enableColor, danmuColor, danmuMode}
}
