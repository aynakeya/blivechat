package internal

import (
	"blivechat/model"
	"encoding/json"
	"github.com/AynaLivePlayer/blivedm-go/api"
	"github.com/AynaLivePlayer/blivedm-go/client"
	"github.com/AynaLivePlayer/blivedm-go/message"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/tidwall/gjson"
	"net/http"
	"strconv"
	"strings"
)

type Backend struct {
	api      api.IApi
	dmClient *client.Client
	roomId   int
	program  *tea.Program
	cred     api.BiliVerify
}

func NewBackend(program *tea.Program, roomId int, cookie string) *Backend {
	iapi := api.NewDefaultClient(cookie)
	var cred api.BiliVerify
	cookie = strings.Trim(strings.TrimSpace(cookie), ";")
	if cookies, err := http.ParseCookie(cookie); err == nil && cookies != nil {
		for _, c := range cookies {
			if c.Name == "SESSDATA" {
				cred.SessData = c.Value
			}
			if c.Name == "bili_jct" {
				cred.Csrf = c.Value
			}
		}
	}
	dmClient := client.NewClientWithApi(roomId, iapi)
	dmClient.OnLiveStart(func(start *message.LiveStart) {
		program.Send((*model.LiveStart)(start))
	})
	dmClient.OnLiveStop(func(stop *message.LiveStop) {
		program.Send((*model.LiveStop)(stop))
	})
	dmClient.OnDanmaku(func(danmu *message.Danmaku) {
		program.Send((*model.Danmaku)(danmu))
	})
	dmClient.OnGift(func(gift *message.Gift) {
		program.Send((*model.Gift)(gift))
	})
	dmClient.OnGuardBuy(func(guard *message.GuardBuy) {
		program.Send((*model.GuardBuy)(guard))
	})
	dmClient.OnSuperChat(func(chat *message.SuperChat) {
		program.Send((*model.SuperChat)(chat))
	})
	dmClient.RegisterCustomEventHandler("INTERACT_WORD", func(s string) {
		var interact = &message.InteractWord{}
		log.Debug(gjson.Get(s, "data").String())
		err := json.Unmarshal(str2bytes(gjson.Get(s, "data").String()), interact)
		if err != nil {
			return
		}
		program.Send((*model.InteractWord)(interact))
	})

	return &Backend{
		api:      iapi,
		dmClient: dmClient,
		roomId:   roomId,
		program:  program,
		cred:     cred,
	}
}

func (b *Backend) Run() error {
	return b.dmClient.Start()
}

func (b *Backend) Stop() error {
	b.dmClient.Stop()
	return nil
}

func (b *Backend) SendDanmuku(request api.DanmakuRequest) error {
	if request.Bubble == "" {
		request.Bubble = "0"
	}
	if request.Color == "" {
		request.Color = "16777215"
	}
	if request.FontSize == "" {
		request.FontSize = "25"
	}
	if request.Mode == "" {
		request.Mode = "1"
	}
	if request.RoomID == "" {
		request.RoomID = strconv.Itoa(b.roomId)
	}
	_, err := api.SendDanmaku(&request, &b.cred)
	return err
}

func (b *Backend) UpdateRoomInfo() (*api.RoomInfo, error) {
	info, err := api.GetRoomInfo(b.roomId)
	if err == nil && info != nil {
		b.program.Send((*model.RoomInfo)(info))
	}
	return info, err
}
