package main

import (
	"blivechat/internal"
	"blivechat/ui/got"
	"blivechat/ui/tab"
	"os"

	"github.com/charmbracelet/log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := tab.NewModel()
	p := tea.NewProgram(m, tea.WithAltScreen())

	f, err := os.OpenFile("blivechat.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	log.SetOutput(f)
	log.Info("blivechat started")

	backend := internal.NewBackend(p, 3819533, "")
	got.Backend = backend
	go func() {
		go backend.UpdateRoomInfo()
		err := backend.Run()
		if err != nil {
			log.Fatal(err)
			return
		}
	}()

	if _, err := p.Run(); err != nil {
		log.Info(err)
	}

	if err := backend.Stop(); err != nil {
		log.Info(err)
	}
	f.Close()
}
