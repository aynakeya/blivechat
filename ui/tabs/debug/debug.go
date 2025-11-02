package debug

import (
	"blivechat/model"
	"strings"

	"github.com/muesli/reflow/wordwrap"
	"github.com/muesli/reflow/wrap"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type Style interface {
	SeparatorLine(width int) string // 你可以用一个小接口从全局 style 里渲染分隔线
}

type DebugTab struct {
	name     string
	lines    []string
	vp       viewport.Model
	width    int
	height   int
	sepLineF func(int) string
}

func NewDebugTab(sepLine func(int) string) *DebugTab {
	vp := viewport.New(0, 0)
	return &DebugTab{
		lines:    make([]string, 0, 256),
		vp:       vp,
		sepLineF: sepLine,
	}
}

func (t *DebugTab) TabName() string {
	return "Debug"
}

func (t *DebugTab) Init() tea.Cmd { return nil }

func (t *DebugTab) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m := msg.(type) {
	case *model.DebugLineMsg:
		if m.Line != "" {
			t.lines = append(t.lines, m.Line)
			t.refresh()
			t.vp.GotoBottom()
		}
	case tea.WindowSizeMsg:
		t.width, t.height = m.Width, m.Height
		t.vp.Width = t.width - 2
		t.vp.Height = t.height - 6
		t.refresh()
	}
	return t, nil
}

func (t *DebugTab) View() string {
	header := ""
	if t.sepLineF != nil {
		header = t.sepLineF(t.width) + "\n"
	}
	footer := ""
	if t.sepLineF != nil {
		footer = "\n" + t.sepLineF(t.width)
	}
	return header + t.vp.View() + footer
}

func (t *DebugTab) refresh() {
	content := strings.Join(t.lines, "\n")

	ww := wordwrap.NewWriter(t.vp.Width)
	ww.KeepNewlines = true
	_, _ = ww.Write([]byte(content))
	_ = ww.Close()
	softWrapped := ww.String()

	final := wrap.String(softWrapped, t.vp.Width)

	t.vp.SetContent(final)
}
