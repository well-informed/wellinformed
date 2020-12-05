package database

import (
	"errors"

	"github.com/well-informed/wellinformed/graph/model"
)

func (db DB) SaveUserRelationship(followerID int64, followeeID int64) (*model.UserRelationship, error) {
	return nil, errors.New("not implemented")
}
func (db DB) DeleteUserRelationship(followerID int64, followeeID int64) error {
	return errors.New("not implemented")
}
func (db DB) ListUserRelationshipsByFollowerID(followerID int64) ([]*model.UserRelationship, error) {
	return nil, errors.New("not implemented")
}
func (db DB) ListUserRelationshipsByFolloweeID(followeeID int64) ([]*model.UserRelationship, error) {
	return nil, errors.New("not implemented")
}
