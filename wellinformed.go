package wellinformed

import (
	"context"
	"time"

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
	ServeContentItems([]*model.SrcRSSFeed, model.SortType, *time.Time, *time.Time) ([]*model.ContentItem, error)
	GetUserByEmail(string) (*model.User, error)
	GetUserByUsername(string) (*model.User, error)
	GetUserById(int64) (*model.User, error)
	CreateUser(model.User) (model.User, error)
	CreatePreferenceSet(*model.PreferenceSet) (*model.PreferenceSet, error)
	ListPreferenceSetsByUser(int64) ([]*model.PreferenceSet, error)
	GetPreferenceSetByID(int64) (*model.PreferenceSet, error)
	GetPreferenceSetByName(int64, string) (*model.PreferenceSet, error)
	UpdatePreferenceSet(int64, string, *model.PreferenceSetInput) (*model.PreferenceSet, error)
	SaveHistory(int64, *model.HistoryInput) (*model.History, error)
	ListUserHistory(int64) ([]*model.History, error)
	GetHistoryByContentID(int64, int64) (*model.History, error)
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
