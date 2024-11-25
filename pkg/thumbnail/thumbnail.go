package thumbnail

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
	service         *youtube.Service
	errSetThumbnail = errors.New("failed to set thumbnail")
)

type thumbnail struct {
	File    string `yaml:"file" json:"file"`
	VideoId string `yaml:"video_id" json:"video_id"`
}

type Thumbnail interface {
	Set(output string)
}

type Option func(*thumbnail)

func NewThumbnail(opts ...Option) Thumbnail {
	t := &thumbnail{}
	for _, opt := range opts {
		opt(t)
	}
	return t
}

func (t *thumbnail) Set(output string) {
	file, err := os.Open(t.File)
	if err != nil {
		utils.PrintJSON(t)
		log.Fatalln(errors.Join(errSetThumbnail, err))
	}
	call := service.Thumbnails.Set(t.VideoId).Media(file)
	res, err := call.Do()
	if err != nil {
		utils.PrintJSON(t)
		log.Fatalln(errors.Join(errSetThumbnail, err))
	}

	switch output {
	case "json":
		utils.PrintJSON(res)
	case "yaml":
		utils.PrintYAML(res)
	case "silent":
	default:
		fmt.Println("Thumbnail set done")
	}
}

func WithVideoId(videoId string) Option {
	return func(t *thumbnail) {
		t.VideoId = videoId
	}
}

func WithFile(file string) Option {
	return func(t *thumbnail) {
		t.File = file
	}
}

func WithService(svc *youtube.Service) Option {
	return func(t *thumbnail) {
		if svc != nil {
			service = svc
		} else {
			service = auth.NewY2BService()
		}
	}
}
