package channelBanner

import (
	"errors"
	"fmt"
	"github.com/zamai/yutu/pkg/auth"
	"github.com/zamai/yutu/pkg/utils"
	"google.golang.org/api/youtube/v3"
	"log"
	"os"
)

var (
	service                *youtube.Service
	errInsertChannelBanner = errors.New("failed to insert channelBanner")
)

type channelBanner struct {
	File string `yaml:"file" json:"file"`

	OnBehalfOfContentOwner        string `yaml:"on_behalf_of_content_owner" json:"on_behalf_of_content_owner"`
	OnBehalfOfContentOwnerChannel string `yaml:"on_behalf_of_content_owner_channel" json:"on_behalf_of_content_owner_channel"`
}

type ChannelBanner interface {
	Insert(output string)
}

type Option func(banner *channelBanner)

func NewChannelBanner(opts ...Option) ChannelBanner {
	cb := &channelBanner{}

	for _, opt := range opts {
		opt(cb)
	}

	return cb
}

func (cb *channelBanner) Insert(output string) {
	file, err := os.Open(cb.File)
	if err != nil {
		utils.PrintJSON(cb)
		log.Fatalln(errors.Join(errInsertChannelBanner, err))
	}
	defer file.Close()
	cbr := &youtube.ChannelBannerResource{}

	call := service.ChannelBanners.Insert(cbr).Media(file)
	if cb.OnBehalfOfContentOwner != "" {
		call = call.OnBehalfOfContentOwner(cb.OnBehalfOfContentOwner)
	}
	if cb.OnBehalfOfContentOwnerChannel != "" {
		call = call.OnBehalfOfContentOwnerChannel(cb.OnBehalfOfContentOwnerChannel)
	}

	res, err := call.Do()
	if err != nil {
		utils.PrintJSON(cb)
		log.Fatalln(errors.Join(errInsertChannelBanner, err))
	}

	switch output {
	case "json":
		utils.PrintJSON(res)
	case "yaml":
		utils.PrintYAML(res)
	case "silent":
	default:
		fmt.Printf("ChannelBanner inserted: %s\n", res.Url)
	}
}

func WithFile(file string) Option {
	return func(cb *channelBanner) {
		cb.File = file
	}
}

func WithOnBehalfOfContentOwner(onBehalfOfContentOwner string) Option {
	return func(cb *channelBanner) {
		cb.OnBehalfOfContentOwner = onBehalfOfContentOwner
	}
}

func WithOnBehalfOfContentOwnerChannel(onBehalfOfContentOwnerChannel string) Option {
	return func(cb *channelBanner) {
		cb.OnBehalfOfContentOwnerChannel = onBehalfOfContentOwnerChannel
	}
}

func WithService(svc *youtube.Service) Option {
	return func(cb *channelBanner) {
		if svc != nil {
			service = svc
		} else {
			service = auth.NewY2BService()
		}
	}
}
