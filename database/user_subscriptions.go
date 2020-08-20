package database

import (
	"database/sql"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/graph/model"
)

func (db DB) InsertUserSubscription(user model.User, src model.SrcRSSFeed) (subscription *model.UserSubscription, err error) {
	subscription = &model.UserSubscription{}
	stmt, err := db.Prepare(`INSERT INTO user_subscriptions
	( user_id,
		source_id,
		created_at)
		VALUES($1,$2,$3)
		RETURNING id`)
	if err != nil {
		log.Error("failed to prepare user_subscriptions insert", err)
		return nil, err
	}
	var id int64
	err = stmt.QueryRow(
		user.ID,
		src.ID,
		time.Now(),
	).Scan(&id)
	if err != nil {
		log.Error("failed to insert row to user_subscriptions. err: ", err)
		return nil, err
	}
	subscription.ID = id
	return subscription, err
}

func (db DB) GetUserSubscription(userID int64, srcID int64) (*model.UserSubscription, error) {
	var userSub model.UserSubscription

	stmt := `SELECT * FROM user_subscriptions WHERE user_id = $1 AND source_id = $2`
	err := db.QueryRow(stmt, userID, srcID).Scan(
		&userSub.ID,
		&userSub.UserID,
		&userSub.SrcRSSFeedID,
		&userSub.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &userSub, err
}

func (db DB) DeleteUserSubscription(userID int64, srcID int64) (int, error) {
	stmt := `DELETE FROM user_subscriptions WHERE user_id = $1 AND source_id = $2`
	result, err := db.Exec(stmt, userID, srcID)
	if err != nil {
		log.Error("unable to delete user subscription", err)
		return 0, err
	}
	numDeleted, err := result.RowsAffected()
	if err != nil {
		log.Error("error getting rows affected by user subscription deletion. err: ", err)
	}
	return int(numDeleted), err
}

// func (db DB) PageUserSubscriptions(userID int64, input *model.ConnectionInput) (*model.Connection, error) {
// 	stmt := `SELECT * FROM user_subscriptions WHERE user_id = $1`

// 	userSubscriptions := make([]*model.UserSubscription, 0)
// 	err := db.Select(&userSubscriptions, stmt, userID)
// 	if err != nil {
// 		log.Error("error listing subscriptions for user. err: ", err)
// 		return nil, err
// 	}
// 	return page.BuildPage(input.First, input.After, page.UserSubscriptionsToNodes(userSubscriptions))
// }
