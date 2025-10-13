package base

import (
	"blivechat/ui"
	"blivechat/ui/styles"
	"blivechat/ui/tabs"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	activeTab int
	tabs      []ui.Tab

	width  int
	height int
	style  *styles.Styles
}

func NewModel() *Model {
	return &Model{
		activeTab: 0,
		tabs: []ui.Tab{
			0: tabs.NewChatTab(),
			1: tabs.NewConfigTab(),
		},
		style: styles.Default,
	}
}

type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg { return tickMsg(t) })
}

func (m *Model) Init() tea.Cmd {
	return m.tabs[m.activeTab].Init()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, 0)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "tab":
			m.activeTab++
			if m.activeTab >= len(m.tabs) {
				m.activeTab = 0
			}
			return m, nil
		}
	case tickMsg:
		cmds = append(cmds, tick())
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	}

	tab := m.tabs[m.activeTab]
	_, cmd := tab.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m *Model) View() string {
	var sb strings.Builder
	sb.WriteString(m.render())
	sb.WriteString("\n")
	sb.WriteString(m.style.TabLine.Render(strings.Repeat("─", m.width)))
	sb.WriteString("\n")
	sb.WriteString(m.tabs[m.activeTab].View())
	return sb.String()
}

func (m *Model) render() string {
	var tabStrs []string
	for tabIdx, tab := range m.tabs {
		name := tab.TabName()
		if tabIdx == m.activeTab {
			tabStrs = append(tabStrs, m.style.TabActive.Render("["+name+"]"))
		} else {
			tabStrs = append(tabStrs, m.style.TabInactive.Render(name))
		}
	}
	return strings.Join(tabStrs, "  ")
}
