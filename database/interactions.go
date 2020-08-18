package database

import (
	"database/sql"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/graph/model"
	"github.com/well-informed/wellinformed/pagination"
)

func (db DB) SaveInteraction(userID int64, input *model.InteractionInput) (*model.ContentItem, error) {
	stmt := `INSERT INTO interactions
	( user_id,
		content_item_id,
		read_state,
		percent_read,
		created_at,
		updated_at
	)
	VALUES($1, $2, $3, $4, $5, $6)
	ON CONFLICT (user_id, content_item_id)
	DO UPDATE SET
	user_id = $1,
	content_item_id = $2,
	read_state = $3,
	percent_read = $4,
	updated_at = $6
	RETURNING id, created_at, updated_at`
	var ID int64
	var CreatedAt time.Time
	var UpdatedAt time.Time
	err := db.QueryRowx(stmt,
		userID,
		input.ContentItemID,
		input.ReadState,
		input.PercentRead,
		time.Now(),
		time.Now(),
	).Scan(&ID, &CreatedAt, &UpdatedAt)
	if err != nil {
		log.Error("failed to save interactions entry: err: ", err)
		return nil, err
	}
	return db.GetContentItem(input.ContentItemID)
}

func (db DB) GetInteractionByContentID(userID int64, contentItemID int64) (*model.Interaction, error) {
	var itemInteraction model.Interaction
	err := db.Get(&itemInteraction,
		`SELECT * FROM interactions WHERE user_id = $1 AND content_item_id = $2`, userID, contentItemID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Error("failed to select interactions by user_id and content_item_id. err: ", err)
		return nil, err
	}
	return &itemInteraction, nil
}

func interactionsToNodes(interactions []*model.Interaction) []*model.Node {
	nodes := make([]*model.Node, 0)
	for _, interaction := range interactions {
		nodes = append(nodes, &model.Node{
			Value: interaction,
			ID:    interaction.ID,
		})
	}
	return nodes
}

func (db DB) PageUserInteractions(userID int64, readState *model.ReadState, input *model.ConnectionInput) (*model.Connection, error) {
	var interactions []*model.Interaction
	var err error
	if readState == nil {
		interactions, err = db.listUserInteractionsByQuery(`SELECT * FROM interactions WHERE user_id = $1`, userID)
	} else {
		interactions, err = db.listUserInteractionsByQuery(`SELECT * FROM interactions WHERE user_id = $1 AND read_state = $2`, userID, readState)
	}
	if err != nil {
		log.Error("error selecting list of interactions. err: ", err)
	}
	return pagination.BuildPage(input.First, input.After, interactionsToNodes(interactions))
}

func (db DB) listUserInteractionsByQuery(stmt string, args ...interface{}) ([]*model.Interaction, error) {
	interactions := make([]*model.Interaction, 0)
	err := db.Select(&interactions, stmt, args...)
	if err != nil {
		log.Error("error selecting interactions. err: ", err)
		return nil, err
	}
	return interactions, nil
}
