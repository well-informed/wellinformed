package database

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/well-informed/wellinformed/graph/model"
)

func TestUserSQL(t *testing.T) {
	db := NewDB()
	currentTime := time.Now()
	user := &model.User{
		Firstname:           "JohnJacob",
		Lastname:            "JingleHeimerSchmidt",
		Username:            "hisnameismyname",
		Email:               "longname@email.com",
		Password:            "asldkfjas;dfk",
		ActivePreferenceSet: "default",
		CreatedAt:           currentTime,
		UpdatedAt:           currentTime,
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
		t.Errorf("user: %+v", user)
		t.Errorf("byEmail: %+v", byEmail)
	}
	byID, err := db.GetUserById(returnedUser.ID)
	if err != nil {
		t.Error("error getting user by id. err: ", err)
	}
	if !cmp.Equal(user, byID, opt) {
		t.Error("byID did not match.")
		t.Errorf("user: %+v", user)
		t.Errorf("byID %+v", byID)
	}
	byUsername, err := db.GetUserByUsername("hisnameismyname")
	if err != nil {
		t.Error("error getting user by email. err: ", err)
	}
	if !cmp.Equal(user, byUsername, opt) {
		t.Error("byUsername did not match.")
		t.Errorf("user: %+v", user)
		t.Errorf("byUsername: %+v", byUsername)
	}
}
