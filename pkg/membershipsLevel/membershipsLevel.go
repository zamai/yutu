package membershipsLevel

import (
	"errors"
	"fmt"
	"github.com/zamai/yutu/pkg/auth"
	"github.com/zamai/yutu/pkg/utils"
	"google.golang.org/api/youtube/v3"
	"log"
)

var (
	service                *youtube.Service
	errGetMembershipsLevel = errors.New("failed to get memberships level")
)

type membershipsLevel struct{}

type MembershipsLevel interface {
	List([]string, string)
	get([]string) []*youtube.MembershipsLevel
}

type Option func(*membershipsLevel)

func NewMembershipsLevel(opts ...Option) MembershipsLevel {
	m := &membershipsLevel{}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

func (m *membershipsLevel) get(parts []string) []*youtube.MembershipsLevel {
	call := service.MembershipsLevels.List(parts)
	res, err := call.Do()
	if err != nil {
		utils.PrintJSON(m)
		log.Fatalln(errors.Join(errGetMembershipsLevel, err))
	}

	return res.Items
}

func (m *membershipsLevel) List(parts []string, output string) {
	membershipsLevels := m.get(parts)
	switch output {
	case "json":
		utils.PrintJSON(membershipsLevels)
	case "yaml":
		utils.PrintYAML(membershipsLevels)
	default:
		fmt.Println("id\tdisplayName")
		for _, membershipsLevel := range membershipsLevels {
			fmt.Printf(
				"%v\t%v\n", membershipsLevel.Id,
				membershipsLevel.Snippet.LevelDetails.DisplayName,
			)
		}
	}
}

func WithService(svc *youtube.Service) Option {
	return func(m *membershipsLevel) {
		if svc != nil {
			service = svc
		} else {
			service = auth.NewY2BService()
		}
	}
}
