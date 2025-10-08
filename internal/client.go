package internal

import (
	"blivechat/model"
	"github.com/AynaLivePlayer/blivedm-go/api"
	"github.com/AynaLivePlayer/blivedm-go/client"
	"github.com/AynaLivePlayer/blivedm-go/message"
	tea "github.com/charmbracelet/bubbletea"
)

type Backend struct {
	api      api.IApi
	dmClient *client.Client
	roomId   int
	program  *tea.Program
}

func NewBackend(program *tea.Program, roomId int, cookie string) *Backend {
	iapi := api.NewDefaultClient(cookie)
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

	return &Backend{
		api:      iapi,
		dmClient: dmClient,
		roomId:   roomId,
		program:  program,
	}
}

func (b *Backend) Run() error {
	return b.dmClient.Start()
}

func (b *Backend) Stop() error {
	b.dmClient.Stop()
	return nil
}
