package playlistItem

import (
	"github.com/spf13/cobra"
	"github.com/zamai/yutu/pkg/auth"
	"github.com/zamai/yutu/pkg/playlistItem"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update playlist item",
	Long:  "Update playlist item's info, such as title, description, etc.",
	Run: func(cmd *cobra.Command, args []string) {
		pi := playlistItem.NewPlaylistItem(
			playlistItem.WithID(id),
			playlistItem.WithTitle(title),
			playlistItem.WithDescription(description),
			playlistItem.WithPrivacy(privacy),
			playlistItem.WithOnBehalfOfContentOwner(onBehalfOfContentOwner),
			playlistItem.WithService(auth.NewY2BService(
				auth.WithCredential(credential),
				auth.WithCacheToken(cacheToken),
			)),
		)
		pi.Update(output)
	},
}

func init() {
	playlistItemCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&id, "id", "i", "", "ID of the playlist item")
	updateCmd.Flags().StringVarP(
		&title, "title", "t", "", "Title of the playlist item",
	)
	updateCmd.Flags().StringVarP(
		&description, "description", "d", "", "Description of the playlist item",
	)
	updateCmd.Flags().StringVarP(
		&privacy, "privacy", "p", "", "Privacy status of the playlist item",
	)
	updateCmd.Flags().StringVarP(&onBehalfOfContentOwner, "onBehalfOfContentOwner", "b", "", "")
	updateCmd.Flags().StringVarP(&output, "output", "o", "", "json, yaml or silent")
}
