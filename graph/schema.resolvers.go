package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/auth"
	"github.com/well-informed/wellinformed/graph/generated"
	"github.com/well-informed/wellinformed/graph/model"
)

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
	log.Debug("user: %v", user)

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

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	// TODO: add validation on input

	existingUser, err := r.DB.GetUserByEmail(input.Email)

	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		log.Printf("error while GetUserByEmail: %v", err)
		return nil, errors.New("email already in used")
	}

	existingUser, err = r.DB.GetUserByUsername(input.Username)

	if existingUser != nil {
		return nil, errors.New("username already in used")
	}

	user := &model.User{
		Username:  input.Username,
		Email:     input.Email,
		Firstname: input.Firstname,
		Lastname:  input.Lastname,
	}

	hashedPassword, err := auth.HashPassword(input.Password)
	if err != nil {
		log.Printf("error while hashing password: %v", err)
		return nil, errors.New("something went wrong")
	}

	user.Password = hashedPassword
	createdUser, err := r.DB.CreateUser(*user)

	if err != nil {
		log.Printf("error creating a user: %v", err)
		return nil, err
	}

	token, err := auth.GenAccessToken(user.ID)
	if err != nil {
		log.Printf("error while generating the token: %v", err)
		return nil, errors.New("something went wrong")
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      &createdUser,
	}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	// log.Printf("context: %v", ctx)
	existingUser, err := r.DB.GetUserByEmail(input.Email)
	log.Printf("existingUser: %v", existingUser)

	if existingUser == nil || err != nil {
		log.Printf("GetUserByEmail err: %v", err)
		return nil, errors.New("email/password combination don't work 1")
	}

	err = auth.ComparePassword(input.Password, existingUser.Password)
	if err != nil {
		log.Printf("ComparePassword err: %v", err)
		return nil, errors.New("email/password combination don't work 2")
	}

	accessToken, err := auth.GenAccessToken(existingUser.ID)
	// refreshToken, rerr := user.GenRefreshToken()
	if err != nil {
		return nil, errors.New("something went wrong")
	}

	return &model.AuthResponse{
		AuthToken: accessToken,
		User:      existingUser,
	}, nil
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

func (r *srcRSSFeedResolver) ContentItems(ctx context.Context, obj *model.SrcRSSFeed) ([]*model.ContentItem, error) {
	log.Debug("resolving ContentItems")
	contentItems, err := r.DB.ListContentItemsBySource(obj)
	if err != nil {
		return nil, err
	}
	return contentItems, nil
}

func (r *userResolver) SrcRSSFeeds(ctx context.Context, obj *model.User) ([]*model.SrcRSSFeed, error) {
	return r.DB.ListSrcRSSFeedsByUser(obj)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// SrcRSSFeed returns generated.SrcRSSFeedResolver implementation.
func (r *Resolver) SrcRSSFeed() generated.SrcRSSFeedResolver { return &srcRSSFeedResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type srcRSSFeedResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
