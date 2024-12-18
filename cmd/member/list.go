package member

import (
	"github.com/spf13/cobra"
	"github.com/zamai/yutu/pkg/auth"
	"github.com/zamai/yutu/pkg/member"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List channel's members",
	Long:  "List channel's members' info, such as channelId, displayName, etc.",
	Run: func(cmd *cobra.Command, args []string) {
		m := member.NewMember(
			member.WithMemberChannelId(memberChannelId),
			member.WithHasAccessToLevel(hasAccessToLevel),
			member.WithMaxResults(maxResults),
			member.WithMode(mode),
			member.WithService(auth.NewY2BService(
				auth.WithCredential(credential),
				auth.WithCacheToken(cacheToken),
			)),
		)
		m.List(parts, output)
	},
}

func init() {
	memberCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(
		&memberChannelId, "memberChannelId", "c", "",
		"Comma separated list of channel IDs. Only data about members that are part of this list will be included",
	)
	listCmd.Flags().StringVarP(
		&hasAccessToLevel, "hasAccessToLevel", "a", "",
		"Filter members in the results set to the ones that have access to a level",
	)
	listCmd.Flags().Int64VarP(
		&maxResults, "maxResults", "n", 5, "The maximum number of items that should be returned",
	)
	listCmd.Flags().StringVarP(
		&mode, "mode", "m", "",
		"listMembersModeUnknown, updates or all_current(default)",
	)
	listCmd.Flags().StringSliceVarP(
		&parts, "parts", "p", []string{"snippet"}, "Comma separated parts",
	)
	listCmd.Flags().StringVarP(
		&output, "output", "o", "", "json or yaml",
	)
}
