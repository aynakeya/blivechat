package tab

import "github.com/charmbracelet/lipgloss"

//TabActive:   lipgloss.NewStyle().Foreground(lipgloss.Color("#00BFFF")).Bold(true),
//TabInactive: lipgloss.NewStyle().Foreground(lipgloss.Color("#8A8FA3")),
//TabLine:     lipgloss.NewStyle().Foreground(lipgloss.Color("#4B5563")),

type Style struct {
	TabActive   lipgloss.Style
	TabInactive lipgloss.Style
	TabLine     lipgloss.Style
}

func DefaultStyle() Style {
	return Style{
		TabActive:   lipgloss.NewStyle().Foreground(lipgloss.Color("#00BFFF")).Bold(true),
		TabInactive: lipgloss.NewStyle().Foreground(lipgloss.Color("#8A8FA3")),
		TabLine:     lipgloss.NewStyle().Foreground(lipgloss.Color("#4B5563")),
	}
}
