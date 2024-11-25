package comment

import (
	"github.com/spf13/cobra"
	"github.com/zamai/yutu/pkg/auth"
	"github.com/zamai/yutu/pkg/comment"
)

var markAsSpamCmd = &cobra.Command{
	Use:   "markAsSpam",
	Short: "Mark YouTube comments as spam",
	Long:  "Mark YouTube comments as spam by ids",
	Run: func(cmd *cobra.Command, args []string) {
		c := comment.NewComment(
			comment.WithIDs(ids),
			comment.WithService(auth.NewY2BService(
				auth.WithCredential(credential),
				auth.WithCacheToken(cacheToken),
			)),
		)
		c.MarkAsSpam(output)
	},
}

func init() {
	commentCmd.AddCommand(markAsSpamCmd)

	markAsSpamCmd.Flags().StringSliceVarP(&ids, "ids", "i", []string{}, "Comma separated ids of comments")
	markAsSpamCmd.Flags().StringVarP(&output, "output", "o", "", "json, yaml or silent")
}
