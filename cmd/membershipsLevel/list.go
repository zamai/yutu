package membershipsLevel

import (
	"github.com/spf13/cobra"
	"github.com/zamai/yutu/pkg/auth"
	"github.com/zamai/yutu/pkg/membershipsLevel"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List memberships levels",
	Long:  "List memberships levels' info, such as id, displayName, etc.",
	Run: func(cmd *cobra.Command, args []string) {
		m := membershipsLevel.NewMembershipsLevel(
			membershipsLevel.WithService(auth.NewY2BService(
				auth.WithCredential(credential),
				auth.WithCacheToken(cacheToken),
			)))
		m.List(parts, output)
	},
}

func init() {
	membershipsLevelCmd.AddCommand(listCmd)

	listCmd.Flags().StringSliceVarP(
		&parts, "parts", "p", []string{"id, snippet"}, "Comma separated parts",
	)
	listCmd.Flags().StringVarP(
		&output, "output", "o", "", "format: json or yaml",
	)
}
