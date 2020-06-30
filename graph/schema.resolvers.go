package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/graph/generated"
	"github.com/well-informed/wellinformed/graph/model"
)

func (r *mutationResolver) AddSrcRSSFeed(ctx context.Context, feedLink string) (*model.SrcRSSFeed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	feed, err := r.DB.SelectSrcRSSFeed(model.SrcRSSFeedInput{FeedLink: &feedLink})
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if feed != (model.SrcRSSFeed{}) {
		return &feed, nil
	}
	feed, err = r.RSS.FetchSrcFeed(feedLink, ctx)
	if err != nil {
		log.Errorf("couldn't fetch SrcFeed in order to add it.")
		return nil, err
	}
	feed, err = r.DB.InsertSrcRSSFeed(feed)
	if err != nil {
		return nil, err
	}

	// json, err := json.Marshal(feed)
	// if err != nil {
	// 	log.Error("feed object can't be json marshalled", err)
	// }
	// log.Info("manual json: ", string(json))
	// log.Infof("feed object to return: %+v", feed)
	return &feed, nil
}

func (r *queryResolver) SrcRSSFeed(ctx context.Context, input *model.SrcRSSFeedInput) (*model.SrcRSSFeed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	feed, err := r.DB.SelectSrcRSSFeed(*input)
	log.Debug("after db select")
	if err != nil {
		return nil, err
	}
	return &feed, nil
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
