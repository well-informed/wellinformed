package wellinformed

import (
	"context"

	"github.com/well-informed/wellinformed/graph/model"
)

type Persistor interface {
	InsertSrcRSSFeed(model.SrcRSSFeed) (*model.SrcRSSFeed, error)
	SelectSrcRSSFeed(model.SrcRSSFeedInput) (*model.SrcRSSFeed, error)
	InsertContentItem(model.ContentItem) (*model.ContentItem, error)
	ListContentItemsBySource(*model.SrcRSSFeed) ([]*model.ContentItem, error)
	GetUserByEmail(string) (model.User, error)
	GetUserByUsername(string) (model.User, error)
	GetUserById(string) (model.User, error)
	CreateUser(model.User) (model.User, error)
}

type RSS interface {
	FetchSrcFeed(feedLink string, ctx context.Context) (model.SrcRSSFeed, []*model.ContentItem, error)
	WatchSrcFeed(feedLink string) error
}
