package subscription

import (
	"reflect"
	"testing"
)

func TestNewSubscription(t *testing.T) {
	type args struct {
		opts []Option
	}
	tests := []struct {
		name string
		args args
		want Subscription
	}{
		{
			name: "TestNewSubscription",
			args: args{
				opts: []Option{
					WithId("id"),
					WithSubscriberChannelId("subscriberChannelId"),
					WithDescription("description"),
					WithChannelId("channelId"),
					WithForChannelId("forChannelId"),
					WithMaxResults(10),
					WithMine("true"),
					WithMyRecentSubscribers("true"),
					WithMySubscribers("true"),
					WithOnBehalfOfContentOwner("contentOwner"),
					WithOnBehalfOfContentOwnerChannel("contentOwnerChannel"),
					WithOrder("relevance"),
				},
			},
			want: &subscription{
				id:                            "id",
				subscriberChannelId:           "subscriberChannelId",
				description:                   "description",
				channelId:                     "channelId",
				forChannelId:                  "forChannelId",
				maxResults:                    10,
				mine:                          "true",
				myRecentSubscribers:           "true",
				mySubscribers:                 "true",
				onBehalfOfContentOwner:        "contentOwner",
				onBehalfOfContentOwnerChannel: "contentOwnerChannel",
				order:                         "relevance",
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := NewSubscription(tt.args.opts...); !reflect.DeepEqual(
					got, tt.want,
				) {
					t.Errorf("NewSubscription() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}