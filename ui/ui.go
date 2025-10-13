package ui

import (
	"blivechat/model"
	"blivechat/ui/styles"
	tea "github.com/charmbracelet/bubbletea"
)

type Tab interface {
	tea.Model
	TabName() string
}

type ChatRenderer interface {
	UseStyle(style *styles.Styles)
	Styles() *styles.Styles // return current styles, modifiable
	Danmuku(msg *model.Danmaku) string
	LiveStart(msg *model.LiveStart) string
	LiveStop(msg *model.LiveStop) string
	Gift(msg *model.Gift) string
	GuardBuy(msg *model.GuardBuy) string
	SuperChat(msg *model.SuperChat) string
	SystemMsg(msg *model.SystemMsg) string
	InteractWord(msg *model.InteractWord) string
}
