package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/auth"
	"github.com/well-informed/wellinformed/graph/generated"
	"github.com/well-informed/wellinformed/graph/model"
)

func (r *contentItemResolver) Interaction(ctx context.Context, obj *model.ContentItem, input *model.ContentItemInteractionsInput) (*model.Interaction, error) {
	var userIdToUse int64

	currentUser, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	if input != nil {
		if input.UserID == nil {
			userIdToUse = currentUser.ID
		} else {
			userIdToUse = *input.UserID
		}
	} else {
		userIdToUse = currentUser.ID
	}

	return r.DB.GetInteractionByContentID(userIdToUse, obj.ID)
}

func (r *engineResolver) User(ctx context.Context, obj *model.Engine) (*model.User, error) {
	return r.DB.GetUserByID(obj.UserID)
}

func (r *interactionResolver) User(ctx context.Context, obj *model.Interaction) (*model.User, error) {
	return r.DB.GetUserByInteraction(obj.ID)
}

func (r *interactionResolver) ContentItem(ctx context.Context, obj *model.Interaction) (*model.ContentItem, error) {
	return r.DB.GetContentItemByInteraction(obj.ID)
}

func (r *mutationResolver) AddUserFeed(ctx context.Context, input model.AddUserFeedInput) (*model.UserFeed, error) {
	//Adding without cloning engine or source list for a start...
	return nil, errors.New("not implemented")
}

func (r *mutationResolver) AddSrcRSSFeed(ctx context.Context, feedLink string) (*model.SrcRSSFeed, error) {
	user, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	link, err := url.Parse(feedLink)
	if err != nil {
		log.Error("couldn't parse feedLink: ", feedLink)
		return nil, errors.New("couldn't parse feedLink")
	}
	if link.Scheme == "" {
		link.Scheme = "https"
	}
	feedLink = link.String()

	existingFeed, err := r.DB.GetSrcRSSFeed(model.SrcRSSFeedInput{FeedLink: &feedLink})
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

func (r *mutationResolver) AddSource(ctx context.Context, input model.AddSourceInput) (model.Feed, error) {
	panic(fmt.Errorf("not implemented"))
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

func (r *mutationResolver) SaveInteraction(ctx context.Context, input *model.InteractionInput) (*model.ContentItem, error) {
	user, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}
	//Ensure optional arguments are set to defaults before insert into DB
	var f = false
	if input.Completed == nil {
		input.Completed = &f
	}
	if input.SavedForLater == nil {
		input.SavedForLater = &f
	}
	return r.DB.SaveInteraction(user.ID, input)
}

func (r *mutationResolver) SaveEngine(ctx context.Context, engine model.EngineInput) (*model.Engine, error) {
	return r.UserService.SaveEngine(ctx, &engine)
}

func (r *queryResolver) SrcRSSFeed(ctx context.Context, input *model.SrcRSSFeedInput) (*model.SrcRSSFeed, error) {
	_, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}
	feed, err := r.DB.GetSrcRSSFeed(*input)
	if err != nil {
		return nil, err
	}
	log.Info("retrieved feed: ", feed)
	if feed == nil {
		return nil, errors.New("srcRSSFeed not found")
	}
	return feed, nil
}

func (r *queryResolver) Sources(ctx context.Context) ([]*model.SrcRSSFeed, error) {
	sources, err := r.DB.ListSrcRSSFeeds()
	if err != nil {
		return nil, err
	}
	if sources == nil {
		return nil, errors.New("no sources exist")
	}
	return sources, nil
}

func (r *queryResolver) UserFeed(ctx context.Context) (*model.UserFeed, error) {
	currentUser, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		log.Errorf("error while getting user feed: %v", err)
		return nil, errors.New("you are not signed in")
	}
	log.Printf("currentUser: %v", currentUser)
	return r.FeedService.Serve(ctx, currentUser)
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	currentUser, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		log.Errorf("error while getting user feed: %v", err)
		return nil, errors.New("you are not signed in")
	}
	return currentUser, nil
}

func (r *queryResolver) User(ctx context.Context, input *model.GetUserInput) (*model.User, error) {
	var user *model.User
	var err error

	if input.UserID != nil {
		user, err = r.DB.GetUserByID(*input.UserID)
	} else if input.Email != nil {
		user, err = r.DB.GetUserByEmail(*input.Email)
	} else if input.Username != nil {
		user, err = r.DB.GetUserByUsername(*input.Username)
	} else {
		return nil, errors.New("you need to provide an input")
	}

	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *queryResolver) GetContentItem(ctx context.Context, input int64) (*model.ContentItem, error) {
	_, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New("user not signed in")
	}
	contentItem, err := r.DB.GetContentItem(input)
	if err != nil {
		return nil, err
	}
	return contentItem, nil
}

