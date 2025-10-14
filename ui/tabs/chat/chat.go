package chat

import (
	"blivechat/model"
	"blivechat/ui/got"
	"github.com/AynaLivePlayer/blivedm-go/api"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/muesli/reflow/wordwrap"
	"github.com/muesli/reflow/wrap"
	"strings"
)

type ChatTab struct {
	lines    []string
	viewport viewport.Model
	input    textarea.Model
	width    int
	height   int

	infoLine string
	renderer ChatRenderer
	styles   Style
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

	tab := &ChatTab{
		lines:    make([]string, 0, 128),
		input:    ti,
		viewport: vp,
		renderer: &DefaultRenderer{},
		styles:   DefaultStyles(),
	}
	tab.renderer.UseStyle(&tab.styles)
	return tab
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
				go func() {
					err := got.Backend.SendDanmuku(api.DanmakuRequest{Msg: text})
					if err != nil {
						log.Errorf("send danmu failed: %v", err)
					}
				}()
				m.input.Reset()
			}
		default:
			m.input, cmd = m.input.Update(msg)
			cmds = append(cmds, cmd)
		}
	case *model.RoomInfo:
		m.infoLine = m.renderer.RoomTitle(msg)
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
	case *model.InteractWord:
		m.appendLine(m.renderer.InteractWord(msg))
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.viewport.Width = m.width - 2
		m.viewport.Height = m.height - 10
		m.input.SetWidth(m.width - 4)
		m.renderer.Styles().InputBox = m.renderer.Styles().InputBox.Width(m.width - 2)
		m.refreshViewport()
	}

	return m, tea.Batch(cmds...)
}

func (m *ChatTab) View() string {
	body := m.viewport.View()
	inputView := m.renderer.Styles().InputBox.Render(m.input.View())

	line := "no room info"
	if m.infoLine != "" {
		line = m.infoLine
	}

	return strings.Join([]string{
		line,
		m.renderer.Styles().Separator.Render(strings.Repeat("─", m.width)),
		body,
		m.renderer.Styles().Separator.Render(strings.Repeat("─", m.width)),
		inputView,
	}, "\n")
}

func (m *ChatTab) appendLine(line string) {
	if line == "" {
		return
	}
	m.lines = append(m.lines, line)
	m.refreshViewport()
	m.viewport.GotoBottom()
}

func (m *ChatTab) refreshViewport() {
	content := strings.Join(m.lines, "\n")

	ww := wordwrap.NewWriter(m.viewport.Width)
	ww.KeepNewlines = true // 保留你原本拼好的换行
	_, _ = ww.Write([]byte(content))
	_ = ww.Close()
	softWrapped := ww.String()

	final := wrap.String(softWrapped, m.viewport.Width)

	m.viewport.SetContent(final)
}
