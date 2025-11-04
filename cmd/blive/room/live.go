package room

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"slices"
	"strconv"
	"strings"
	"time"
)

var liveCmd = &cobra.Command{
	Use:   "live",
	Short: "manage live status",
}

func init() {
	liveStartCmd.Flags().Int64VarP(&areaV2Code, "area", "a", 192, "sub area code (area_v2)")
	liveCmd.AddCommand(liveStartCmd, liveStopCmd)
}

var areaV2Code int64

type streamingInfo struct {
	addr     string
	code     string
	newLink  string
	protocol string
	provider string
}

// startLive starts live. areaCode should be sub area code (which is area_v2)
func startLive(roomId int64, areaCode int64) ([]streamingInfo, error) {
	formData := map[string]string{
		"access_key": "",
		"appkey":     "aae92bc66f3edfab",
		"area_v2":    strconv.FormatInt(areaCode, 10),
		"build":      "9343", // to update
		"room_id":    strconv.FormatInt(roomId, 10),
		"platform":   "pc_link",
		"csrf_token": csrfToken,
		"csrf":       csrfToken,
		"ts":         strconv.FormatInt(time.Now().Unix(), 10),
	}
	keys := make([]string, 0)
	for k := range formData {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	sortedForm := make([]string, 0)
	for _, k := range keys {
		sortedForm = append(sortedForm, k+"="+formData[k])
	}
	md5hash := md5.Sum([]byte(strings.Join(sortedForm, "&") + "af125a0d5279fd576c1b4418a3e8276d"))
	formData["sign"] = hex.EncodeToString(md5hash[:])
	post, err := client.R().SetFormData(formData).Post("https://api.live.bilibili.com/room/v1/Room/startLive")
	if err != nil {
		return nil, err
	}
	data := gjson.ParseBytes(post.Bytes())
	msg := data.Get("message").String()
	if msg != "" {
		return nil, errors.New(msg)
	}
	info := streamingInfo{
		addr:     data.Get("data.rtmp.addr").String(),
		code:     data.Get("data.rtmp.code").String(),
		newLink:  data.Get("data.rtmp.new_link").String(),
		protocol: "rtmp",
		provider: "live",
	}
	return []streamingInfo{info}, nil
}

var liveStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start live room",
	RunE: func(cmd *cobra.Command, args []string) error {
		if areaV2Code == 0 {
			log.Error("area code is not valid")
			return nil
		}
		info, err := getInfo()
		if err != nil {
			log.Errorf("fail to get room id: %v", err)
			return nil
		}
		log.Infof("starting room %d with area=%d", info.roomId, areaV2Code)
		streams, err := startLive(info.roomId, areaV2Code)
		if err != nil {
			log.Errorf("fail to start live room: %v", err)
			return nil
		}
		log.Info("-- stream info --")
		for idx, stream := range streams {
			fmt.Printf("%02d: protocol=%s\n", idx+1, stream.protocol)
			fmt.Printf("Addr: %s\n", stream.addr)
			fmt.Printf("Code: %s\n", stream.code)
		}
		log.Info("-- end stream info --")
		return nil
	},
}

func stopLive(roomId int64) error {
	post, err := client.R().SetFormData(map[string]string{
		"room_id":    strconv.FormatInt(roomId, 10),
		"platform":   "pc",
		"csrf_token": csrfToken,
		"csrf":       csrfToken,
	}).Post("https://api.live.bilibili.com/room/v1/Room/stopLive")
	if err != nil {
		return err
	}
	msg := gjson.GetBytes(post.Bytes(), "message").String()
	if msg != "0" {
		return errors.New(msg)
	}
	return nil
}

var liveStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop live room",
	RunE: func(cmd *cobra.Command, args []string) error {
		info, err := getInfo()
		if err != nil {
			log.Errorf("fail to get room id: %v", err)
			return nil
		}
		log.Infof("stopping room %d", info.roomId)
		err = stopLive(info.roomId)
		if err != nil {
			log.Errorf("fail to stop live room: %v", err)
			return nil
		}
		log.Info("room stopped")
		return nil
	},
}
