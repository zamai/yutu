package membershipsLevel

import (
	"github.com/spf13/cobra"
	"github.com/zamai/yutu/cmd"
)

var (
	parts      []string
	output     string
	credential string
	cacheToken string
)

var membershipsLevelCmd = &cobra.Command{
	Use:   "membershipsLevel",
	Short: "List YouTube memberships levels",
	Long:  "List YouTube memberships levels",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	cmd.RootCmd.AddCommand(membershipsLevelCmd)

	membershipsLevelCmd.PersistentFlags().StringVarP(&credential, "credential", "", "client_secret.json", "Path to client secret file")
	membershipsLevelCmd.PersistentFlags().StringVarP(&cacheToken, "cacheToken", "", "youtube.token.json", "Path to token cache file")
}
