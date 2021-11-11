package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cast"
	"os"
	"os/exec"
)

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
}

type RoomInfoV2Data struct {
	AllSpecialTypes []interface{} `json:"all_special_types"`
	Encrypted       bool          `json:"encrypted"`
	HiddenTill      int           `json:"hidden_till"`
	IsHidden        bool          `json:"is_hidden"`
	IsLocked        bool          `json:"is_locked"`
	IsPortrait      bool          `json:"is_portrait"`
	LiveStatus      int           `json:"live_status"`
	LiveTime        int           `json:"live_time"`
	LockTill        int           `json:"lock_till"`
	PlayurlInfo     interface{}   `json:"playurl_info"`
	PwdVerified     bool          `json:"pwd_verified"`
	RoomID          int           `json:"room_id"`
	RoomShield      int           `json:"room_shield"`
	ShortID         int           `json:"short_id"`
	UID             int           `json:"uid"`
}

type PlayurlData struct {
	CurrentQn          int `json:"current_qn"`
	QualityDescription []struct {
		Qn   int    `json:"qn"`
		Desc string `json:"desc"`
	} `json:"quality_description"`
	Durl []struct {
		Url        string `json:"url"`
		Length     int    `json:"length"`
		Order      int    `json:"order"`
		StreamType int    `json:"stream_type"`
		Ptag       int    `json:"ptag"`
		P2PType    int    `json:"p2p_type"`
	} `json:"durl"`
	IsDashAuto bool `json:"is_dash_auto"`
}

type RoomInfoV2Response struct {
	BaseResponse
	Data RoomInfoV2Data `json:"data"`
}

type PlayurlDataResponse struct {
	BaseResponse
	Data PlayurlData `json:"data"`
}

func GetRoomId(shortId string) int {
	get, err := resty.New().R().
		SetQueryParam("room_id", shortId).
		Get("https://api.live.bilibili.com/xlive/web-room/v2/index/getRoomPlayInfo")
	if err != nil {
		return -1
	}
	var resp RoomInfoV2Response
	if err := json.Unmarshal(get.Body(), &resp); err != nil {
		return -1
	}
	fmt.Println("Live status: ", resp.Data.LiveStatus)
	if resp.Data.LiveStatus != 1 {
		return -1
	}
	return resp.Data.RoomID
}

func GetPlayUrl(roomId int) string {
	get, err := resty.New().R().
		// fucking stupid api
		SetQueryParams(map[string]string{
			"cid":           cast.ToString(roomId),
			"qn":            "0",
			"platform":      "h5",
			"https_url_req": "1",
			"ptype":         "16",
		}).
		Get("https://api.live.bilibili.com/xlive/web-room/v1/playUrl/playUrl")
	if err != nil {
		return ""
	}
	var resp PlayurlDataResponse
	if err := json.Unmarshal(get.Body(), &resp); err != nil {
		return ""
	}
	return resp.Data.Durl[0].Url
}

func main() {
	defer func() {
		fmt.Println("Press entet to quit")
		var a string
		if _, err := fmt.Scanln(&a); err != nil {
			return
		}
	}()
	fmt.Println("Enter room number> ")
	var shortId string
	if _, err := fmt.Scanln(&shortId); err != nil {
		fmt.Println("fail to read room id")
	}
	fmt.Printf("Try starting %s\n", shortId)
	roomId := GetRoomId(shortId)
	if roomId == -1 {
		fmt.Println("fail to get real room id, please check if you enter the wrong room id")
	}
	fmt.Println("Real room id", roomId)
	durl := GetPlayUrl(roomId)
	if durl == "" {
		fmt.Println("fail to get url")
	}
	fmt.Printf("Get url success %s\n", durl)
	cmd := exec.Command("mpv", "--demuxer-max-back-bytes=16777216", "--demuxer-max-bytes=16777216", "--no-video", durl)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
