package video

import (
	"github.com/spf13/cobra"
	"github.com/zamai/yutu/pkg/auth"
	"github.com/zamai/yutu/pkg/video"
)

var getRatingCmd = &cobra.Command{
	Use:   "getRating",
	Short: "Get the rating of a video",
	Long:  "Get the rating of a video, with the specified video ID",
	Run: func(cmd *cobra.Command, args []string) {
		v := video.NewVideo(
			video.WithID(id),
			video.WithOnBehalfOfContentOwner(onBehalfOfContentOwner),
			video.WithService(auth.NewY2BService(
				auth.WithCredential(credential),
				auth.WithCacheToken(cacheToken),
			)),
		)
		v.GetRating()
	},
}

func init() {
	videoCmd.AddCommand(getRatingCmd)

	getRatingCmd.Flags().StringVarP(&id, "id", "i", "", "ID of the video")
	getRatingCmd.Flags().StringVarP(
		&onBehalfOfContentOwner, "onBehalfOfContentOwner", "b", "", "",
	)
	getRatingCmd.MarkFlagRequired("id")
}
