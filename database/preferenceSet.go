package database

import (
	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/graph/model"
)

func (db DB) CreatePreferenceSet(prefSet *model.PreferenceSet) (*model.PreferenceSet, error) {
	log.Debugf("prefSet: %+v", prefSet)
	stmt, err := db.Prepare(`INSERT into preference_sets
	( user_id,
		name,
		sort)
		VALUES($1,$2,$3)
		RETURNING id
		`)
	if err != nil {
		log.Error("failed to prepare preference_sets insert. err: ", err)
		return prefSet, err
	}
	var ID int64
	err = stmt.QueryRow(
		prefSet.UserID,
		prefSet.Name,
		prefSet.Sort,
	).Scan(&ID)
	if err != nil {
		log.Errorf("failed to insert preference set. err: ", err)
		return prefSet, err
	}
	prefSet.ID = ID
	return prefSet, nil
}

func (db DB) GetPreferenceSetByID(id int64) (*model.PreferenceSet, error) {
	var prefSet *model.PreferenceSet
	err := db.Get(prefSet, "SELECT * FROM preference_sets WHERE id = $1", id)
	if err != nil {
		log.Errorf("failed to get preferenceSet by id. err: ", err)
		return nil, err
	}
	return prefSet, nil
}

func (db DB) ListPreferenceSetsByUser(userID int64) ([]*model.PreferenceSet, error) {
	prefSets := make([]*model.PreferenceSet, 0)
	err := db.Select(prefSets, "SELECT * FROM preference_sets WHERE user_id = $1", userID)
	if err != nil {
		log.Errorf("failed to select preference sets by user_id. err: ", err)
		return nil, err
	}
	return prefSets, nil
}

func (db DB) GetPreferenceSetByName(userID int64, name string) (*model.PreferenceSet, error) {
	var prefSet model.PreferenceSet
	err := db.Get(&prefSet, "SELECT * FROM preference_sets WHERE user_id = $1 AND name = $2", userID, name)
	if err != nil {
		log.Error("failed to select preferenceSet by name. err: ", err)
		return nil, err
	}
	return &prefSet, nil
}
