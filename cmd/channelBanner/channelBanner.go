package channelBanner

import (
	"github.com/spf13/cobra"
	"github.com/zamai/yutu/cmd"
)

var (
	file                          string
	output                        string
	onBehalfOfContentOwner        string
	onBehalfOfContentOwnerChannel string
	credential                    string
	cacheToken                    string
)

var channelBannerCmd = &cobra.Command{
	Use:   "channelBanner",
	Short: "Insert Youtube channelBanner",
	Long:  "Insert Youtube channelBanner",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	cmd.RootCmd.AddCommand(channelBannerCmd)

	channelBannerCmd.PersistentFlags().StringVarP(&credential, "credential", "", "client_secret.json", "Path to client secret file")
	channelBannerCmd.PersistentFlags().StringVarP(&cacheToken, "cacheToken", "", "youtube.token.json", "Path to token cache file")
}
