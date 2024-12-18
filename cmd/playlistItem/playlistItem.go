package playlistItem

import (
	"github.com/spf13/cobra"
	"github.com/zamai/yutu/cmd"
)

var (
	id                     string
	title                  string
	description            string
	kind                   string
	kVideoId               string
	kChannelId             string
	kPlaylistId            string
	videoId                string
	playlistId             string
	channelId              string
	maxResults             int64
	privacy                string
	output                 string
	parts                  []string
	credential             string
	cacheToken             string
	onBehalfOfContentOwner string
)

var playlistItemCmd = &cobra.Command{
	Use:   "playlistItem",
	Short: "Manipulate YouTube playlist items",
	Long:  "Manipulate YouTube playlist items, such as insert, update, etc.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	cmd.RootCmd.AddCommand(playlistItemCmd)

	playlistItemCmd.PersistentFlags().StringVarP(&credential, "credential", "", "client_secret.json", "Path to client secret file")
	playlistItemCmd.PersistentFlags().StringVarP(&cacheToken, "cacheToken", "", "youtube.token.json", "Path to token cache file")
}
