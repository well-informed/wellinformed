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
	feed, err := r.RSS.FetchSrcFeed(feedLink, ctx)
	if err != nil {
		log.Errorf("couldn't fetch SrcFeed in order to add it.")
		return nil, err
	}
	insertedFeed, err := r.DB.InsertSrcRSSFeed(feed)
	if err != nil {
		return nil, err
	}

	// json, err := json.Marshal(feed)
	// if err != nil {
	// 	log.Error("feed object can't be json marshalled", err)
	// }
	// log.Info("manual json: ", string(json))
	// log.Infof("feed object to return: %+v", feed)
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
