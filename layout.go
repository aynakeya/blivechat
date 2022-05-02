package blivechat

import (
	"blivechat/util"
	"blivechat/widget"
	"fmt"
	"github.com/awesome-gocui/gocui"
	"github.com/aynakeya/blivedm"
	"github.com/spf13/cast"
)

func createLayoutManager() {
	createMainLayout()
	createConfigLayout()
}

var roomInfoView, danmuView, debugView, sendView, configView *widget.CommonPanel

func createMainLayout() {
	roomInfoView = &widget.CommonPanel{
		BaseWidget: widget.BaseWidget{
			ViewName:       ViewRoom,
			ParentViewName: "",
			DisplayName:    "RoomInfo",
			GetSize: func(x, y int) (x0, y0, x1, y1 int) {
				return 0, 0, x - 1, y/8 - 1
			},
		},
		LinkedWidget: widget.LinkedWidget{
			NextViewName: ViewDanmu,
		},
		SwitchKey: widget.KeyCombo{
			Key:      gocui.KeyTab,
			Modifier: gocui.ModNone,
		},
	}

	danmuView = &widget.CommonPanel{
		BaseWidget: widget.BaseWidget{
			ViewName:       ViewDanmu,
			ParentViewName: "",
			DisplayName:    "Danmu",
			GetSize: func(x, y int) (x0, y0, x1, y1 int) {
				return 0, y / 8, x*5/8 - 1, y*6/8 - 1
			},
		},
		LinkedWidget: widget.LinkedWidget{
			NextViewName: ViewDebug,
		},
		SwitchKey: widget.KeyCombo{
			Key:      gocui.KeyTab,
			Modifier: gocui.ModNone,
		},
	}

	debugView = &widget.CommonPanel{
		BaseWidget: widget.BaseWidget{
			ViewName:       ViewDebug,
			ParentViewName: "",
			DisplayName:    "Debug",
			GetSize: func(x, y int) (x0, y0, x1, y1 int) {
				return x * 5 / 8, y / 8, x - 1, y*6/8 - 1
			},
		},
		LinkedWidget: widget.LinkedWidget{
			NextViewName: ViewSend,
		},
		SwitchKey: widget.KeyCombo{
			Key:      gocui.KeyTab,
			Modifier: gocui.ModNone,
		},
	}

	sendView = &widget.CommonPanel{
		BaseWidget: widget.BaseWidget{
			ViewName:       ViewSend,
			ParentViewName: "",
			DisplayName:    "Input",
			GetSize: func(x, y int) (x0, y0, x1, y1 int) {
				return 0, y * 6 / 8, x*5/8 - 1, y - 1
			},
		},
		LinkedWidget: widget.LinkedWidget{
			NextViewName: ViewConfig,
		},
		SwitchKey: widget.KeyCombo{
			Key:      gocui.KeyTab,
			Modifier: gocui.ModNone,
		},
	}

	configView = &widget.CommonPanel{
		BaseWidget: widget.BaseWidget{
			ViewName:       ViewConfig,
			ParentViewName: "",
			DisplayName:    "Config",
			GetSize: func(x, y int) (x0, y0, x1, y1 int) {
				return x * 5 / 8, y * 6 / 8, x - 1, y - 1
			},
		},
		LinkedWidget: widget.LinkedWidget{
			NextViewName: ViewRoom,
		},
		SwitchKey: widget.KeyCombo{
			Key:      gocui.KeyTab,
			Modifier: gocui.ModNone,
		},
	}

	MainManager = append(MainManager, roomInfoView, danmuView, debugView, sendView, configView)
}

