package rss

import (
	"context"
	"time"

	"github.com/mmcdole/gofeed"
	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/graph/model"
)

type RSS struct {
	*gofeed.Parser
}

func NewRSS() *RSS {
	return &RSS{
		gofeed.NewParser(),
	}
}

func (rss *RSS) FetchSrcFeed(feedLink string, ctx context.Context) (model.SrcRSSFeed, error) {
	feed, err := rss.ParseURLWithContext(feedLink, ctx)
	if err != nil {
		log.Errorf("could not parse feed at url: %v. err: ", feedLink, err)
		return model.SrcRSSFeed{}, err
	}

	convertedFeed := convertToModelFeed(feed)
	return *convertedFeed, nil
}

func convertToModelFeed(feed *gofeed.Feed) *model.SrcRSSFeed {
	return &model.SrcRSSFeed{
		Title:         feed.Title,
		Description:   &feed.Description,
		Link:          feed.Link,
		FeedLink:      feed.FeedLink,
		Updated:       *feed.UpdatedParsed,
		LastFetchedAt: time.Now(),
		Language:      &feed.Language,
		Generator:     &feed.Generator,
	}
}
