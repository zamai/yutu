package channelSection

import (
	"github.com/eat-pray-ai/yutu/cmd"
	"github.com/spf13/cobra"
)

var (
	id                     string
	channelId              string
	hl                     string
	mine                   bool
	onBehalfOfContentOwner string
	parts                  []string
	output                 string
)

var channelSectionCmd = &cobra.Command{
	Use:   "channelSection",
	Short: "Manipulate channel section",
	Long:  "List or delete channel section",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	cmd.RootCmd.AddCommand(channelSectionCmd)
}
