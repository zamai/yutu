package subscription

import (
	"errors"
	"fmt"
	"github.com/eat-pray-ai/yutu/pkg/auth"
	"github.com/eat-pray-ai/yutu/pkg/utils"
	"google.golang.org/api/youtube/v3"
	"log"
)

var (
	service               *youtube.Service
	errGetSubscription    = errors.New("failed to get subscription")
	errDeleteSubscription = errors.New("failed to delete subscription")
	errInsertSubscription = errors.New("failed to insert subscription")
)

type subscription struct {
	id                            string
	subscriberChannelId           string
	description                   string
	channelId                     string
	forChannelId                  string
	maxResults                    int64
	mine                          *bool
	myRecentSubscribers           *bool
	mySubscribers                 *bool
	onBehalfOfContentOwner        string
	onBehalfOfContentOwnerChannel string
	order                         string
	title                         string
}

type Subscription interface {
	get([]string) []*youtube.Subscription
	List([]string, string)
	Insert()
	Delete()
}

type Option func(*subscription)

func NewSubscription(opts ...Option) Subscription {
	s := &subscription{}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *subscription) get(parts []string) []*youtube.Subscription {
	call := service.Subscriptions.List(parts)
	if s.id != "" {
		call = call.Id(s.id)
	}
	if s.channelId != "" {
		call = call.ChannelId(s.channelId)
	}
	if s.forChannelId != "" {
		call = call.ForChannelId(s.forChannelId)
	}
	call = call.MaxResults(s.maxResults)

	if s.mine != nil {
		call = call.Mine(*s.mine)
	}
	if s.myRecentSubscribers != nil {
		call = call.MyRecentSubscribers(*s.myRecentSubscribers)
	}
	if s.mySubscribers != nil {
		call = call.MySubscribers(*s.mySubscribers)
	}

	if s.onBehalfOfContentOwner != "" {
		call = call.OnBehalfOfContentOwner(s.onBehalfOfContentOwner)
	}
	if s.onBehalfOfContentOwnerChannel != "" {
		call = call.OnBehalfOfContentOwnerChannel(s.onBehalfOfContentOwnerChannel)
	}
	if s.order != "" {
		call = call.Order(s.order)
	}

	res, err := call.Do()
	if err != nil {
		log.Fatalln(errors.Join(errGetSubscription, err))
	}

	return res.Items
}

func (s *subscription) List(parts []string, output string) {
	subscriptions := s.get(parts)
	switch output {
	case "json":
		utils.PrintJSON(subscriptions)
	case "yaml":
		utils.PrintYAML(subscriptions)
	default:
		fmt.Println("ID\tChannel ID\tChannel Title")
		for _, subscription := range subscriptions {
			fmt.Printf(
				"%s\t%s\t%s\n", subscription.Id,
				subscription.Snippet.ResourceId.ChannelId, subscription.Snippet.Title,
			)
		}
	}
}

func (s *subscription) Insert() {
	subscription := &youtube.Subscription{
		Snippet: &youtube.SubscriptionSnippet{
			ChannelId:   s.subscriberChannelId,
			Description: s.description,
			ResourceId: &youtube.ResourceId{
				ChannelId: s.channelId,
			},
			Title: s.title,
		},
	}

	call := service.Subscriptions.Insert([]string{"snippet"}, subscription)
	res, err := call.Do()
	if err != nil {
		log.Fatalln(errors.Join(errInsertSubscription, err))
	}
	fmt.Println("Subscription inserted")
	utils.PrintYAML(res)
}

func (s *subscription) Delete() {
	call := service.Subscriptions.Delete(s.id)
	err := call.Do()
	if err != nil {
		log.Fatalln(errors.Join(errDeleteSubscription, err), s.id)
	}

	fmt.Printf("Subscription %s deleted", s.id)
}

func WithId(id string) Option {
	return func(s *subscription) {
		s.id = id
	}
}

func WithSubscriberChannelId(id string) Option {
	return func(s *subscription) {
		s.subscriberChannelId = id
	}
}

func WithDescription(description string) Option {
	return func(s *subscription) {
		s.description = description
	}
}

func WithChannelId(channelId string) Option {
	return func(s *subscription) {
		s.channelId = channelId
	}
}

func WithForChannelId(forChannelId string) Option {
	return func(s *subscription) {
		s.forChannelId = forChannelId
	}
}

func WithMaxResults(maxResults int64) Option {
	return func(s *subscription) {
		s.maxResults = maxResults
	}
}

func WithMine(mine bool, changed bool) Option {
	return func(s *subscription) {
		if changed {
			s.mine = &mine
		}
	}
}

func WithMyRecentSubscribers(myRecentSubscribers bool, changed bool) Option {
	return func(s *subscription) {
		if changed {
			s.myRecentSubscribers = &myRecentSubscribers
		}
	}
}

func WithMySubscribers(mySubscribers bool, changed bool) Option {
	return func(s *subscription) {
		if changed {
			s.mySubscribers = &mySubscribers
		}
	}
}

func WithOnBehalfOfContentOwner(onBehalfOfContentOwner string) Option {
	return func(s *subscription) {
		s.onBehalfOfContentOwner = onBehalfOfContentOwner
	}
}

func WithOnBehalfOfContentOwnerChannel(onBehalfOfContentOwnerChannel string) Option {
	return func(s *subscription) {
		s.onBehalfOfContentOwnerChannel = onBehalfOfContentOwnerChannel
	}
}

func WithOrder(order string) Option {
	return func(s *subscription) {
		s.order = order
	}
}

func WithTitle(title string) Option {
	return func(s *subscription) {
		s.title = title
	}
}

func WithService() Option {
	return func(s *subscription) {
		service = auth.NewY2BService()
	}
}
