package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/auth"
	"github.com/well-informed/wellinformed/graph/generated"
	"github.com/well-informed/wellinformed/graph/model"
	page "github.com/well-informed/wellinformed/pagination"
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

func (r *feedSubscriptionResolver) UserFeed(ctx context.Context, obj *model.FeedSubscription) (*model.UserFeed, error) {
	return r.DB.GetUserFeedByID(obj.UserFeedID)
}

func (r *feedSubscriptionResolver) Source(ctx context.Context, obj *model.FeedSubscription) (model.Feed, error) {
	if obj.SourceType == model.SourceTypeSrcRSSFeed {
		return r.DB.GetSrcRSSFeedByID(obj.SourceID)
	} else if obj.SourceType == model.SourceTypeUserFeed {
		return r.DB.GetUserFeedByID(obj.SourceID)
	}
	return nil, errors.New("sourceType not recognized")
}

func (r *interactionResolver) User(ctx context.Context, obj *model.Interaction) (*model.User, error) {
	return r.DB.GetUserByInteraction(obj.ID)
}

func (r *interactionResolver) ContentItem(ctx context.Context, obj *model.Interaction) (*model.ContentItem, error) {
	return r.DB.GetContentItemByInteraction(obj.ID)
}

func (r *mutationResolver) AddUserFeed(ctx context.Context, input model.AddUserFeedInput) (*model.UserFeed, error) {
	user, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	newEngine := &model.Engine{
		UserID: user.ID,
		Name:   input.Name,
		Sort:   model.SortTypeChronological,
	}

	savedEngine, err := r.DB.SaveEngine(newEngine)
	if err != nil {
		log.Error("couldn't create new engine for new user feed. err: ", err)
		return nil, err
	}
	//Adding without cloning engine or source list for a start...
	userFeed := &model.UserFeed{
		UserID:   user.ID,
		EngineID: savedEngine.ID,
		Title:    input.Name,
		Name:     input.Name,
	}
	createdUserFeed, err := r.DB.CreateUserFeed(userFeed)
	if err != nil {
		log.Error("could not create new userFeed. err: ", err)
		return nil, err
	}
	_, err = r.switchActiveUserFeed(user, createdUserFeed.ID)
	if err != nil {
		log.Error("could not switch active user feed")
		return nil, err
	}
	return createdUserFeed, nil
}

func (r *mutationResolver) AddSrcRSSFeed(ctx context.Context, feedLink string, targetFeedID int64) (*model.SrcRSSFeed, error) {
	_, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}
	return r.addSrcRSSFeed(ctx, feedLink, targetFeedID)
}

func (r *mutationResolver) AddSource(ctx context.Context, input model.AddSourceInput) (*model.FeedSubscription, error) {
	user, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}
	//TODO: Go back and handle concept of UserSubscription
	//If target feed is not supplied explicitly get the user's active feed
	var feedSubscription *model.FeedSubscription
	if input.TargetFeedID == nil {
		targetFeed, err := r.DB.GetUserFeedByID(user.ActiveUserFeedID)
		if err != nil {
			return nil, err
		}
		input.TargetFeedID = &targetFeed.ID
	}

	//Record Feed subscription
	feedSubscription, err = r.DB.CreateFeedSubscription(*input.TargetFeedID, input.SourceFeedID, input.SourceType)
	if err != nil {
		return nil, err
	}

	return feedSubscription, nil
}

func (r *mutationResolver) DeleteSubscription(ctx context.Context, srcRssfeedID int64) (*model.DeleteResponse, error) {
	user, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
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
	return r.saveInteraction(user, input)
}

func (r *mutationResolver) SaveEngine(ctx context.Context, engine model.EngineInput) (*model.Engine, error) {
	return r.UserService.SaveEngine(ctx, &engine)
}

func (r *mutationResolver) SwitchActiveUserFeed(ctx context.Context, feedID int64) (*model.User, error) {
	user, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}
	user, err = r.switchActiveUserFeed(user, feedID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *queryResolver) SrcRSSFeed(ctx context.Context, input *model.SrcRSSFeedInput) (*model.SrcRSSFeed, error) {
	_, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	var feed *model.SrcRSSFeed
	if input.ID != nil {
		feed, err = r.DB.GetSrcRSSFeedByID(*input.ID)
	} else if input.Link != nil {
		feed, err = r.DB.GetSrcRSSFeedByLink(*input.Link)
	} else if input.FeedLink != nil {
		feed, err = r.DB.GetSrcRSSFeedByFeedLink(*input.FeedLink)
	} else {
		return nil, errors.New("no SrcRSSFeedInput key found")
	}
	if err != nil {
		return nil, err
	}
	log.Debug("retrieved feed: ", feed)
	if feed == nil {
		return nil, errors.New("srcRSSFeed not found")
	}
	return feed, nil
}

func (r *queryResolver) Sources(ctx context.Context, input *model.SrcRSSFeedConnectionInput) (*model.SrcRSSFeedConnection, error) {
	feeds, err := r.DB.ListSrcRSSFeeds()
	if err != nil {
		return nil, err
	}
	if feeds == nil {
		return nil, errors.New("no sources exist")
	}
	return page.BuildSrcRSSFeedPage(input.First, input.After, feeds)
}

