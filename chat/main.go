// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"blivechat"
	"fmt"
	"github.com/aynakeya/blivedm"
	"github.com/aynakeya/gocui"
	"github.com/spf13/cast"
	"log"
	"os"
	"time"
)

func main() {
	args :=  os.Args[1:]

	if len(args) <1 {
		fmt.Println("Usage blivechat <room_id> Optional[<uid> <sessdata> <bilijct>]")
		return
	}
	roomId := cast.ToInt(args[0])
	if roomId == 0 {
		fmt.Println("Room id is not proper")
		return
	}
	uid :=0
	sessdata := ""
	bilijct := ""
	if len(args) >= 4 {
		uid = cast.ToInt(args[1])
		sessdata = cast.ToString(args[2])
		bilijct = cast.ToString(args[3])
	}
	g := blivechat.CreateGUI()
	defer g.Close()
	cl := blivedm.BLiveWsClient{ShortId: roomId,
		Account: blivedm.DanmuAccount{
			UID:         uid,
			SessionData: sessdata,
			BilibiliJCT: bilijct,
		},
		HearbeatInterval: 25 * time.Second}
	blivechat.SetKeybindings(g)
	blivechat.SetupDanmuClient(g,&cl)
	go func() {
		time.Sleep(1 * time.Second)
		err := blivechat.InitEditView(g)
		if err != nil {
			log.Panicln(err)
		}
	}()
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
