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

//TODO: fix this once data is all in place.
func (f feedService) ServeContent(ctx context.Context, userFeed *model.UserFeed) ([]*model.ContentItem, error) {
	subscriptions, err := f.db.ListFeedSubscriptionsByFeedID(userFeed.ID)
	if err != nil {
		return nil, err
	}
	log.Tracef("Got feed subscription for feedID %v: %+v", userFeed.ID, subscriptions)
	engine, err := f.db.GetEngineByID(userFeed.EngineID)
	if err != nil {
		return nil, err
	}
	log.Tracef("Got engine to serve content: %+v", engine)

	var contentItems []*model.ContentItem
	var srcRSSFeeds []*model.SrcRSSFeed

	for _, sub := range subscriptions {
		if sub.SourceType == model.SourceTypeUserFeed {
			sourceUserFeed, err := f.db.GetUserFeedByID(sub.SourceID)
			if err != nil {
				return nil, err
			}
			sourceContentItems, err := f.ServeContent(ctx, sourceUserFeed)
			if err != nil {
				return nil, err
			}
			log.Tracef("found %v sourceContentItems from subscribed userFeed %v", len(sourceContentItems), sub.SourceID)
			contentItems = append(contentItems, sourceContentItems...)
		} else if sub.SourceType == model.SourceTypeSrcRSSFeed {
			log.Debug("sub source ID: ", sub.SourceID)
			srcRSSFeed, err := f.db.GetSrcRSSFeed(model.SrcRSSFeedInput{ID: &sub.SourceID})
			if err != nil {
				return nil, err
			}
			log.Debug("got srcRSSFeeds from DB. appending to srcList to build feed. srcRSSFeed: ", srcRSSFeed)
			if srcRSSFeed != nil {
				srcRSSFeeds = append(srcRSSFeeds, srcRSSFeed)
			}
		}
	}
	//Get all sources for feed, figure out what types they are
	//Recursively serve the other user feeds
	//Grab all the unique content items from the userFeeds' SrcRSSFeeds
	//Pass it all through the engine...

	srcContentItems, err := f.db.ServeContentItems(srcRSSFeeds, engine.Sort, engine.StartDate, engine.EndDate)
	if err != nil {
		log.Error("could not serve feed. err: ", err)
	}
	log.Debug("length of srcContentItems: ", len(srcContentItems))
	//Combine with source UserFeeds content
	contentItems = append(srcContentItems, contentItems...)
	//Sort it again?
	//TODO SORT AND FILTER THIS SHIT
	log.Debug("length of contentItems: ", len(contentItems))
	return contentItems, nil
}
