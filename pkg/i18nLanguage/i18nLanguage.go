package i18nLanguage

import (
	"errors"
	"fmt"
	"github.com/zamai/yutu/pkg/auth"
	"github.com/zamai/yutu/pkg/utils"
	"google.golang.org/api/youtube/v3"
	"log"
)

var (
	service            *youtube.Service
	errGetI18nLanguage = errors.New("failed to get i18n language")
)

type i18nLanguage struct {
	Hl string `yaml:"hl" json:"hl"`
}

type I18nLanguage interface {
	get(parts []string) []*youtube.I18nLanguage
	List(parts []string, output string)
}

type Option func(*i18nLanguage)

func NewI18nLanguage(opts ...Option) I18nLanguage {
	i := &i18nLanguage{}

	for _, opt := range opts {
		opt(i)
	}

	return i
}

func (i *i18nLanguage) get(parts []string) []*youtube.I18nLanguage {
	call := service.I18nLanguages.List(parts)
	if i.Hl != "" {
		call = call.Hl(i.Hl)
	}

	res, err := call.Do()
	if err != nil {
		utils.PrintJSON(i)
		log.Fatalln(errors.Join(errGetI18nLanguage, err))
	}

	return res.Items
}

func (i *i18nLanguage) List(parts []string, output string) {
	i18nLanguages := i.get(parts)
	switch output {
	case "json":
		utils.PrintJSON(i18nLanguages)
	case "yaml":
		utils.PrintYAML(i18nLanguages)
	default:
		fmt.Println("id\thl\tname")
		for _, i18nLanguage := range i18nLanguages {
			fmt.Printf(
				"%v\t%v\t%v\n",
				i18nLanguage.Id, i18nLanguage.Snippet.Hl, i18nLanguage.Snippet.Name,
			)
		}
	}
}

func WithHl(hl string) Option {
	return func(i *i18nLanguage) {
		i.Hl = hl
	}
}

func WithService(svc *youtube.Service) Option {
	return func(i *i18nLanguage) {
		if svc != nil {
			service = svc
		} else {
			service = auth.NewY2BService()
		}
	}
}
