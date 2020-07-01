package wellinformed

import (
	"context"

	"github.com/well-informed/wellinformed/graph/model"
)

type Persistor interface {
	InsertSrcRSSFeed(model.SrcRSSFeed) (*model.SrcRSSFeed, error)
	SelectSrcRSSFeed(model.SrcRSSFeedInput) (*model.SrcRSSFeed, error)
	InsertContentItem(model.ContentItem) (*model.ContentItem, error)
	ListContentItems() ([]model.ContentItem, error)
}

type RSS interface {
	FetchSrcFeed(feedLink string, ctx context.Context) (model.SrcRSSFeed, error)
	WatchSrcFeed(feedLink string) error
}
