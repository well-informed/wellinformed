package mock

import (
	"context"

	"github.com/well-informed/wellinformed/graph/model"
)

//Mock Subscriber implementation. simply does nothing for testing purposes
type Subscriber struct{}

func (sub Subscriber) SubscribeToRSSFeed(ctx context.Context, feedLink string) (*model.SrcRSSFeed, error) {
	return nil, nil
}

func (sub Subscriber) AddUserSubscription(user *model.User, srcRSSFeed *model.SrcRSSFeed) (*model.UserSubscription, error) {
	return nil, nil
}
