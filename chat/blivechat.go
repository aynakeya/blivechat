// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"blivechat"
	"fmt"
	"github.com/awesome-gocui/gocui"
	"github.com/aynakeya/blivedm"
	"github.com/spf13/cast"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"strings"
	"time"
)

const DEFAULT_CONFIG_NAME = "blivechat.ini"

type ConfigFile struct {
	Uid         int
	SessionData string
	BilibiliJCT string

	VisualColorMode bool
}

func saveToConfig(cfg ConfigFile) {
	cfgFile := ini.Empty()
	cfgFile.Section("blivechat").Key("Uid").SetValue(cast.ToString(cfg.Uid))
	cfgFile.Section("blivechat").Key("VisualColorMode").SetValue(cast.ToString(cfg.VisualColorMode))
	cfgFile.Section("blivechat").Key("SessionData").SetValue(cast.ToString(cfg.SessionData))
	cfgFile.Section("blivechat").Key("BilibiliJCT").SetValue(cast.ToString(cfg.BilibiliJCT))
	err := cfgFile.SaveTo(DEFAULT_CONFIG_NAME)
	if err != nil {
		return
	}
}

func getOrDefault(input string, defaultValue string) string {
	if input == "" {
		return defaultValue
	}
	return input
}

func startFromArgs(roomId int, args []string) *gocui.Gui {
	uid := 0
	sessdata := ""
	bilijct := ""
	if len(args) >= 4 {
		uid = cast.ToInt(args[1])
		sessdata = cast.ToString(args[2])
		bilijct = cast.ToString(args[3])
	} else {
		g := startFromConfig(roomId, DEFAULT_CONFIG_NAME)
		if g != nil {
			return g
		}
	}
	g := blivechat.CreateGUI()
	cl := blivedm.BLiveWsClient{ShortId: roomId,
		Account: blivedm.DanmuAccount{
			UID:         uid,
			SessionData: sessdata,
			BilibiliJCT: bilijct,
		},
		HearbeatInterval: 25 * time.Second}
	blivechat.SetupDanmuClient(g, &cl)
	saveToConfig(ConfigFile{
		Uid:             uid,
		SessionData:     sessdata,
		BilibiliJCT:     bilijct,
		VisualColorMode: false,
	})
	return g
}

func startFromConfig(roomId int, filename string) *gocui.Gui {
	cfg := ConfigFile{
		Uid:             0,
		SessionData:     "",
		BilibiliJCT:     "",
		VisualColorMode: false,
	}
	cfgFile, err := ini.Load(filename)
	if err != nil {
		return nil
	}

	cfg.Uid = cast.ToInt(getOrDefault(cfgFile.Section("blivechat").
		Key("Uid").Value(), "0"))
	cfg.VisualColorMode = cast.ToBool(
		getOrDefault(cfgFile.Section("blivechat").
			Key("VisualColorMode").Value(), "false"))
	cfg.SessionData = getOrDefault(cfgFile.Section("blivechat").
		Key("SessionData").Value(), "")
	cfg.BilibiliJCT = getOrDefault(cfgFile.Section("blivechat").
		Key("BilibiliJCT").Value(), "")
	g := blivechat.CreateGUI()
	cl := blivedm.BLiveWsClient{ShortId: roomId,
		Account: blivedm.DanmuAccount{
			UID:         cfg.Uid,
			SessionData: cfg.SessionData,
			BilibiliJCT: cfg.BilibiliJCT,
		},
		HearbeatInterval: 25 * time.Second}
	blivechat.SetupDanmuClient(g, &cl)
	return g
}

func main() {
	var g *gocui.Gui
	args := os.Args[1:]
	roomId := 0

	if len(args) < 1 {
		defer func() {
			fmt.Print("Press entet to quit")
			var a string
			if _, err := fmt.Scanln(&a); err != nil {
				return
			}
		}()
		fmt.Println("Usage blivechat <room_id> Optional[<uid> <sessdata> <bilijct>]\n" +
			"	blivechat <room_id> --c=<config file name>")
		fmt.Println("Please enter room id > ")
		_, err := fmt.Scanln(&roomId)
		if err != nil {
			return
		}
	} else {
		roomId = cast.ToInt(args[0])
	}
	if roomId == 0 {
		fmt.Println("Room id is not proper")
		return
	}
	useConfig := false
	configName := "blivechat.ini"
	for _, arg := range args {
		if strings.HasPrefix(arg, "--c") {
			useConfig = true
			configName = strings.ReplaceAll(arg, "--c=", "")
			break
		}
	}
	if useConfig {
		g = startFromConfig(roomId, configName)
	} else {
		g = startFromArgs(roomId, args)
	}
	if g == nil {
		return
	}
	defer func() {
		g.Cursor = false
		g.Close()
	}()
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
	return
}
