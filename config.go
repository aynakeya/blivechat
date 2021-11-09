package blivechat

import "github.com/aynakeya/blivedm"

var SendFormConfig *blivedm.DanmakuSendForm = &blivedm.DanmakuSendForm{
	Bubble:   0,
	Message:  "",
	Color:    "16777215",
	Fontsize: 25,
	Rnd:      0,
}

var Config = &struct {
	DanmuColor bool
}{
	DanmuColor: true,
}
