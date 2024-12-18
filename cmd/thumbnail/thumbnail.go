package thumbnail

import (
	"github.com/spf13/cobra"
	"github.com/zamai/yutu/cmd"
)

var (
	file       string
	videoId    string
	output     string
	credential string
	cacheToken string
)

var thumbnailCmd = &cobra.Command{
	Use:   "thumbnail",
	Short: "Set thumbnail for a video",
	Long:  "Set thumbnail for a video",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	cmd.RootCmd.AddCommand(thumbnailCmd)

	thumbnailCmd.PersistentFlags().StringVarP(&credential, "credential", "", "client_secret.json", "Path to client secret file")
	thumbnailCmd.PersistentFlags().StringVarP(&cacheToken, "cacheToken", "", "youtube.token.json", "Path to token cache file")
}
