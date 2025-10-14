package chat

import (
	"github.com/charmbracelet/lipgloss"
)

type StyleSc struct {
	Tag       lipgloss.Style // [SC] tag
	Price     lipgloss.Style // price
	User      lipgloss.Style // 用户名
	Msg       lipgloss.Style // SC 内容
	Container lipgloss.Style // 整块背景（可动态换色）
}

type StyleGift struct {
	Tag   lipgloss.Style // [礼物] tag
	User  lipgloss.Style // 送礼用户
	Name  lipgloss.Style // 礼物名称
	Count lipgloss.Style // 礼物数量
	Price lipgloss.Style // 礼物单价
}

type StyleMedal struct {
	Name    lipgloss.Style
	Level   lipgloss.Style
	Bracket lipgloss.Style
}

type StyleUser struct {
	Name  lipgloss.Style
	Admin lipgloss.Style
}

type StyleGuardBadge struct {
	Base     lipgloss.Style // 徽标底样式（边框/间距）
	Captain  lipgloss.Style // Lv1 舰长
	Admiral  lipgloss.Style // Lv2 提督
	Governor lipgloss.Style // Lv3 总督
}

type StyleGuard struct {
	Tag   lipgloss.Style // 新上舰 tag
	User  lipgloss.Style
	Level lipgloss.Style
	Price lipgloss.Style
}

type StyleDanmaku struct {
	Timestamp lipgloss.Style
	Username  lipgloss.Style
	Message   lipgloss.Style
}

type StyleSystemMsg struct {
	Info  lipgloss.Style
	Warn  lipgloss.Style
	Error lipgloss.Style
}

type StyleLiveStatus struct {
	Start lipgloss.Style
	Stop  lipgloss.Style
}

// Style chat render styles
type Style struct {
	SuperChat  StyleSc
	Gift       StyleGift
	Guard      StyleGuard
	GuardBadge StyleGuardBadge
	Medal      StyleMedal
	User       StyleUser
	Danmaku    StyleDanmaku
	SystemMsg  StyleSystemMsg
	LiveStatus StyleLiveStatus

	Separator lipgloss.Style
	InputBox  lipgloss.Style
}
