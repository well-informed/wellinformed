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

func (rss *RSS) FetchSrcFeed(feedLink string, ctx context.Context) (model.SrcRSSFeed, []*model.ContentItem, error) {
	feed, err := rss.ParseURLWithContext(feedLink, ctx)
	if err != nil {
		log.Errorf("could not parse feed at url: %v. err: %v", feedLink, err)
		return model.SrcRSSFeed{}, nil, err
	}

	convertedFeed := convertToModelFeed(feed)

	contentItems := convertModelContentItems(feed.Items, feed)

	return *convertedFeed, contentItems, nil
}

func convertToModelFeed(feed *gofeed.Feed) *model.SrcRSSFeed {
	modelFeed := &model.SrcRSSFeed{
		Title:         feed.Title,
		Description:   &feed.Description,
		Link:          feed.Link,
		FeedLink:      feed.FeedLink,
		Updated:       *feed.UpdatedParsed,
		LastFetchedAt: time.Now(),
		Language:      &feed.Language,
		Generator:     &feed.Generator,
	}

	return modelFeed
}

//Converts the items returned on the gofeed.Feed struct into their internal representation
//Requires a model.SrcRSSFeed so the source ID can be attached to the contentItems for reference
func convertModelContentItems(items []*gofeed.Item, feed *gofeed.Feed) []*model.ContentItem {
	var contentItems []*model.ContentItem
	for _, item := range items {
		contentItems = append(contentItems, convertToModelContentItem(item, feed))
	}
	return contentItems
}

func convertToModelContentItem(item *gofeed.Item, feed *gofeed.Feed) *model.ContentItem {
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
