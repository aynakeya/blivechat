package tabs

import (
	"blivechat/ui/styles"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type ConfigTab struct {
	width  int
	height int
	style  *styles.Styles
}

func (m *ConfigTab) TabName() string {
	return "Config"
}

func NewConfigTab() *ConfigTab {
	return &ConfigTab{style: styles.Default}
}

func (m *ConfigTab) Init() tea.Cmd { return nil }

func (m *ConfigTab) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	}
	return m, nil
}

func (m *ConfigTab) View() string {
	content := `
配置页

- 房间号: 123456
- 弹幕颜色: 默认
- 弹幕过滤: 无
- 自动重连: 开启
`
	return m.style.ConfigText.Render(strings.TrimSpace(content))
}
