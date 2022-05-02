package blivechat

import (
	"blivechat/util"
	"fmt"
	"github.com/awesome-gocui/gocui"
	"github.com/aynakeya/blivedm"
	"time"
)

var Client *blivedm.BLiveWsClient
var ClientSet = make(chan int)

func danmuMsgRecv(context *blivedm.Context) {
	msg, _ := context.ToDanmakuMessage()
	util.PrintDanmu(MainGui, ViewDanmu, Config.VisualColorMode, Config.ShowMedal, msg)
	MainGui.Update(func(gui *gocui.Gui) error {
		return nil
	})
}

func sendDanmuMsg(msg string) (blivedm.DanmakuSendResponse, error) {
	return Client.SendMessage(blivedm.DanmakuSendForm{
		Bubble:   SendFormConfig.Bubble,
		Message:  msg,
		Mode:     SendFormConfig.Mode,
		Color:    SendFormConfig.Color,
		Fontsize: SendFormConfig.Fontsize,
		Rnd:      int(time.Now().Unix()),
	})
}

func initializeClient() {
	util.ViewPrintWithTime(MainGui, ViewDebug, fmt.Sprintf("try get room info"))
	util.ViewPrintWithTime(MainGui, ViewDebug, fmt.Sprintf("GetRoomInfo: %t", Client.GetRoomInfo()))
	MainGui.Update(func(gui *gocui.Gui) error { return nil })
	util.ViewPrintWithTime(MainGui, ViewDebug, fmt.Sprintf("try danmu room info"))
	util.ViewPrintWithTime(MainGui, ViewDebug, fmt.Sprintf("GetDanmuInfo: %t", Client.GetDanmuInfo()))
	MainGui.Update(func(gui *gocui.Gui) error { return nil })
	util.ViewPrintWithTime(MainGui, ViewDebug, fmt.Sprintf("try connect to danmu server"))
	util.ViewPrintWithTime(MainGui, ViewDebug, fmt.Sprintf("ConnectToDanmuServer: %t", Client.ConnectDanmuServer()))
	MainGui.Update(func(gui *gocui.Gui) error { return nil })
	roomV, err := MainGui.View(ViewRoom)
	if err != nil {
		util.ViewPrintWithTime(MainGui, ViewDebug, err)
		return
	}
	var upname string
	if info, err := blivedm.ApiGetUpInfo(Client.RoomInfo.Uid); err != nil {
		upname = "Unknown"
	} else {
		upname = info.Data.Name
	}
	fmt.Fprintf(roomV,
		fmt.Sprintf("%s > %s | %s (%d) - Live: %t | %s (%d) | login as uid=%d",
			Client.RoomInfo.ParentAreaName, Client.RoomInfo.AreaName,
			Client.RoomInfo.Title, Client.RoomInfo.RoomId, Client.RoomInfo.LiveStatus == 1,
			upname, Client.RoomInfo.Uid,
			Client.Account.UID))
	MainGui.Update(func(gui *gocui.Gui) error { return nil })
	ClientSet <- 1
}

func SetupDanmuClient(cl *blivedm.BLiveWsClient) {
	Client = cl
	Client.RegHandler(blivedm.CmdDanmaku, danmuMsgRecv)

	MainGui.Update(func(gui *gocui.Gui) error {
		go initializeClient()
		return nil
	})
}
