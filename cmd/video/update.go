package video

import (
	"github.com/eat-pray-ai/yutu/pkg/yutuber/video"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a video on YouTube",
	Long:  "Update a video on YouTube, with the specified title, description, tags, etc.",
	Run: func(cmd *cobra.Command, args []string) {
		v := video.NewVideo(
			video.WithId(id),
			video.WithTitle(title),
			video.WithDescription(description),
			video.WithTags(tags),
			video.WithLanguage(language),
			video.WithLicense(license),
			video.WithPlaylistId(playListId),
			video.WithThumbnail(thumbnail),
			video.WithCategory(categoryId),
			video.WithPrivacy(privacy),
			video.WithEmbeddable(embeddable),
			video.WithService(),
		)
		v.Update()
	},
}

func init() {
	videoCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&id, "id", "i", "", "ID of the video")
	updateCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the video")
	updateCmd.Flags().StringVarP(
		&description, "description", "d", "", "Description of the video",
	)
	updateCmd.Flags().StringArrayVarP(
		&tags, "tags", "a", []string{}, "Comma separated tags",
	)
	updateCmd.Flags().StringVarP(
		&language, "language", "l", "", "Language of the video",
	)
	updateCmd.Flags().StringVarP(
		&license, "license", "L", "youtube", "youtube(default) or creativeCommon",
	)
	updateCmd.Flags().StringVarP(
		&thumbnail, "thumbnail", "T", "", "Path to the thumbnail file",
	)
	updateCmd.Flags().StringVarP(
		&playListId, "playlistId", "y", "", "Playlist ID of the video",
	)
	updateCmd.Flags().StringVarP(
		&categoryId, "categoryId", "g", "", "Category of the video",
	)
	updateCmd.Flags().StringVarP(
		&privacy, "privacy", "p", "",
		"Privacy status of the video: public, private or unlisted",
	)
	updateCmd.Flags().BoolVarP(
		&embeddable, "embeddable", "e", true, "Whether the video is embeddable",
	)

	updateCmd.MarkFlagRequired("id")
}
