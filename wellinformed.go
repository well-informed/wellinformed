package wellinformed

import (
	"context"

	"github.com/well-informed/wellinformed/graph/model"
)

type Persistor interface {
	InsertSrcRSSFeed(model.SrcRSSFeed) (*model.SrcRSSFeed, error)
	SelectSrcRSSFeed(model.SrcRSSFeedInput) (*model.SrcRSSFeed, error)
	ListSrcRSSFeeds() ([]*model.SrcRSSFeed, error)
	ListSrcRSSFeedsByUser(*model.User) ([]*model.SrcRSSFeed, error)
	InsertUserSubscription(model.User, model.SrcRSSFeed) (*model.UserSubscription, error)
	SelectUserSubscription(int64, int64) (*model.UserSubscription, error)
	DeleteUserSubscription(int64, int64) (int, error)
	SelectContentItem(int64) (*model.ContentItem, error)
	InsertContentItem(model.ContentItem) (*model.ContentItem, error)
	ListContentItemsBySource(*model.SrcRSSFeed) ([]*model.ContentItem, error)
	GetUserByEmail(string) (*model.User, error)
	GetUserByUsername(string) (*model.User, error)
	GetUserById(int64) (*model.User, error)
	CreateUser(model.User) (model.User, error)
}

type RSS interface {
	FetchSrcFeed(feedLink string, ctx context.Context) (model.SrcRSSFeed, []*model.ContentItem, error)
}

type Subscriber interface {
	SubscribeToRSSFeed(ctx context.Context, feedLink string) (*model.SrcRSSFeed, error)
	AddUserSubscription(user *model.User, srcRSSFeed *model.SrcRSSFeed) (*model.UserSubscription, error)
}

type FeedService interface {
	Serve(ctx context.Context, user *model.User) (*model.UserFeed, error)
}
