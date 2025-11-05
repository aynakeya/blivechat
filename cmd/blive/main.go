package main

import (
	"blivechat/cmd/blive/chat"
	"blivechat/cmd/blive/room"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"os"
)

var mainCmd = &cobra.Command{
	Use: "blive",
}

func init() {
	mainCmd.AddCommand(chat.ChatCmd, room.RoomCmd)
}

func main() {
	if err := mainCmd.Execute(); err != nil {
		log.SetOutput(os.Stdout)
	}
}
