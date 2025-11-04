package room

import (
	"errors"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

func setTitle(title string) error {
	post, err := client.R().
		SetFormData(map[string]string{
			"platform":   "pc",
			"mobi_app":   "pc",
			"build":      "1",
			"title":      title,
			"csrf_token": csrfToken,
			"csrf":       csrfToken,
		}).
		Post("https://api.live.bilibili.com/xlive/app-blink/v1/preLive/UpdatePreLiveInfo")
	if err != nil {
		return err
	}
	data := gjson.GetBytes(post.Bytes(), "message").String()
	if data != "0" {
		return errors.New("failed: " + data)
	}
	return nil
}

var setTitleCmd = &cobra.Command{
	Use:   "set-title <title>",
	Short: "set title for current room",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		if err := setTitle(title); err != nil {
			log.Errorf("failed to set title: %v", err)
			return
		}
		log.Infof("set title to %s ok", title)
		return
	},
}