func createConfigLayout() {
	sizeFunc := func(x, y int) (x0, y0, x1, y1 int) {
		xa, ya, xb, yb := x*5/8, y*6/8, x-1, y-1
		dx, dy := xb-xa, yb-ya
		return xa + dx/8, ya + dy*2/8, xa + dx*7/8, ya + dy*7/8
	}

	enableColor := &widget.ConfigOptionPanel{
		BaseWidget: widget.BaseWidget{
			ViewName:       ViewConfigVisualColorMode,
			ParentViewName: ViewConfig,
			DisplayName:    "VisualMode",
			GetSize:        sizeFunc,
		},
		DoubleLinkedWidget: widget.DoubleLinkedWidget{
			PrevViewName: ViewConfigDanmuMode,
			NextViewName: ViewConfigShowMedal,
		},
		Option: widget.NewConfigOption(map[string]interface{}{
			"On":  true,
			"Off": false,
		}),
		SetConfig: func(value interface{}, origin *widget.ConfigOptionPanel) {
			Config.VisualColorMode = value.(bool)
		},
	}
	enableColor.Option.SetByValue(Config.VisualColorMode)

	showMedal := &widget.ConfigOptionPanel{
		BaseWidget: widget.BaseWidget{
			ViewName:       ViewConfigShowMedal,
			ParentViewName: ViewConfig,
			DisplayName:    "ShowMedal",
			GetSize:        sizeFunc,
		},
		DoubleLinkedWidget: widget.DoubleLinkedWidget{
			PrevViewName: ViewConfigVisualColorMode,
			NextViewName: ViewConfigShowDebug,
		},
		Option: widget.NewConfigOption(map[string]interface{}{
			"On":  true,
			"Off": false,
		}),
		SetConfig: func(value interface{}, origin *widget.ConfigOptionPanel) {
			Config.ShowMedal = value.(bool)
		},
	}

	showMedal.Option.SetByValue(Config.ShowMedal)

	showDebug := &widget.ConfigOptionPanel{
		BaseWidget: widget.BaseWidget{
			ViewName:       ViewConfigShowDebug,
			ParentViewName: ViewConfig,
			DisplayName:    "ShowDebug",
			GetSize:        sizeFunc,
		},
		DoubleLinkedWidget: widget.DoubleLinkedWidget{
			PrevViewName: ViewConfigShowMedal,
			NextViewName: ViewConfigDanmuColor,
		},
		Option: widget.NewConfigOption(map[string]interface{}{
			"On":  true,
			"Off": false,
		}),
		SetConfig: func(value interface{}, origin *widget.ConfigOptionPanel) {
			Config.ShowDebug = value.(bool)
			if Config.ShowDebug {
				danmuView.NextViewName = ViewDebug
				danmuView.GetSize = func(x, y int) (x0, y0, x1, y1 int) {
					return 0, y / 8, x*5/8 - 1, y*6/8 - 1
				}
			} else {
				danmuView.NextViewName = ViewSend
				danmuView.GetSize = func(x, y int) (x0, y0, x1, y1 int) {
					return 0, y / 8, x - 1, y*6/8 - 1

				}
			}
			MainGui.Update(func(gui *gocui.Gui) error {
				gui.SetViewOnTop(ViewDanmu)
				return nil
			})
		},
	}

	showDebug.Option.SetByValue(Config.ShowDebug)

	danmuColor := &widget.ConfigOptionPanel{
		BaseWidget: widget.BaseWidget{
			ViewName:       ViewConfigDanmuColor,
			ParentViewName: ViewConfig,
			DisplayName:    "DanmuColor",
			GetSize:        sizeFunc,
		},
		DoubleLinkedWidget: widget.DoubleLinkedWidget{
			PrevViewName: ViewConfigShowMedal,
			NextViewName: ViewConfigDanmuMode,
		},
		Option: widget.NewConfigOption(map[string]interface{}{
			"白色": "16777215",
		}),
		SetConfig: func(value interface{}, origin *widget.ConfigOptionPanel) {
			SendFormConfig.Color = value.(string)
			go func() {
				resp, err := blivedm.ApiSetDanmuConfig(Client.Account, Client.RoomId,
					"color", "0x"+util.IntToRGB(cast.ToInt(value)).ToHex())
				if err != nil {
					util.ViewPrintWithTime(MainGui, ViewDebug, "Fail to set color")
					return
				}
				util.ViewPrintWithTime(MainGui, ViewDebug, fmt.Sprintf("set color result - result:%d msg: %s", resp.Code, resp.Message))
				MainGui.Update(func(gui *gocui.Gui) error {
					return nil
				})
			}()
		},
	}
	danmuColor.Option.SetByValue(SendFormConfig.Color)

	danmuMode := &widget.ConfigOptionPanel{
		BaseWidget: widget.BaseWidget{
			ViewName:       ViewConfigDanmuMode,
			ParentViewName: ViewConfig,
			DisplayName:    "DanmuMode",
			GetSize:        sizeFunc,
		},
		DoubleLinkedWidget: widget.DoubleLinkedWidget{
			PrevViewName: ViewConfigDanmuColor,
			NextViewName: ViewConfigVisualColorMode,
		},
		Option: widget.NewConfigOption(map[string]interface{}{
			"滚动": 1,
		}),
		SetConfig: func(value interface{}, origin *widget.ConfigOptionPanel) {
			SendFormConfig.Mode = value.(int)
			go func() {
				resp, err := blivedm.ApiSetDanmuConfig(Client.Account, Client.RoomId,
					"mode", cast.ToString(value))
				if err != nil {
					util.ViewPrintWithTime(MainGui, ViewDebug, "Fail to set mode")
					return
				}
				util.ViewPrintWithTime(MainGui, ViewDebug, fmt.Sprintf("set mode result - result:%d msg: %s", resp.Code, resp.Message))
				MainGui.Update(func(gui *gocui.Gui) error {
					return nil
				})
			}()
		},
	}
	danmuMode.Option.SetByValue(SendFormConfig.Mode)

	go func() {
		_ = <-ClientSet
		defer MainGui.Update(func(gui *gocui.Gui) error {
			return nil
		})
		config, err := blivedm.ApiGetRoomDanmuConfig(Client.Account, Client.RoomId)
		if err != nil || config.Code != 0 {
			util.ViewPrintWithTime(MainGui, ViewDebug, "Load room danmu config fail")
			return
		}
		colors := make([]string, 0)
		colorvals := make([]interface{}, 0)
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
		danmuColor.Option.SetByValue(SendFormConfig.Color)
		modes := make([]string, 0)
		modevals := make([]interface{}, 0)
		for _, mode := range config.Data.Mode {
			if mode.Status == 1 {
				modes = append(modes, mode.Name)
				modevals = append(modevals, mode.Mode)
			}
		}
		danmuMode.Option.Options = modes
		danmuMode.Option.OptionValues = modevals
		danmuMode.Option.SetByValue(SendFormConfig.Mode)

		util.ViewPrintWithTime(MainGui, ViewDebug, "Load room danmu config success")
	}()

	MainManager = append(MainManager, enableColor, showMedal, showDebug, danmuColor, danmuMode)
	return
}
