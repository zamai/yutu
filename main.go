package main

import (
	"github.com/zamai/yutu/cmd"
	_ "github.com/zamai/yutu/cmd/activity"
	_ "github.com/zamai/yutu/cmd/caption"
	_ "github.com/zamai/yutu/cmd/channel"
	_ "github.com/zamai/yutu/cmd/channelBanner"
	_ "github.com/zamai/yutu/cmd/channelSection"
	_ "github.com/zamai/yutu/cmd/comment"
	_ "github.com/zamai/yutu/cmd/commentThread"
	_ "github.com/zamai/yutu/cmd/i18nLanguage"
	_ "github.com/zamai/yutu/cmd/i18nRegion"
	_ "github.com/zamai/yutu/cmd/member"
	_ "github.com/zamai/yutu/cmd/membershipsLevel"
	_ "github.com/zamai/yutu/cmd/playlist"
	_ "github.com/zamai/yutu/cmd/playlistItem"
	_ "github.com/zamai/yutu/cmd/search"
	_ "github.com/zamai/yutu/cmd/subscription"
	_ "github.com/zamai/yutu/cmd/superChatEvent"
	_ "github.com/zamai/yutu/cmd/thumbnail"
	_ "github.com/zamai/yutu/cmd/video"
	_ "github.com/zamai/yutu/cmd/videoAbuseReportReason"
	_ "github.com/zamai/yutu/cmd/videoCategory"
	_ "github.com/zamai/yutu/cmd/watermark"
)

//go:generate go-winres make --arch amd64 --product-version git-tag --file-version git-tag
func main() {
	cmd.Execute()
}
