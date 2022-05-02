package blivechat

import (
	"github.com/aynakeya/blivedm"
)

var SendFormConfig *blivedm.DanmakuSendForm = &blivedm.DanmakuSendForm{
	Bubble:   0,
	Message:  "",
	Color:    "16777215",
	Mode:     1,
	Fontsize: 25,
	Rnd:      0,
}

type GUIConfig struct {
	VisualColorMode bool
	ShowMedal       bool
	ShowDebug       bool
}

var Config = GUIConfig{
	VisualColorMode: false,
	ShowMedal:       true,
	ShowDebug:       false,
}
