package i18nRegion

import (
	"github.com/spf13/cobra"
	"github.com/zamai/yutu/cmd"
)

var (
	hl         string
	parts      []string
	output     string
	credential string
	cacheToken string
)

var i18nRegionCmd = &cobra.Command{
	Use:   "i18nRegion",
	Short: "List YouTube i18nRegions",
	Long:  "List YouTube i18nRegions",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	cmd.RootCmd.AddCommand(i18nRegionCmd)

	i18nRegionCmd.PersistentFlags().StringVarP(&credential, "credential", "", "client_secret.json", "Path to client secret file")
	i18nRegionCmd.PersistentFlags().StringVarP(&cacheToken, "cacheToken", "", "youtube.token.json", "Path to token cache file")
}
