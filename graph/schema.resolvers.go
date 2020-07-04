package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/graph/generated"
	"github.com/well-informed/wellinformed/graph/model"
)

func (r *mutationResolver) AddSrcRSSFeed(ctx context.Context, feedLink string) (*model.SrcRSSFeed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	existingFeed, err := r.DB.SelectSrcRSSFeed(model.SrcRSSFeedInput{FeedLink: &feedLink})
	if err != nil {
		return nil, err
	}
	log.Debug("existingFeed: ", existingFeed)
	if existingFeed != nil {
		return existingFeed, nil
	}
	log.Debug("passed select, fetching feed")
	//TODO: Figure out how to resolve the problem where we can't be sure by the type def
	//whether or not the feed/contentItem has been assigned a database ID.
	//Maybe just assigned UUID internally and make the database accept them?
	feed, contentItems, err := r.RSS.FetchSrcFeed(feedLink, ctx)
	if err != nil {
		log.Errorf("couldn't fetch SrcFeed in order to add it.")
		return nil, err
	}
	insertedFeed, err := r.DB.InsertSrcRSSFeed(feed)
	if err != nil {
		return nil, err
	}
	log.Debug("inserted feed ID: ", insertedFeed.ID)
	for _, item := range contentItems {

		item.SourceID = insertedFeed.ID
		r.DB.InsertContentItem(*item)
	}
	return insertedFeed, nil
}

func (r *queryResolver) SrcRSSFeed(ctx context.Context, input *model.SrcRSSFeedInput) (*model.SrcRSSFeed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	feed, err := r.DB.SelectSrcRSSFeed(*input)
	log.Debug("after db select")
	if err != nil {
		return nil, err
	}
	return feed, nil
}

func (r *queryResolver) UserFeed(ctx context.Context, input int64) (*model.UserFeed, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *srcRSSFeedResolver) ContentItems(ctx context.Context, obj *model.SrcRSSFeed) ([]*model.ContentItem, error) {
	log.Debug("resolving ContentItems")
	contentItems, err := r.DB.ListContentItemsBySource(obj)
	if err != nil {
		return nil, err
	}
	return contentItems, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// SrcRSSFeed returns generated.SrcRSSFeedResolver implementation.
func (r *Resolver) SrcRSSFeed() generated.SrcRSSFeedResolver { return &srcRSSFeedResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type srcRSSFeedResolver struct{ *Resolver }
