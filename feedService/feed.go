package feed

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed"
	"github.com/well-informed/wellinformed/graph/model"
)

type feedService struct {
	db wellinformed.Persistor
}

func NewFeedService(db wellinformed.Persistor) *feedService {
	return &feedService{
		db: db,
	}
}

func (f feedService) Serve(ctx context.Context, user *model.User) (*model.UserFeed, error) {
	userSources, err := f.db.ListSrcRSSFeedsByUser(user)
	if err != nil {
		return nil, err
	}
	var contentItems []*model.ContentItem
	for _, src := range userSources {
		srcItems, err := f.db.ListContentItemsBySource(src)
		if err != nil {
			log.Error("could not retrieve content for source: ", src)
		}
		contentItems = append(contentItems, srcItems...)
	}
	return &model.UserFeed{
		UserID:       user.ID,
		Name:         "default",
		ContentItems: contentItems,
	}, nil
}
