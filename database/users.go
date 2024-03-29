package database

import (
	"database/sql"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/graph/model"
)

func (db DB) getUserByField(selection string, whereClause string, args ...interface{}) (*model.User, error) {
	var user model.User

	s := []string{"SELECT", selection, "FROM users WHERE", whereClause}
	stmt := strings.Join(s, " ")

	err := db.Get(&user, stmt, args...)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &user, err
}

func (db DB) GetUserByEmail(value string) (*model.User, error) {
	return db.getUserByField("*", "email = $1", value)
}

func (db DB) GetUserByUsername(value string) (*model.User, error) {
	return db.getUserByField("*", "user_name = $1", value)
}

func (db DB) GetUserByID(value int64) (*model.User, error) {
	return db.getUserByField("*", "id = $1", value)
}

func (db DB) CreateUser(user *model.User) (*model.User, error) {
	stmt, err := db.Prepare(`INSERT INTO users
	( email,
		first_name,
		last_name,
		user_name,
		password,
		active_user_feed,
		created_at,
		updated_at)
		values($1,$2,$3,$4,$5,$6,$7,$8)
		RETURNING id
		`)
	if err != nil {
		log.Error("failed to prepare user insert: ", err)
		return nil, err
	}

	var ID int64
	err = stmt.QueryRow(
		user.Email,
		user.Firstname,
		user.Lastname,
		user.Username,
		user.Password,
		user.ActiveUserFeedID,
		time.Now(),
		time.Now(),
	).Scan(&ID)
	if err != nil {
		log.Error("failed to insert row to create user. err: ", err)
		return nil, err
	}
	user.ID = ID
	return user, nil
}

func (db DB) UpdateUser(user *model.User) (*model.User, error) {
	stmt := `UPDATE users SET
	email = $1,
	first_name = $2,
	last_name = $3,
	user_name = $4,
	password = $5,
	active_user_feed = $6,
	updated_at = $7
	WHERE id = $8
	RETURNING updated_at`

	err := db.QueryRowx(stmt,
		user.Email,
		user.Firstname,
		user.Lastname,
		user.Username,
		user.Password,
		user.ActiveUserFeedID,
		time.Now(),
		user.ID).Scan(&user.UpdatedAt)
	if err != nil {
		log.Error("failed to update user: err: ", err)
		return nil, err
	}
	return user, nil
}

func (db DB) GetUserByInteraction(interactionId int64) (*model.User, error) {
	var user model.User

	stmt := `
		SELECT u.* FROM users u
		INNER JOIN interactions i on u.id = i.user_id
		WHERE i.id = $1
		LIMIT 1
	`

	err := db.QueryRow(stmt, interactionId).Scan(
		&user.ID,
		&user.Email,
		&user.Firstname,
		&user.Lastname,
		&user.Username,
		&user.Password,
		&user.ActiveUserFeedID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		log.Error("failed to GetUserByInteraction: err: ", err)
		return nil, err
	}

	return &user, nil
}
