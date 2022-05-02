// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"blivechat"
	"fmt"
	"github.com/aynakeya/blivedm"
	"github.com/spf13/cast"
	"gopkg.in/ini.v1"
	"os"
	"strings"
	"time"
)

const DEFAULT_CONFIG_NAME = "blivechat.ini"

type ConfigFile struct {
	*blivechat.GUIConfig
	Uid         int
	SessionData string
	BilibiliJCT string
}

var BlivechatConfig = ConfigFile{
	GUIConfig: &blivechat.Config,
}

func SaveConfig(filename string) error {
	cfgFile := ini.Empty()
	section := cfgFile.Section("blivechat")
	section.Key("Uid").SetValue(cast.ToString(BlivechatConfig.Uid))
	section.Key("SessionData").SetValue(cast.ToString(BlivechatConfig.SessionData))
	section.Key("BilibiliJCT").SetValue(cast.ToString(BlivechatConfig.BilibiliJCT))

	section.Key("VisualColorMode").SetValue(cast.ToString(BlivechatConfig.VisualColorMode))
	section.Key("ShowDebug").SetValue(cast.ToString(BlivechatConfig.ShowDebug))
	section.Key("ShowMedal").SetValue(cast.ToString(BlivechatConfig.ShowMedal))
	return cfgFile.SaveTo(filename)
}

func LoadConfig(filename string) error {
	cfgFile, err := ini.Load(filename)
	if err != nil {
		return err
	}
	section := cfgFile.Section("blivechat")
	BlivechatConfig.Uid = section.Key("Uid").MustInt(0)
	BlivechatConfig.SessionData = section.Key("SessionData").Value()
	BlivechatConfig.BilibiliJCT = section.Key("BilibiliJCT").Value()

	BlivechatConfig.VisualColorMode = section.Key("VisualColorMode").MustBool(false)
	BlivechatConfig.ShowDebug = section.Key("ShowDebug").MustBool(true)
	BlivechatConfig.ShowMedal = section.Key("ShowMedal").MustBool(true)
	return nil
}

func startFromArgs(roomId int, args []string) error {
	if len(args) < 4 {
		if err := startFromConfig(roomId, DEFAULT_CONFIG_NAME); err == nil {
			return nil
		}
	} else {
		BlivechatConfig.Uid = cast.ToInt(args[1])
		BlivechatConfig.SessionData = args[2]
		BlivechatConfig.BilibiliJCT = args[3]
	}
	err := blivechat.CreateGUI()
	if err != nil {
		return err
	}
	cl := blivedm.BLiveWsClient{ShortId: roomId,
		Account: blivedm.DanmuAccount{
			UID:         BlivechatConfig.Uid,
			SessionData: BlivechatConfig.SessionData,
			BilibiliJCT: BlivechatConfig.BilibiliJCT,
		},
		HearbeatInterval: 25 * time.Second}
	blivechat.SetupDanmuClient(&cl)
	SaveConfig(DEFAULT_CONFIG_NAME)
	return nil
}

func startFromConfig(roomId int, filename string) error {
	if err := LoadConfig(filename); err != nil {
		return err
	}
	if err := blivechat.CreateGUI(); err != nil {
		return err
	}
	cl := blivedm.BLiveWsClient{ShortId: roomId,
		Account: blivedm.DanmuAccount{
			UID:         BlivechatConfig.Uid,
			SessionData: BlivechatConfig.SessionData,
			BilibiliJCT: BlivechatConfig.BilibiliJCT,
		},
		HearbeatInterval: 25 * time.Second}
	blivechat.SetupDanmuClient(&cl)
	return nil
}

func main() {
	args := os.Args[1:]
	roomId := 0

	if len(args) < 1 {
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
	configName := DEFAULT_CONFIG_NAME
	for _, arg := range args {
		if strings.HasPrefix(arg, "--c") {
			useConfig = true
			configName = strings.ReplaceAll(arg, "--c=", "")
			break
		}
	}
	var err error
	if useConfig {
		err = startFromConfig(roomId, configName)
	} else {
		err = startFromArgs(roomId, args)
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		SaveConfig(configName)
		blivechat.MainGui.Cursor = false
		blivechat.MainGui.Close()
	}()
	if err := blivechat.MainGui.MainLoop(); err != nil {
	}
	return
}
