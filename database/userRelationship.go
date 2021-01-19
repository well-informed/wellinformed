package database

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/graph/model"
)

func (db DB) SaveUserRelationship(followerID int64, followeeID int64) (*model.UserRelationship, error) {
	stmt := `INSERT INTO user_relationships
	(
		follower_id,
		followee_id,
		created_at,
		updated_at
	)
	VALUES($1, $2, $3, $4)
	ON CONFLICT (follower_id, followee_id)
	DO UPDATE SET
	follower_id = $1,
	followee_id = $2,
	updated_at = $3
	RETURNING id, created_at, updated_at`
	var ID int64
	var CreatedAt time.Time
	var UpdatedAt time.Time
	err := db.QueryRowx(stmt,
		followerID,
		followeeID,
		time.Now(),
		time.Now(),
	).Scan(&ID, &CreatedAt, &UpdatedAt)
	if err != nil {
		log.Error("failed to save interactions entry: err: ", err)
		return nil, err
	}
	userRelationship := &model.UserRelationship{
		ID:        ID,
		Follower:  followerID,
		Followee:  followeeID,
		CreatedAt: CreatedAt,
		UpdatedAt: UpdatedAt,
	}
	return userRelationship, nil
}

func (db DB) DeleteUserRelationship(followerID int64, followeeID int64) error {
	stmt := `DELETE FROM user_relationships WHERE follower_id = $1 AND followee_id = $2`
	_, err := db.Exec(stmt, followerID, followeeID)
	if err != nil {
		log.Error("unable to delete user relationship", err)
		return err
	}
	return nil
}

func (db DB) ListUserRelationshipsByFollowerID(followerID int64) ([]*model.UserRelationship, error) {
	stmt := `SELECT * FROM user_relationships WHERE follower_id = $1`

	userRelationships := make([]*model.UserRelationship, 0)
	err := db.Select(&userRelationships, stmt, followerID)
	if err != nil {
		log.Error("error listing relationships by follower ID", err)
		return nil, err
	}
	return userRelationships, nil
}

func (db DB) ListUserRelationshipsByFolloweeID(followeeID int64) ([]*model.UserRelationship, error) {
	stmt := `SELECT * FROM user_relationships WHERE followee_id = $1`

	userRelationships := make([]*model.UserRelationship, 0)
	err := db.Select(&userRelationships, stmt, followeeID)
	if err != nil {
		log.Error("error listing relationships by followee ID", err)
		return nil, err
	}
	return userRelationships, nil
}
