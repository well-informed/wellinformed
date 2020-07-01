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

func (rss *RSS) WatchSrcFeed(feedLink string) error {
	log.Panic("not implemented")
	return nil
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

func convertToModelContentItem(item *gofeed.Item, feed *model.SrcRSSFeed) *model.ContentItem {
	var author *string
	if item.Author != nil {
		author = &item.Author.Name
	}
	var imageTitle *string
	var imageURL *string
	if item.Image != nil {
		imageTitle = &item.Image.Title
		imageURL = &item.Image.URL
	}
	return &model.ContentItem{
		SourceID:    feed.ID,
		SourceTitle: feed.Title,
		SourceLink:  feed.Link,
		Title:       item.Title,
		Description: item.Description,
		Content:     item.Content,
		Link:        item.Link,
		Updated:     item.UpdatedParsed,
		Published:   item.PublishedParsed,
		Author:      author,
		GUID:        &item.GUID,
		ImageTitle:  imageTitle,
		ImageURL:    imageURL,
	}
}
