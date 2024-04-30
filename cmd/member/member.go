package member

import (
	"github.com/eat-pray-ai/yutu/cmd"
	"github.com/spf13/cobra"
)

var (
	memberChannelId string
	parts           []string
	output          string
)

var memberCmd = &cobra.Command{
	Use:   "member",
	Short: "manipulate YouTube members",
	Long:  "manipulate YouTube members, only list for now",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	cmd.RootCmd.AddCommand(memberCmd)
}