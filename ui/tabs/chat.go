package tabs

import (
	"blivechat/model"
	"blivechat/ui"
	"blivechat/ui/renderer"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/wordwrap"
	"github.com/muesli/reflow/wrap"
)

type ChatTab struct {
	lines    []string
	viewport viewport.Model
	input    textarea.Model
	width    int
	height   int

	renderer ui.ChatRenderer
}

func (m *ChatTab) TabName() string { return "Chat" }

func NewChatTab() *ChatTab {
	ti := textarea.New()
	ti.Placeholder = "输入弹幕 (Enter 发送, Ctrl+C 退出)"
	ti.Focus()
	ti.CharLimit = 200
	ti.SetHeight(3)

	vp := viewport.New(0, 0)
	vp.SetContent("")

	return &ChatTab{
		lines:    make([]string, 0, 128),
		input:    ti,
		viewport: vp,
		renderer: renderer.NewDefaultRenderer(),
	}
}

func (m *ChatTab) Init() tea.Cmd {
	return nil
}

func (m *ChatTab) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			text := strings.TrimSpace(m.input.Value())
			if text != "" {
				m.appendLine(text)
				m.input.Reset()
			}
		default:
			m.input, cmd = m.input.Update(msg)
			cmds = append(cmds, cmd)
		}

	case *model.Danmaku:
		m.appendLine(m.renderer.Danmuku(msg))
	case *model.LiveStart:
		m.appendLine(m.renderer.LiveStart(msg))
	case *model.LiveStop:
		m.appendLine(m.renderer.LiveStop(msg))
	case *model.SystemMsg:
		m.appendLine(m.renderer.SystemMsg(msg))
	case *model.Gift:
		m.appendLine(m.renderer.Gift(msg))
	case *model.GuardBuy:
		m.appendLine(m.renderer.GuardBuy(msg))
	case *model.SuperChat:
		m.appendLine(m.renderer.SuperChat(msg))
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.viewport.Width = m.width - 2
		m.viewport.Height = m.height - 8
		m.input.SetWidth(m.width - 4)
		m.renderer.Styles().InputBox = m.renderer.Styles().InputBox.Width(m.width - 2)
		m.refreshViewport()
	}

	return m, tea.Batch(cmds...)
}

func (m *ChatTab) View() string {
	body := m.viewport.View()
	inputView := m.renderer.Styles().InputBox.Render(m.input.View())

	return strings.Join([]string{
		body,
		m.renderer.Styles().Separator.Render(strings.Repeat("─", m.width)),
		inputView,
	}, "\n")
}

func (m *ChatTab) appendLine(line string) {
	m.lines = append(m.lines, line)
	m.refreshViewport()
	m.viewport.GotoBottom()
}

func (m *ChatTab) refreshViewport() {
	content := strings.Join(m.lines, "\n")

	// Step 1: 词级换行（ANSI-aware）
	ww := wordwrap.NewWriter(m.viewport.Width)
	ww.KeepNewlines = true // 保留你原本拼好的换行
	_, _ = ww.Write([]byte(content))
	_ = ww.Close()
	softWrapped := ww.String()

	// Step 2: 硬折行（ANSI-aware）
	// 解决没有空格的极长行（中文长句、URL、连续表情）仍然可能超宽的问题
	final := wrap.String(softWrapped, m.viewport.Width)

	m.viewport.SetContent(final)
}
