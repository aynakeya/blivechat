package chat

import "github.com/charmbracelet/lipgloss"

var Default = DefaultStyles()

func DefaultStyles() Style {
	return Style{
		SuperChat: StyleSc{
			Tag:       lipgloss.NewStyle().Foreground(lipgloss.Color("#FF69B4")).Bold(true),
			Price:     lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Bold(true),
			User:      lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true),
			Msg:       lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")),
			Container: lipgloss.NewStyle().Background(lipgloss.Color("#7C3AED")).Padding(0, 1),
		},
		Gift: StyleGift{
			Tag:   lipgloss.NewStyle().Foreground(lipgloss.Color("#22C55E")).Bold(true),
			User:  lipgloss.NewStyle().Foreground(lipgloss.Color("#E5E7EB")).Bold(true),
			Name:  lipgloss.NewStyle().Foreground(lipgloss.Color("#93C5FD")).Bold(true),
			Count: lipgloss.NewStyle().Foreground(lipgloss.Color("#BBF7D0")).Bold(true),
			Price: lipgloss.NewStyle().Foreground(lipgloss.Color("#F59E0B")).Bold(true),
		},
		Guard: StyleGuard{
			Tag:   lipgloss.NewStyle().Foreground(lipgloss.Color("#38BDF8")).Bold(true),
			User:  lipgloss.NewStyle().Foreground(lipgloss.Color("#E5E7EB")).Bold(true),
			Level: lipgloss.NewStyle().Foreground(lipgloss.Color("#60A5FA")).Bold(true),
			Price: lipgloss.NewStyle().Foreground(lipgloss.Color("#F59E0B")).Bold(true),
		},
		GuardBadge: StyleGuardBadge{
			Base: lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#475569")). // 细灰边
				Padding(0, 1).
				MarginRight(1),
			Captain:  lipgloss.NewStyle().Foreground(lipgloss.Color("#60A5FA")).Bold(true), // 蓝
			Admiral:  lipgloss.NewStyle().Foreground(lipgloss.Color("#A78BFA")).Bold(true), // 紫
			Governor: lipgloss.NewStyle().Foreground(lipgloss.Color("#F59E0B")).Bold(true), // 橙
		},
		Medal: StyleMedal{
			Name:    lipgloss.NewStyle().Bold(true), // 颜色运行时设置
			Level:   lipgloss.NewStyle(),            // 颜色运行时设置
			Bracket: lipgloss.NewStyle().Foreground(lipgloss.Color("#64748B")),
		},
		User: StyleUser{

			Admin: lipgloss.NewStyle().Foreground(lipgloss.Color("#F59E0B")).Bold(true), // 橙
		},
		Danmaku: StyleDanmaku{
			Timestamp: lipgloss.NewStyle().Foreground(lipgloss.Color("#94A3B8")),
			Username:  lipgloss.NewStyle().Bold(true),
			Message:   lipgloss.NewStyle(), // 颜色运行时设置
		},
		SystemMsg: StyleSystemMsg{
			Info:  lipgloss.NewStyle().Foreground(lipgloss.Color("#22C55E")).Bold(true),
			Warn:  lipgloss.NewStyle().Foreground(lipgloss.Color("#F59E0B")).Bold(true),
			Error: lipgloss.NewStyle().Foreground(lipgloss.Color("#EF4444")).Bold(true),
		},
		LiveStatus: StyleLiveStatus{
			Start: lipgloss.NewStyle().Foreground(lipgloss.Color("#10B981")).Bold(true),
			Stop:  lipgloss.NewStyle().Foreground(lipgloss.Color("#F87171")).Bold(true),
		},
		Separator: lipgloss.NewStyle().Foreground(lipgloss.Color("#4B5563")),
		InputBox:  lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("#4B5563")).PaddingLeft(1),
	}
}
