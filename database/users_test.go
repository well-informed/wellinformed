package database

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed"
	"github.com/well-informed/wellinformed/graph/model"
)

//Drops database and recreates it for testing
func NewTestDB(conf wellinformed.Config) DB {
	/*NewDB Creates a new handle on the database
	and creates necessary tables if they do not already exist*/
	format := "postgres://%v:%v@%v:5432/%v?sslmode=disable"
	connStr := fmt.Sprintf(format, conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBName)
	defaultDB, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("could not connect to database. err: ", err)
	}
	_, err = defaultDB.Queryx("DROP DATABASE unittest")
	if err != nil {
		log.Error("could not drop unittest database. err:", err)
	}
	_, err = defaultDB.Queryx("CREATE DATABASE unittest")
	if err != nil {
		log.Error("could not create unittest database. err: ", err)
	}
	connStr = fmt.Sprintf(format, conf.DBUser, conf.DBPassword, conf.DBHost, "unittest")
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Error("couldn't connect to unittest db. err: ", err)
	}
	MigrateSchema(connStr)
	return DB{db}
}

func TestUserSQL(t *testing.T) {
	db := NewTestDB(wellinformed.GetConfig())
	currentTime := time.Now()
	user := &model.User{
		Firstname:        "JohnJacob",
		Lastname:         "JingleHeimerSchmidt",
		Username:         "hisnameismyname",
		Email:            "longname@email.com",
		Password:         "asldkfjas;dfk",
		ActiveUserFeedID: 1,
		CreatedAt:        currentTime,
		UpdatedAt:        currentTime,
	}
	opt := cmpopts.IgnoreFields(model.User{}, "ID", "CreatedAt", "UpdatedAt")

	returnedUser, err := db.CreateUser(*user)
	if err != nil {
		t.Error("error creating user: ", err)
	}
	if returnedUser.ID == 0 {
		t.Error("returned user should have had an ID")
	}
	byEmail, err := db.GetUserByEmail("longname@email.com")
	if err != nil {
		t.Error("error getting user by email. err: ", err)
	}
	if !cmp.Equal(user, byEmail, opt) {
		t.Error("byEmail did not match.")
		t.Errorf("correct: %+v", user)
		t.Errorf("byEmail: %+v", byEmail)
	}
	byID, err := db.GetUserByID(returnedUser.ID)
	if err != nil {
		t.Error("error getting user by id. err: ", err)
	}
	if !cmp.Equal(user, byID, opt) {
		t.Error("byID did not match.")
		t.Errorf("correct: %+v", user)
		t.Errorf("byID %+v", byID)
	}
	byUsername, err := db.GetUserByUsername("hisnameismyname")
	if err != nil {
		t.Error("error getting user by email. err: ", err)
	}
	if !cmp.Equal(user, byUsername, opt) {
		t.Error("byUsername did not match.")
		t.Errorf("correct: %+v", user)
		t.Errorf("byUsername: %+v", byUsername)
	}
}
