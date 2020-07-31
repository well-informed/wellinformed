package database

import (
	"database/sql"

	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/graph/model"
)

func (db DB) SaveHistory(userID int64, input *model.HistoryInput) (*model.History, error) {
	stmt := `INSERT INTO history
	( user_id,
		content_item_id,
		read_state,
		percent_read
	)
	VALUES($1, $2, $3, $4)
	RETURNING id
	ON CONFLICT DO UPDATE SET
	user_id = $1,
	content_item_id = $2,
	read_state = $3,
	percent_read = $4
	RETURNING id`
	var ID int64
	err := db.QueryRowx(stmt, userID, input).Scan(&ID)
	if err != nil {
		log.Error("failed to save history entry: err: ", err)
		return nil, err
	}
	return &model.History{
		ID:            ID,
		UserID:        userID,
		ContentItemID: input.ContentItemID,
		ReadState:     input.ReadState,
		PercentRead:   input.PercentRead,
	}, nil
}

func (db DB) GetHistoryByContentID(userID int64, contentItemID int64) (*model.History, error) {
	var itemHistory model.History
	err := db.Get(&itemHistory,
		`SELECT * FROM history WHERE user_id = $1 AND content_item_id = $2`, userID, contentItemID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Error("failed to select History by user_id and content_item_id. err: ", err)
		return nil, err
	}
	return &itemHistory, nil
}

func (db DB) ListUserHistory(userID int64) ([]*model.History, error) {
	return db.listUserHistoryByQuery(`SELECT * FROM history WHERE user_id = $1`, userID)
}

func (db DB) listUserHistoryByQuery(stmt string, args ...interface{}) ([]*model.History, error) {
	histories := make([]*model.History, 0)
	err := db.Select(&histories, stmt, args)
	if err != nil {
		log.Error("error selecting all Histories. err: ", err)
		return nil, err
	}
	return histories, nil
}

