package userFeed

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed"
	"github.com/well-informed/wellinformed/graph/model"
	"github.com/well-informed/wellinformed/pagination"
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
	engine, err := f.db.GetEngineByName(user.ID, user.ActiveEngineName)
	if err != nil {
		return nil, errors.New("could not find user preference set")
	}

	userSources, err := f.db.ListSrcRSSFeedsByUser(user)
	if err != nil {
		return nil, err
	}

	contentItems, err := f.db.ServeContentItems(userSources, engine.Sort, engine.StartDate, engine.EndDate)
	if err != nil {
		log.Error("could not serve feed. err: ", err)
	}
	//TODO: Accept page input to fill this out properly
	contentItemsPage, err := pagination.BuildContentItemPage(100, nil, contentItems)
	if err != nil {
		log.Error("could not build contentItemsPage in order to serve feed. err: ", err)
	}
	return &model.UserFeed{
		UserID:       user.ID,
		Name:         "default",
		ContentItems: contentItemsPage,
	}, nil
}