func (r *queryResolver) GetInteractionByContentID(ctx context.Context, input int64) (*model.Interaction, error) {
	currentUser, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	return r.DB.GetInteractionByContentID(currentUser.ID, input)
}

func (r *queryResolver) Engines(ctx context.Context) ([]*model.Engine, error) {
	user, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}
	return r.DB.ListEnginesByUser(user.ID)
}

func (r *srcRSSFeedResolver) ContentItems(ctx context.Context, obj *model.SrcRSSFeed) ([]*model.ContentItem, error) {
	log.Debug("resolving ContentItems")
	contentItems, err := r.DB.ListContentItemsBySource(obj)
	if err != nil {
		return nil, err
	}
	return contentItems, nil
}

func (r *srcRSSFeedResolver) IsSubscribed(ctx context.Context, obj *model.SrcRSSFeed) (bool, error) {
	user, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return false, err
	}
	subscription, err := r.DB.GetUserSubscription(user.ID, obj.ID)
	if err != nil {
		return false, err
	}
	if subscription == nil {
		return false, nil
	}
	return true, nil
}

func (r *userResolver) Feed(ctx context.Context, obj *model.User) (*model.UserFeed, error) {
	return r.Query().UserFeed(ctx)
}

func (r *userResolver) Feeds(ctx context.Context, obj *model.User) ([]*model.UserFeed, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) SrcRSSFeeds(ctx context.Context, obj *model.User) ([]*model.SrcRSSFeed, error) {
	return r.DB.ListSrcRSSFeedsByUser(obj)
}

func (r *userResolver) Engines(ctx context.Context, obj *model.User) ([]*model.Engine, error) {
	return r.DB.ListEnginesByUser(obj.ID)
}

func (r *userResolver) ActiveEngine(ctx context.Context, obj *model.User) (*model.Engine, error) {
	return r.DB.GetEngineByName(obj.ID, obj.ActiveEngineName)
}

func (r *userResolver) Subscriptions(ctx context.Context, obj *model.User) ([]*model.UserSubscription, error) {
	user, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}
	return r.DB.ListUserSubscriptions(user.ID)
}

func (r *userResolver) Interactions(ctx context.Context, obj *model.User, input *model.UserInteractionsInput) ([]*model.Interaction, error) {
	var interactionInput *model.ReadState
	_, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		log.Errorf("error while getting Interactions for a user: %v", err)
		return nil, errors.New("you are not signed in")
	}

	if input == nil {
		interactionInput = nil
	} else {
		interactionInput = input.ReadState
	}

	return r.DB.ListUserInteractions(obj.ID, interactionInput)
}

func (r *userSubscriptionResolver) User(ctx context.Context, obj *model.UserSubscription) (*model.User, error) {
	user, err := r.DB.GetUserByID(obj.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("subscription's userID not found")
	}
	return user, nil
}

func (r *userSubscriptionResolver) SrcRSSFeed(ctx context.Context, obj *model.UserSubscription) (*model.SrcRSSFeed, error) {
	input := model.SrcRSSFeedInput{ID: &obj.SrcRSSFeedID}
	src, err := r.DB.GetSrcRSSFeed(input)
	if err != nil {
		return nil, err
	}
	return src, nil
}

// ContentItem returns generated.ContentItemResolver implementation.
func (r *Resolver) ContentItem() generated.ContentItemResolver { return &contentItemResolver{r} }

// Engine returns generated.EngineResolver implementation.
func (r *Resolver) Engine() generated.EngineResolver { return &engineResolver{r} }

// Interaction returns generated.InteractionResolver implementation.
func (r *Resolver) Interaction() generated.InteractionResolver { return &interactionResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// SrcRSSFeed returns generated.SrcRSSFeedResolver implementation.
func (r *Resolver) SrcRSSFeed() generated.SrcRSSFeedResolver { return &srcRSSFeedResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// UserSubscription returns generated.UserSubscriptionResolver implementation.
func (r *Resolver) UserSubscription() generated.UserSubscriptionResolver {
	return &userSubscriptionResolver{r}
}

type contentItemResolver struct{ *Resolver }
type engineResolver struct{ *Resolver }
type interactionResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type srcRSSFeedResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type userSubscriptionResolver struct{ *Resolver }
