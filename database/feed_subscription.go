package database

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/graph/model"
)

func (db DB) CreateFeedSubscription(feedID int64, sourceID int64, sourceType model.SourceType) (*model.FeedSubscription, error) {
	stmt := `INSERT INTO feed_subscriptions
	( feed_id,
		source_type,
		source_id,
		created_at,
		updated_at
	)
		VALUES($1,$2,$3,$4,$5)
		RETURNING id, created_at, updated_at`

	var ID int64
	var createdAt time.Time
	var updatedAt time.Time
	err := db.QueryRowx(stmt,
		feedID,
		sourceType,
		sourceID,
		time.Now(),
		time.Now(),
	).Scan(&ID, &createdAt, &updatedAt)
	if err != nil {
		log.Error("failed to create feed subscription. err: ", err)
		return nil, err
	}
	feedSubscription := &model.FeedSubscription{
		ID:         ID,
		UserFeedID: feedID,
		SourceID:   sourceID,
		SourceType: sourceType,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}
	return feedSubscription, nil
}

func (db DB) ListFeedSubscriptionsByFeedID(feedID int64) (feedSubscriptions []*model.FeedSubscription, err error) {
	stmt := `SELECT * FROM feed_subscriptions WHERE feed_id = $1`

	err = db.Select(&feedSubscriptions, stmt, feedID)
	if err != nil {
		log.Errorf("error listing feed subscriptions for feed ID %v. err: %v", feedID, err)
		return nil, err
	}
	return feedSubscriptions, nil
}
