package room

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"resty.dev/v3"
	"strings"
)

var logfile string
var flagCookie string
var client *resty.Client
var csrfToken string

func init() {
	log.SetOutput(os.Stdout)
	RoomCmd.PersistentFlags().StringVarP(&logfile, "log-file", "l", "", "log file path, if not set, no log file will be created")
	RoomCmd.PersistentFlags().StringVarP(&flagCookie, "cookie", "c", "", "cookie to use for session")
	RoomCmd.AddCommand(setTitleCmd, infoCmd, liveCmd)
}

var RoomCmd = &cobra.Command{
	Use: "room",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if flagCookie == "" {
			flagCookie = os.Getenv("bilibili_cookie")
		}
		if flagCookie == "" {
			log.Error("no cookie provided, please provide cookie")
			return
		}
		cookies, err := http.ParseCookie(strings.TrimSpace(flagCookie))
		if err != nil {
			log.Error("invalid cookie provided, please provide correct cookie")
			return
		}
		for _, c := range cookies {
			if c.Name == "bili_jct" {
				csrfToken = c.Value
			}
		}
		if csrfToken == "" {
			log.Error("no bili_jct cookie provided, please provide correct cookie")
			return
		}
		client = resty.New().SetCookies(cookies)
	},
}
