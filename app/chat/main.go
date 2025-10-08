package main

import (
	"blivechat/internal"
	"blivechat/ui/base"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := base.NewModel()
	p := tea.NewProgram(m, tea.WithAltScreen())

	f, _ := os.OpenFile("blivechat.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	log.SetOutput(f)

	backend := internal.NewBackend(p, 3819533, "buvid3=284A792C-37F4-40FD-6A4E-45E26D1871CE92156infoc;  SESSDATA=a1f6bd3e%2C1775462777%2Ca3f75%2Aa1CjDf9qgCWGm1qwB-2oVoXAXGdXz6NaBlN-6lcbQHVjvLs015eBkcdYGuFdUp2sAvYvQSVjRBbkVrYlFHa0lQWDhNNkhnOVBTMnQyY0taV0l5ajh1ZU12NWNuVXB5WVJhTTRHZi11ZG1obzh1ZmRaZnFBYUpnU0hqSHlFS1hGa2plSnQ4YkZlYU5RIIEC; bili_jct=420558e735cc4365b5836e96cd0f3150;")

	go func() {
		err := backend.Run()
		if err != nil {
			log.Fatal(err)
			return
		}
	}()

	if _, err := p.Run(); err != nil {
		log.Println(err)
	}

	if err := backend.Stop(); err != nil {
		log.Println(err)
	}
	f.Close()
}