func (r *queryResolver) UserFeed(ctx context.Context) (*model.UserFeed, error) {
	currentUser, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		log.Errorf("error while getting user feed: %v", err)
		return nil, errors.New("you are not signed in")
	}
	log.Debugf("fetching currentUser %v userFeed", currentUser)
	activeFeed, err := r.DB.GetUserFeedByID(currentUser.ActiveUserFeedID)
	if err != nil {
		return nil, err
	}
	log.Debugf("retrieved UserFeed as active: %+v", activeFeed)
	return activeFeed, nil
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

func (r *srcRSSFeedResolver) ContentItems(ctx context.Context, obj *model.SrcRSSFeed, input model.ContentItemConnectionInput) (*model.ContentItemConnection, error) {
	items, err := r.DB.ListContentItemsBySource(obj)
	if err != nil {
		return nil, err
	}
	return page.BuildContentItemPage(input.First, input.After, items)
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
	//TODO: paginate
	return r.DB.ListUserFeedsByUser(obj.ID)
}

func (r *userResolver) SrcRSSFeeds(ctx context.Context, obj *model.User, input *model.SrcRSSFeedConnectionInput) (*model.SrcRSSFeedConnection, error) {
	feeds, err := r.DB.ListSrcRSSFeedsByUser(obj)
	if err != nil {
		return nil, err
	}
	return page.BuildSrcRSSFeedPage(input.First, input.After, feeds)
}

func (r *userResolver) Engines(ctx context.Context, obj *model.User) ([]*model.Engine, error) {
	return r.DB.ListEnginesByUser(obj.ID)
}

func (r *userResolver) Subscriptions(ctx context.Context, obj *model.User, input *model.UserSubscriptionConnectionInput) (*model.UserSubscriptionConnection, error) {
	user, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}
	subs, err := r.DB.ListUserSubscriptions(user.ID)
	if err != nil {
		return nil, err
	}
	return page.BuildUserSubscriptionPage(input.First, input.After, subs)
}

func (r *userResolver) Interactions(ctx context.Context, obj *model.User, readState *model.ReadState, input model.InteractionConnectionInput) (*model.InteractionConnection, error) {
	_, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		log.Errorf("error while getting Interactions for a user: %v", err)
		return nil, errors.New("you are not signed in")
	}
	interactions, err := r.DB.ListUserInteractions(obj.ID, readState)
	if err != nil {
		return nil, err
	}
	return page.BuildInteractionPage(input.First, input.After, interactions)
}

func (r *userFeedResolver) User(ctx context.Context, obj *model.UserFeed) (*model.User, error) {
	return r.DB.GetUserByID(obj.UserID)
}

func (r *userFeedResolver) ContentItems(ctx context.Context, obj *model.UserFeed, input model.ContentItemConnectionInput) (*model.ContentItemConnection, error) {
	contentItems, err := r.FeedService.ServeContent(ctx, obj)
	if err != nil {
		return nil, err
	}
	//Set default value due to nil panic when 0 is passed to pager
	return page.BuildContentItemPage(input.First, input.After, contentItems)
}

func (r *userFeedResolver) Subscriptions(ctx context.Context, obj *model.UserFeed) ([]*model.FeedSubscription, error) {
	return r.DB.ListFeedSubscriptionsByFeedID(obj.ID)
}

func (r *userFeedResolver) Engine(ctx context.Context, obj *model.UserFeed) (*model.Engine, error) {
	return r.DB.GetEngineByID(obj.EngineID)
}

func (r *userFeedResolver) IsActive(ctx context.Context, obj *model.UserFeed) (bool, error) {
	user, err := r.DB.GetUserByID(obj.UserID)
	if err != nil {
		return false, err
	}
	if user.ActiveUserFeedID == obj.ID {
		return true, nil
	}
	return false, nil
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
	src, err := r.DB.GetSrcRSSFeedByID(obj.SrcRSSFeedID)
	if err != nil {
		return nil, err
	}
	return src, nil
}

// ContentItem returns generated.ContentItemResolver implementation.
func (r *Resolver) ContentItem() generated.ContentItemResolver { return &contentItemResolver{r} }

// Engine returns generated.EngineResolver implementation.
func (r *Resolver) Engine() generated.EngineResolver { return &engineResolver{r} }

// FeedSubscription returns generated.FeedSubscriptionResolver implementation.
func (r *Resolver) FeedSubscription() generated.FeedSubscriptionResolver {
	return &feedSubscriptionResolver{r}
}

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

// UserFeed returns generated.UserFeedResolver implementation.
func (r *Resolver) UserFeed() generated.UserFeedResolver { return &userFeedResolver{r} }

// UserSubscription returns generated.UserSubscriptionResolver implementation.
func (r *Resolver) UserSubscription() generated.UserSubscriptionResolver {
	return &userSubscriptionResolver{r}
}

type contentItemResolver struct{ *Resolver }
type engineResolver struct{ *Resolver }
type feedSubscriptionResolver struct{ *Resolver }
type interactionResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type srcRSSFeedResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type userFeedResolver struct{ *Resolver }
type userSubscriptionResolver struct{ *Resolver }
