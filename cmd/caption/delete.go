package caption

import (
	"github.com/spf13/cobra"
	"github.com/zamai/yutu/pkg/auth"
	"github.com/zamai/yutu/pkg/caption"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete caption",
	Long:  "Delete caption of a video",
	Run: func(cmd *cobra.Command, args []string) {
		c := caption.NewCation(
			caption.WithID(id),
			caption.WithOnBehalfOf(onBehalfOf),
			caption.WithOnBehalfOfContentOwner(onBehalfOfContentOwner),
			caption.WithService(auth.NewY2BService(
				auth.WithCredential(credential),
				auth.WithCacheToken(cacheToken),
			)),
		)
		c.Delete()
	},
}

func init() {
	captionCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVarP(&id, "id", "i", "", "ID of the caption")
	deleteCmd.Flags().StringVarP(&onBehalfOf, "onBehalfOf", "b", "", "")
	deleteCmd.Flags().StringVarP(&onBehalfOfContentOwner, "onBehalfOfContentOwner", "B", "", "")
}
