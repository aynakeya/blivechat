package internal

import (
	"blivechat/model"
	"bytes"
	tea "github.com/charmbracelet/bubbletea"
	"io"
	"os"
	"strings"
)

type teaLogWriter struct {
	p   *tea.Program
	buf bytes.Buffer
}

func NewTeaLogWriter(p *tea.Program) io.Writer {
	return &teaLogWriter{p: p}
}

func (w *teaLogWriter) Write(p []byte) (int, error) {
	n, _ := w.buf.Write(p)

	for {
		data := w.buf.Bytes()
		idx := bytes.IndexByte(data, '\n')
		if idx < 0 {
			break
		}
		line := string(data[:idx])
		w.buf.Next(idx + 1)
		w.p.Send(&model.DebugLineMsg{Line: strings.TrimRight(line, "\r\n")})
	}
	return n, nil
}

func CreateLogWriter(p *tea.Program, logfile string) (io.Writer, func() error) {
	if logfile != "" {
		f, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		return io.MultiWriter(f, NewTeaLogWriter(p)), func() error {
			return f.Close()
		}
	}
	return NewTeaLogWriter(p), func() error {
		return nil
	}
}
