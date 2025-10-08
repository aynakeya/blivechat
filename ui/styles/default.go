package styles

import "github.com/charmbracelet/lipgloss"

var Default = NewDefaultStyles()

func NewDefaultStyles() *Styles {
	return &Styles{
		TabActive:   lipgloss.NewStyle().Foreground(lipgloss.Color("#00BFFF")).Bold(true),
		TabInactive: lipgloss.NewStyle().Foreground(lipgloss.Color("#8A8FA3")),
		TabLine:     lipgloss.NewStyle().Foreground(lipgloss.Color("#4B5563")),

		Separator: lipgloss.NewStyle().Foreground(lipgloss.Color("#4B5563")),
		InputBox:  lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("#4B5563")).PaddingLeft(1),

		ConfigText: lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#AAAAAA")),

		Timestamp: lipgloss.NewStyle().Foreground(lipgloss.Color("#94A3B8")),
		Username:  lipgloss.NewStyle().Bold(true),

		MedalBracket: lipgloss.NewStyle().Foreground(lipgloss.Color("#64748B")),
		MedalName:    lipgloss.NewStyle().Bold(true), // 颜色运行时设置
		MedalLevel:   lipgloss.NewStyle(),            // 颜色运行时设置

		DanmakuText: lipgloss.NewStyle(), // 颜色运行时设置

		SystemInfo:  lipgloss.NewStyle().Foreground(lipgloss.Color("#22C55E")).Bold(true),
		SystemWarn:  lipgloss.NewStyle().Foreground(lipgloss.Color("#F59E0B")).Bold(true),
		SystemError: lipgloss.NewStyle().Foreground(lipgloss.Color("#EF4444")).Bold(true),
		LiveStart:   lipgloss.NewStyle().Foreground(lipgloss.Color("#10B981")).Bold(true),
		LiveStop:    lipgloss.NewStyle().Foreground(lipgloss.Color("#F87171")).Bold(true),

		SCTag:       lipgloss.NewStyle().Foreground(lipgloss.Color("#FF69B4")).Bold(true),
		SCPrice:     lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Bold(true),
		SCUser:      lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true),
		SCMsg:       lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")),
		SCContainer: lipgloss.NewStyle().Background(lipgloss.Color("#7C3AED")).Padding(0, 1),

		GiftTag:   lipgloss.NewStyle().Foreground(lipgloss.Color("#22C55E")).Bold(true),
		GiftUser:  lipgloss.NewStyle().Foreground(lipgloss.Color("#E5E7EB")).Bold(true),
		GiftName:  lipgloss.NewStyle().Foreground(lipgloss.Color("#93C5FD")).Bold(true),
		GiftCount: lipgloss.NewStyle().Foreground(lipgloss.Color("#BBF7D0")).Bold(true),
		GiftPrice: lipgloss.NewStyle().Foreground(lipgloss.Color("#F59E0B")).Bold(true),

		GuardTag:   lipgloss.NewStyle().Foreground(lipgloss.Color("#38BDF8")).Bold(true),
		GuardUser:  lipgloss.NewStyle().Foreground(lipgloss.Color("#E5E7EB")).Bold(true),
		GuardLevel: lipgloss.NewStyle().Foreground(lipgloss.Color("#60A5FA")).Bold(true),
		GuardPrice: lipgloss.NewStyle().Foreground(lipgloss.Color("#F59E0B")).Bold(true),
	}
}
