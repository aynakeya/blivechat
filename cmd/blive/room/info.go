package room

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"text/tabwriter"
)

var infoShowCode bool

func init() {
	infoCmd.Flags().BoolVarP(&infoShowCode, "with-code", "", false, "show code for current room")
}

type roomInfo struct {
	areaSubName string
	areaName    string
	title       string
	uname       string
	uid         int64
	roomId      int64
}

func getInfo() (roomInfo, error) {
	get, err := client.R().Get("https://api.live.bilibili.com/xlive/app-blink/v1/room/GetInfo?platform=pc")
	if err != nil {
		return roomInfo{}, err
	}
	jdata := gjson.ParseBytes(get.Bytes())
	if jdata.Get("message").String() != "0" {
		return roomInfo{}, errors.New("get info failed: " + jdata.Get("message").String())
	}
	info := roomInfo{
		areaSubName: jdata.Get("data.area_v2_name").String(),
		areaName:    jdata.Get("data.parent_name").String(),
		title:       jdata.Get("data.title").String(),
		uname:       jdata.Get("data.uname").String(),
		uid:         jdata.Get("data.uid").Int(),
		roomId:      jdata.Get("data.room_id").Int(),
	}
	return info, nil
}

func getLiveToken() (string, error) {
	get, err := client.R().SetFormData(
		map[string]string{
			"action":     "1",
			"csrf_token": csrfToken,
			"csrf":       csrfToken,
		}).Post("https://api.live.bilibili.com/xlive/open-platform/v1/common/operationOnBroadcastCode")
	if err != nil {
		return "", err
	}
	data := gjson.GetBytes(get.Bytes(), "message").String()
	if data != "0" {
		return "", errors.New("failed: " + data)
	}
	return gjson.GetBytes(get.Bytes(), "data.code").String(), nil
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "get info for current room",
	Run: func(cmd *cobra.Command, args []string) {
		info, err := getInfo()
		if err != nil {
			log.Errorf("failed to set title: %v", err)
			return
		}
		var liveToken string
		if infoShowCode {
			liveToken, _ = getLiveToken()
		}
		log.Infof("-- start room info --")
		var buf bytes.Buffer
		tabw := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
		bold := lipgloss.NewStyle().Bold(true)
		_, _ = tabw.Write([]byte(fmt.Sprintf("%s: \t%s(%d)\n", bold.Render("User"), info.uname, info.uid)))
		_, _ = tabw.Write([]byte(fmt.Sprintf("%s: \t%s\n", bold.Render("Title"), info.title)))
		_, _ = tabw.Write([]byte(fmt.Sprintf("%s: \t%s - %s", bold.Render("Area"), info.areaName, info.areaSubName)))
		if liveToken != "" {
			_, _ = tabw.Write([]byte(fmt.Sprintf("\n%s: \t%s", bold.Render("LiveCode"), liveToken)))
		}
		_ = tabw.Flush()
		fmt.Println(buf.String())
		log.Infof("-- end room info --")
		return
	},
}
