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

func (db DB) GetHistoryByContentID(int64, int64) (*model.History, error) {
	log.Panic("not implemented")
	return nil, nil
}

func (db DB) ListUserHistory(int64) ([]*model.History, error) {
	log.Panic("not implemented")
	return nil, nil
}
