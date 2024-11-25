package playlist

import (
	"github.com/spf13/cobra"
	"github.com/zamai/yutu/pkg/auth"
	"github.com/zamai/yutu/pkg/playlist"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a playlist",
	Long:  "Delete a playlist",
	Run: func(cmd *cobra.Command, args []string) {
		p := playlist.NewPlaylist(
			playlist.WithID(id),
			playlist.WithOnBehalfOfContentOwner(onBehalfOfContentOwner),
			playlist.WithService(auth.NewY2BService(
				auth.WithCredential(credential),
				auth.WithCacheToken(cacheToken),
			)),
		)
		p.Delete()
	},
}

func init() {
	playlistCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVarP(
		&id, "id", "i", "", "ID of playlist to be deleted",
	)
	deleteCmd.Flags().StringVarP(
		&onBehalfOfContentOwner, "onBehalfOfContentOwner", "b", "", "",
	)
}
