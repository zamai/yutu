package channelBanner

import (
	"github.com/spf13/cobra"
	"github.com/zamai/yutu/pkg/auth"
	"github.com/zamai/yutu/pkg/channelBanner"
)

var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert a ChannelBanner",
	Long:  "Insert a ChannelBanner",
	Run: func(cmd *cobra.Command, args []string) {
		cb := channelBanner.NewChannelBanner(
			channelBanner.WithFile(file),
			channelBanner.WithOnBehalfOfContentOwner(onBehalfOfContentOwner),
			channelBanner.WithOnBehalfOfContentOwnerChannel(onBehalfOfContentOwnerChannel),
			channelBanner.WithService(auth.NewY2BService(
				auth.WithCredential(credential),
				auth.WithCacheToken(cacheToken),
			)),
		)
		cb.Insert(output)
	},
}

func init() {
	channelBannerCmd.AddCommand(insertCmd)

	insertCmd.Flags().StringVarP(
		&file, "file", "f", "", "Path to the banner image",
	)
	insertCmd.Flags().StringVarP(
		&onBehalfOfContentOwner, "onBehalfOfContentOwner", "b", "", "",
	)
	insertCmd.Flags().StringVarP(
		&onBehalfOfContentOwnerChannel, "onBehalfOfContentOwnerChannel", "B", "", "",
	)
	insertCmd.Flags().StringVarP(&output, "output", "o", "", "json, yaml or silent")

	insertCmd.MarkFlagRequired("file")
}
