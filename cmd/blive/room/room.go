package room

import (
	"errors"
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
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if flagCookie == "" {
			flagCookie = os.Getenv("bilibili_cookie")
		}
		if flagCookie == "" {
			log.Error("no cookie provided, please provide cookie")
			return errors.New("no cookie provided, please provide cookie")
		}
		cookies, err := http.ParseCookie(strings.TrimSpace(flagCookie))
		if err != nil {
			log.Errorf("invalid cookie provided, please provide correct cookie: %v", err)
			return errors.New("invalid cookie provided, please provide correct cookie")
		}
		for _, c := range cookies {
			if c.Name == "bili_jct" {
				csrfToken = c.Value
			}
		}
		if csrfToken == "" {
			log.Error("no bili_jct cookie provided, please provide correct cookie")
			return errors.New("no cookie provided, please provide correct cookie")
		}
		client = resty.New().SetCookies(cookies)
		return nil
	},
}
