package main

import (
	"blivechat/cmd/blive/chat"
	"blivechat/cmd/blive/room"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"os"
)

const (
	Version = "0.1.0"
	Hash    = "00000000"
)

var mainCmd = &cobra.Command{
	Use:   "blive",
	Short: fmt.Sprintf("bilibili live cli, %s (%s)", Version, Hash),
}

func init() {
	mainCmd.AddCommand(chat.ChatCmd, room.RoomCmd)
}

func main() {
	if err := mainCmd.Execute(); err != nil {
		log.SetOutput(os.Stdout)
	}
}
