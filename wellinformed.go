package wellinformed

import (
	"context"

	"github.com/well-informed/wellinformed/graph/model"
)

type Persistor interface {
	InsertSrcRSSFeed(model.SrcRSSFeed) (model.SrcRSSFeed, error)
	SelectSrcRSSFeed(model.SrcRSSFeedInput) (model.SrcRSSFeed, error)
	GetUserByField(string, string) (model.User, error)
	CreateUser(model.User) (model.User, error)
}

type RSS interface {
	FetchSrcFeed(feedLink string, ctx context.Context) (model.SrcRSSFeed, error)
}
