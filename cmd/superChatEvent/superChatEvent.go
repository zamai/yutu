package superChatEvent

import (
	"github.com/spf13/cobra"
	"github.com/zamai/yutu/cmd"
)

var (
	hl         string
	maxResults int64
	parts      []string
	output     string
)

var superChatEventCmd = &cobra.Command{
	Use:   "superChatEvent",
	Short: "List Super Chat events for a channel",
	Long:  "List Super Chat events for a channel",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	cmd.RootCmd.AddCommand(superChatEventCmd)
}
