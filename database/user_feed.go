package database

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/graph/model"
)

func (db DB) CreateUserFeed(userFeed *model.UserFeed) (*model.UserFeed, error) {
	stmt, err := db.Prepare(`INSERT INTO user_feeds
	( user_id,
		name,
		engine_id,
		created_at,
		updated_at
	) values($1,$2,$3,$4,$5)
		RETURNING id
		`)
	if err != nil {
		log.Error("failed to prepare user_feed insert: ", err)
		return userFeed, err
	}

	var ID int64
	err = stmt.QueryRow(
		userFeed.UserID,
		userFeed.Name,
		userFeed.EngineID,
		time.Now(),
		time.Now(),
	).Scan(&ID)
	if err != nil {
		log.Error("failed to insert row to create user_feed. err: ", err)
		return userFeed, err
	}
	userFeed.ID = ID
	return userFeed, nil
}

func (db DB) GetUserFeedByID(id int64) (*model.UserFeed, error) {
	var userFeed model.UserFeed
	err := db.Get(&userFeed, "SELECT * FROM user_feeds WHERE id = $1", id)
	if err != nil {
		log.Errorf("failed to get user feed by ID: %v err: %v", id, err)
		return nil, err
	}
	return &userFeed, nil
}

func (db DB) ListUserFeedsByUser(userID int64) ([]*model.UserFeed, error) {
	userFeeds := make([]*model.UserFeed, 0)
	err := db.Select(&userFeeds, "SELECT * FROM user_feeds WHERE user_id = $1", userID)
	if err != nil {
		log.Error("failed to select user feeds by user_id. err: ", err)
		return nil, err
	}
	return userFeeds, nil
}
