package database

import (
	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/graph/model"
)

//Creates a new of updates an existing Engine, attaching the id found in the database
func (db DB) SaveEngine(engine *model.Engine) (*model.Engine, error) {
	log.Debugf("engine: %+v", engine)
	stmt, err := db.Prepare(`INSERT into engines
	( user_id,
		name,
		sort,
		start_date,
		end_date)
		VALUES($1,$2,$3,$4,$5)
		ON CONFLICT (user_id, name)
		DO UPDATE SET
			user_id = $1,
			name = $2,
			sort = $3,
			start_date = $4,
			end_date = $5
		RETURNING id
		`)
	if err != nil {
		log.Error("failed to prepare engines insert. err: ", err)
		return engine, err
	}
	var ID int64
	err = stmt.QueryRow(
		engine.UserID,
		engine.Name,
		engine.Sort,
		engine.StartDate,
		engine.EndDate,
	).Scan(&ID)
	if err != nil {
		log.Error("failed to insert preference set. err: ", err)
		return engine, err
	}
	engine.ID = ID
	return engine, nil
}

func (db DB) GetEngineByID(id int64) (*model.Engine, error) {
	var engine model.Engine
	err := db.Get(&engine, "SELECT * FROM engines WHERE id = $1", id)
	if err != nil {
		log.Error("failed to get Engine by id. err: ", err)
		return nil, err
	}
	return &engine, nil
}

func (db DB) ListEnginesByUser(userID int64) ([]*model.Engine, error) {
	engines := make([]*model.Engine, 0)
	err := db.Select(&engines, "SELECT * FROM engines WHERE user_id = $1", userID)
	if err != nil {
		log.Error("failed to select preference sets by user_id. err: ", err)
		return nil, err
	}
	return engines, nil
}

func (db DB) GetEngineByName(userID int64, name string) (*model.Engine, error) {
	var engine model.Engine
	err := db.Get(&engine, "SELECT * FROM engines WHERE user_id = $1 AND name = $2", userID, name)
	if err != nil {
		log.Error("failed to select Engine by name. err: ", err)
		return nil, err
	}
	return &engine, nil
}
