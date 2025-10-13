package styles

import (
	"github.com/charmbracelet/lipgloss"
)

// Styles 统一样式与渲染工具
type Styles struct {
	// Tabs
	TabActive   lipgloss.Style
	TabInactive lipgloss.Style
	TabLine     lipgloss.Style

	ConfigText lipgloss.Style

	// Chat - 基础
	Separator lipgloss.Style
	InputBox  lipgloss.Style

	// Chat - 文本元素
	Timestamp lipgloss.Style
	Username  lipgloss.Style

	// Chat - 勋章
	MedalBracket lipgloss.Style // 勋章方括号
	MedalName    lipgloss.Style // 勋章名（颜色动态）
	MedalLevel   lipgloss.Style // 勋章等级（颜色动态）

	// Chat - 弹幕正文（颜色动态）
	DanmakuText lipgloss.Style

	// 系统提示
	SystemInfo  lipgloss.Style
	SystemWarn  lipgloss.Style
	SystemError lipgloss.Style
	LiveStart   lipgloss.Style
	LiveStop    lipgloss.Style

	// SuperChat
	SCTag       lipgloss.Style // [SC]
	SCPrice     lipgloss.Style // ¥
	SCUser      lipgloss.Style
	SCMsg       lipgloss.Style
	SCContainer lipgloss.Style // 整块背景（可动态换色）

	// Gift
	GiftTag   lipgloss.Style // [礼物]
	GiftUser  lipgloss.Style
	GiftName  lipgloss.Style
	GiftCount lipgloss.Style
	GiftPrice lipgloss.Style

	// Guard
	GuardTag   lipgloss.Style // [大航海]
	GuardUser  lipgloss.Style
	GuardLevel lipgloss.Style
	GuardPrice lipgloss.Style

	// Guard Badge
	GuardBadgeBase lipgloss.Style // 徽标底样式（边框/间距）
	GuardCaptain   lipgloss.Style // Lv1 舰长
	GuardAdmiral   lipgloss.Style // Lv2 提督
	GuardGovernor  lipgloss.Style // Lv3 总督
}
