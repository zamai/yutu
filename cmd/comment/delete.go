package comment

import (
	"github.com/spf13/cobra"
	"github.com/zamai/yutu/pkg/auth"
	"github.com/zamai/yutu/pkg/comment"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete YouTube comments",
	Long:  "Delete YouTube comments by ids",
	Run: func(cmd *cobra.Command, args []string) {
		c := comment.NewComment(
			comment.WithIDs(ids),
			comment.WithService(auth.NewY2BService(
				auth.WithCredential(credential),
				auth.WithCacheToken(cacheToken),
			)),
		)
		c.Delete()
	},
}

func init() {
	commentCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringSliceVarP(&ids, "ids", "i", []string{}, "Comma separated ids of comments")
}
