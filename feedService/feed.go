package feed

import (
	"context"
	"errors"

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
	prefSet, err := f.db.GetPreferenceSetByName(user.ID, user.ActivePreferenceSetName)
	if err != nil {
		return nil, errors.New("could not find user preference set")
	}

	userSources, err := f.db.ListSrcRSSFeedsByUser(user)
	if err != nil {
		return nil, err
	}

	contentItems, err := f.db.ServeContentItems(userSources, prefSet.Sort, prefSet.StartDate, prefSet.EndDate)
	if err != nil {
		log.Error("could not serve feed. err: ", err)
	}
	return &model.UserFeed{
		UserID:       user.ID,
		Name:         "default",
		ContentItems: contentItems,
	}, nil
}
