package database

import (
	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/graph/model"
)

func (db DB) SaveHistory(user_id int64, input *model.HistoryInput) (*model.History, error) {
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
	err := db.QueryRowx(stmt, user_id, input).Scan(&ID)
	if err != nil {
		log.Error("failed to save history entry: err: ", err)
		return nil, err
	}
	return &model.History{
		ID:            ID,
		UserID:        user_id,
		ContentItemID: input.ContentItemID,
		ReadState:     input.ReadState,
		PercentRead:   input.PercentRead,
	}, nil
}

func (db DB) GetHistoryByContentID(user_id int64, content_item_id int64) (*model.History, error) {
	var itemHistory model.History
	err := db.Get(&itemHistory,
		`SELECT * FROM history WHERE user_id = $1 AND content_item_id = $2`, user_id, content_item_id)
	if err != nil {
		log.Error("failed to select History by user_id and content_item_id. err: ", err)
		return nil, err
	}
	return &itemHistory, nil
}

func (db DB) ListUserHistory(user_id int64) ([]*model.History, error) {
	return db.ListUserHistoryByQuery(`SELECT * FROM history WHERE user_id = $1`, user_id)
}

func (db DB) ListUserHistoryByQuery(stmt string, args ...interface{}) ([]*model.History, error) {
	rows, err := db.Query(stmt, args...)
	defer rows.Close()
	if err != nil {
		log.Error("error selecting all Histories. err: ", err)
		return nil, err
	}
	histories := make([]*model.History, 0)
	for rows.Next() {
		var hist model.History
		err := rows.Scan(
			&hist.ID,
			&hist.UserID,
			&hist.ContentItemID,
			&hist.ReadState,
			&hist.PercentRead,
		)
		if err != nil {
			log.Error("error scanning History row: err: ", err)
		}
		histories = append(histories, &hist)
	}
	if err := rows.Err(); err != nil {
		log.Error("error listing Histories. err: ", err)
		return nil, err
	}
	return histories, nil
}

