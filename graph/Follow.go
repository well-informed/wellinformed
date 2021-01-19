package graph

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/graph/model"
)

func (r *mutationResolver) followUser(ctx context.Context, user *model.User, input model.UserRelationshipInput) (*model.UserRelationship, error) {
	if user.ID != input.FollowerID {
		return nil, errors.New("follower must be signed in user")
	}
	return r.DB.SaveUserRelationship(input.FollowerID, input.FolloweeID)
}

func (r *mutationResolver) unfollowUser(ctx context.Context, user *model.User, input model.UserRelationshipInput) (*model.DeleteResponse, error) {
	if user.ID != input.FollowerID {
		return &model.DeleteResponse{Ok: false}, errors.New("follower must be signed in user")
	}
	err := r.DB.DeleteUserRelationship(input.FollowerID, input.FolloweeID)
	if err != nil {
		log.Error("unable to delete user relationship")
		return &model.DeleteResponse{Ok: false}, err
	}
	return &model.DeleteResponse{Ok: true}, nil
}
