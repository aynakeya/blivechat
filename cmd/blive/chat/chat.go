package chat

import (
	"blivechat/internal"
	"blivechat/ui/got"
	"blivechat/ui/tab"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var logfile string
var flagCookie string

func init() {
	log.SetOutput(os.Stdout)
	ChatCmd.Flags().StringVarP(&logfile, "log-file", "l", "", "log file path, if not set, no log file will be created")
	ChatCmd.Flags().StringVarP(&flagCookie, "cookie", "c", os.Getenv("bilibili_cookie"), "cookie to use for session")
}

var ChatCmd = &cobra.Command{
	Use: "chat <room_id>",
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
			err := input.WithTheme(huh.ThemeBase()).Run()
			if err != nil {
				log.Error("invalid room id")
				return
			}
			roomId, _ = strconv.Atoi(roomIdStr)
		}
		runMain(roomId, flagCookie)
	},
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
