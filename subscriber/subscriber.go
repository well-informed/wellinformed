package subscriber

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed"
	"github.com/well-informed/wellinformed/graph/model"
)

type subscriber struct {
	rss            wellinformed.RSS
	db             wellinformed.Persistor
	subscriptions  chan model.SrcRSSFeed
	updateInterval time.Duration
}

//NewSubscriber provides a persistent service that allows the application to regularly check for updates to
//subscribed RSS feeds
func NewSubscriber(rss wellinformed.RSS, db wellinformed.Persistor) (*subscriber, error) {
	sub := &subscriber{
		rss:            rss,
		db:             db,
		subscriptions:  make(chan model.SrcRSSFeed),
		updateInterval: 5 * time.Minute,
	}
	sub.watchSubscriptions()
	err := sub.loadSubscriptions()
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (sub *subscriber) SubscribeToRSSFeed(ctx context.Context, feedLink string) (*model.SrcRSSFeed, error) {
	srcFeed, err := sub.updateSrcRSSFeed(ctx, feedLink)
	if err != nil {
		return nil, err
	}
	sub.subscriptions <- *srcFeed
	return srcFeed, nil
}

//Gets all the recorded subscriptions from the database so that the subscriptions can be loaded at startup
func (sub *subscriber) loadSubscriptions() error {
	srcRSSFeeds, err := sub.db.ListSrcRSSFeeds()
	if err != nil {
		return err
	}
	for _, srcRSSFeed := range srcRSSFeeds {
		go sub.updateSrcRSSFeed(context.Background(), srcRSSFeed.FeedLink)
		sub.subscriptions <- *srcRSSFeed
	}
	return nil
}

//fans out goroutines so that each source feed has it's own routine checking for updates on each interval
func (sub *subscriber) watchSubscriptions() {
	go func() {
		for {
			feed := <-sub.subscriptions
			log.Debug("launching subscription for feed: ", feed)
			go sub.updateSourceOnInterval(feed, sub.updateInterval)
		}
	}()
}

func (sub *subscriber) updateSourceOnInterval(srcFeed model.SrcRSSFeed, duration time.Duration) {
	t := time.NewTicker(duration)
	for {
		<-t.C
		log.Debug("running scheduled update for srcFeed: ", srcFeed.Title)
		sub.updateSrcRSSFeed(context.Background(), srcFeed.FeedLink)
	}
}

func (sub *subscriber) updateSrcRSSFeed(ctx context.Context, feedLink string) (*model.SrcRSSFeed, error) {
	//TODO: might want to wrap all this in a db transaction
	feed, contentItems, err := sub.rss.FetchSrcFeed(feedLink, ctx)
	if err != nil {
		log.Errorf("couldn't fetch SrcFeed in order to add it.")
		return nil, err
	}
	var storedFeed *model.SrcRSSFeed
	storedFeed, err = sub.db.GetSrcRSSFeed(model.SrcRSSFeedInput{FeedLink: &feedLink})
	if err != nil {
		return nil, err
	}
	if storedFeed == nil {
		storedFeed, err = sub.db.InsertSrcRSSFeed(feed)
		if err != nil {
			return nil, err
		}
	}

	log.Debug("stored feed ID: ", storedFeed.ID)
	for _, item := range contentItems {

		item.SourceID = storedFeed.ID
		sub.db.InsertContentItem(*item)
	}
	return storedFeed, nil
}

func (sub *subscriber) AddUserSubscription(user *model.User, srcRSSFeed *model.SrcRSSFeed) (*model.UserSubscription, error) {
	userSub, err := sub.db.GetUserSubscription(user.ID, srcRSSFeed.ID)
	if err != nil {
		return nil, err
	}
	if userSub != nil {
		log.Debug("existing user subscription for source found")
		return userSub, nil
	}
	userSub, err = sub.db.InsertUserSubscription(*user, *srcRSSFeed)
	if err != nil {
		return nil, err
	}
	return userSub, nil
}
