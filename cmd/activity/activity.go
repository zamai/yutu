package activity

import (
	"github.com/spf13/cobra"
	"github.com/zamai/yutu/cmd"
)

var (
	channelId       string
	home            bool
	maxResults      int64
	mine            bool
	publishedAfter  string
	publishedBefore string
	regionCode      string
	parts           []string
	output          string
	credential      string
	cacheToken      string
)

var activityCmd = &cobra.Command{
	Use:   "activity",
	Short: "List YouTube activities",
	Long:  "List YouTube activities",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	cmd.RootCmd.AddCommand(activityCmd)

	activityCmd.PersistentFlags().StringVarP(
		&credential, "credential", "", "client_secret.json", "Path to client secret file",
	)
	activityCmd.PersistentFlags().StringVarP(
		&cacheToken, "cacheToken", "", "youtube.token.json", "Path to token cache file",
	)
}
