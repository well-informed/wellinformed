package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/auth"
	"github.com/well-informed/wellinformed/graph/generated"
	"github.com/well-informed/wellinformed/graph/model"
)

func (r *historyResolver) User(ctx context.Context, obj *model.History) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *historyResolver) ContentItem(ctx context.Context, obj *model.History) (*model.ContentItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddSrcRSSFeed(ctx context.Context, feedLink string) (*model.SrcRSSFeed, error) {
	user, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	existingFeed, err := r.DB.SelectSrcRSSFeed(model.SrcRSSFeedInput{FeedLink: &feedLink})
	if err != nil {
		return nil, err
	}
	log.Debug("existingFeed: ", existingFeed)
	log.Debugf("user: %v", user)

	if existingFeed != nil {
		_, err := r.Sub.AddUserSubscription(user, existingFeed)
		if err != nil {
			return existingFeed, err
		}
		return existingFeed, nil
	}
	insertedFeed, err := r.Sub.SubscribeToRSSFeed(ctx, feedLink)
	if err != nil {
		return nil, err
	}
	_, err = r.Sub.AddUserSubscription(user, insertedFeed)
	if err != nil {
		return nil, err
	}

	return insertedFeed, nil
}

func (r *mutationResolver) DeleteSubscription(ctx context.Context, srcRssfeedID int64) (*model.DeleteResponse, error) {
	user, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New("Something went wrong")
	}
	numDeleted, err := r.DB.DeleteUserSubscription(user.ID, srcRssfeedID)
	if err != nil {
		return nil, errors.New("Error deleting subscription")
	}
	if numDeleted == 0 {
		return nil, errors.New("Subscription ID does not exist")
	}
	return &model.DeleteResponse{
		Ok: true,
	}, nil
}

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	return r.UserService.Register(ctx, input)
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	return r.UserService.Login(ctx, input)
}

func (r *mutationResolver) UpdatePreferenceSet(ctx context.Context, input model.PreferenceSetInput) (*model.PreferenceSet, error) {
	return r.UserService.UpdatePreferenceSet(ctx, &input)
}

func (r *mutationResolver) SaveHistory(ctx context.Context, input *model.HistoryInput) (*model.History, error) {
	user, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}
	return r.DB.SaveHistory(user.ID, input)
}

func (r *preferenceSetResolver) User(ctx context.Context, obj *model.PreferenceSet) (*model.User, error) {
	return r.DB.GetUserById(obj.UserID)
}

func (r *queryResolver) SrcRSSFeed(ctx context.Context, input *model.SrcRSSFeedInput) (*model.SrcRSSFeed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	feed, err := r.DB.SelectSrcRSSFeed(*input)
	if err != nil {
		return nil, err
	}
	return feed, nil
}

func (r *queryResolver) UserFeed(ctx context.Context) (*model.UserFeed, error) {
	currentUser, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		log.Printf("error while getting user feed: %v", err)
		return nil, errors.New("You are not signed in!")
	}
	log.Printf("currentUser: %v", currentUser)
	return r.Feed.Serve(ctx, currentUser)
}

func (r *queryResolver) GetUser(ctx context.Context) (*model.User, error) {
	currentUser, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		log.Printf("error while getting user feed: %v", err)
		return nil, errors.New("You are not signed in!")
	}
	return currentUser, nil
}

func (r *queryResolver) GetContentItem(ctx context.Context, input int64) (*model.ContentItem, error) {
	_, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New("user not signed in")
	}
	contentItem, err := r.DB.SelectContentItem(input)
	if err != nil {
		return nil, err
	}
	return contentItem, nil
}

func (r *queryResolver) GetHistoryByContentID(ctx context.Context, input int64) (*model.History, error) {
	log.Debug("resolving GetHistoryByContentID")
	currentUser, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		log.Printf("error while getting user history: %v", err)
		return nil, errors.New("You are not signed in!")
	}
	history, err := r.DB.GetHistoryByContentID(currentUser.ID, input)
	if err != nil {
		return nil, err
	}
	return history, nil
}

func (r *srcRSSFeedResolver) ContentItems(ctx context.Context, obj *model.SrcRSSFeed) ([]*model.ContentItem, error) {
	log.Debug("resolving ContentItems")
	contentItems, err := r.DB.ListContentItemsBySource(obj)
	if err != nil {
		return nil, err
	}
	return contentItems, nil
}

func (r *userResolver) Feed(ctx context.Context, obj *model.User) (*model.UserFeed, error) {
	return r.Query().UserFeed(ctx)
}

func (r *userResolver) SrcRSSFeeds(ctx context.Context, obj *model.User) ([]*model.SrcRSSFeed, error) {
	return r.DB.ListSrcRSSFeedsByUser(obj)
}

func (r *userResolver) PreferenceSets(ctx context.Context, obj *model.User) ([]*model.PreferenceSet, error) {
	return r.DB.ListPreferenceSetsByUser(obj.ID)
}

func (r *userResolver) ActivePreferenceSet(ctx context.Context, obj *model.User) (*model.PreferenceSet, error) {
	return r.DB.GetPreferenceSetByName(obj.ID, obj.ActivePreferenceSet)
}

func (r *userResolver) History(ctx context.Context, obj *model.User) ([]*model.History, error) {
	panic(fmt.Errorf("not implemented"))
}

// History returns generated.HistoryResolver implementation.
func (r *Resolver) History() generated.HistoryResolver { return &historyResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// PreferenceSet returns generated.PreferenceSetResolver implementation.
func (r *Resolver) PreferenceSet() generated.PreferenceSetResolver { return &preferenceSetResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// SrcRSSFeed returns generated.SrcRSSFeedResolver implementation.
func (r *Resolver) SrcRSSFeed() generated.SrcRSSFeedResolver { return &srcRSSFeedResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type historyResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type preferenceSetResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type srcRSSFeedResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) ChangeActivePreferenceSet(ctx context.Context, input string) (*model.PreferenceSet, error) {
	panic(fmt.Errorf("not implemented"))
}
