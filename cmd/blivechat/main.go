package main

import (
	"blivechat/internal"
	"blivechat/ui/got"
	"blivechat/ui/tab"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

var logfile string
var flagCookie string

func init() {
	log.SetOutput(os.Stdout)
	mainCmd.Flags().StringVarP(&logfile, "log-file", "l", "", "log file path, if not set, no log file will be created")
	mainCmd.Flags().StringVarP(&flagCookie, "cookie", "c", os.Getenv("bilibili_cookie"), "cookie to use for session")
}

var mainCmd = &cobra.Command{
	Use: "blivechat <room_id>",
	Run: func(cmd *cobra.Command, args []string) {
		var roomId int
		var err error
		if flagCookie == "" {
			log.Error("no cookie provided, some features may not work properly")
			return
		}
		if len(args) >= 1 {
			roomId, err = strconv.Atoi(args[0])
			if err != nil {
				log.Error("invalid room id")
				return
			}
		} else {
			var roomIdStr string
			input := huh.NewInput().
				Title("Please Enter Room Id").
				Prompt("> ").
				Validate(func(val string) error {
					_, err := strconv.Atoi(val)
					return err
				}).Value(&roomIdStr)
			err := input.Run()
			if err != nil {
				log.Error("invalid room id")
				return
			}
			roomId, _ = strconv.Atoi(roomIdStr)
		}
		runMain(roomId, flagCookie)
	},
}

func main() {
	if err := mainCmd.Execute(); err != nil {
		log.SetOutput(os.Stdout)
		log.Error(err)
	}
}

func runMain(roomId int, cookie string) {
	m := tab.NewModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	backend := internal.NewBackend(p, roomId, cookie)
	logWriter, logClos := internal.CreateLogWriter(p, logfile)
	defer logClos()
	log.SetOutput(logWriter)
	got.Backend = backend
	go func() {
		go backend.UpdateRoomInfo()
		log.Info("blivechat started")
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
}
