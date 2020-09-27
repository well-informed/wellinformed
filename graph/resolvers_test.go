package graph

import (
	"context"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed"
	"github.com/well-informed/wellinformed/auth"
	"github.com/well-informed/wellinformed/database"
	"github.com/well-informed/wellinformed/feed"
	"github.com/well-informed/wellinformed/graph/model"
	"github.com/well-informed/wellinformed/rss"
	"github.com/well-informed/wellinformed/subscriber"
	"github.com/well-informed/wellinformed/user"
)

//Creates a new test harness, including a unique Database for each test that request it.
//This allows the tests to run in parallel against their own datasets without conflicting with each other.
//Exercises and depends upon basically all of the real modules except for the server.
//Namely the subscriber service, database, user service, and feed service
func NewTestHarness(dbName string) (*Resolver, *sqlx.DB) {
	conf := wellinformed.Config{
		ServerPort: "8081",
		DBHost:     "localhost",
		DBName:     dbName, //dbname cannot be capitalized
		DBUser:     "postgres",
		DBPassword: "password",
		LogLevel:   log.FatalLevel,
	}
	db, sqlxDB := NewTestDB(conf)
	rssHandler := rss.NewRSS()
	sub, err := subscriber.NewSubscriber(rssHandler, db)
	if err != nil {
		log.Fatal("could not initialize test subscriber. err: ", err)
	}
	resolver := &Resolver{
		DB:          db,
		RSS:         rssHandler,
		Sub:         sub,
		FeedService: feed.NewFeedService(db),
		UserService: user.NewUserService(db),
	}

	return resolver, sqlxDB
}

func NewTestDB(conf wellinformed.Config) (database.DB, *sqlx.DB) {
	//Connect to default postgres Database
	format := "postgres://%v:%v@%v:5432/%v?sslmode=disable"
	connStr := fmt.Sprintf(format, conf.DBUser, conf.DBPassword, conf.DBHost, "postgres")
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("could not connect to database. err: ", err)
	}
	db.DB.SetMaxOpenConns(conf.DBMaxOpenConnections)
	db.DB.SetMaxIdleConns(conf.DBMaxIdleConnections)

	//With postgres connection, recreate custom test database
	dropDB := fmt.Sprintf("DROP DATABASE IF EXISTS %v", conf.DBName)
	createDB := fmt.Sprintf("CREATE DATABASE %v", conf.DBName)
	_, err = db.Queryx(dropDB)
	if err != nil {
		log.Fatal("could not drop database. err: ", err)
	}
	_, err = db.Queryx(createDB)
	if err != nil {
		log.Fatal("could not create db. err: ", err)
	}

	//Migrate schema and connect to test database
	testConnStr := fmt.Sprintf(format, conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBName)
	db, err = sqlx.Connect("postgres", testConnStr)
	if err != nil {
		log.Fatal("could not connect to test database. err: ", err)
	}
	database.MigrateSchema(testConnStr)
	return database.DB{DB: db}, db
}

func NewMockAuthMiddlewareContext(db wellinformed.Persistor, id int64) context.Context {
	//Mimic middleware authentication by setting user to key in context.
	user, err := db.GetUserByID(id)
	if err != nil {
		log.Error("could not fetch registered user for mock authentication middleware. err: ", err)
	}
	return context.WithValue(context.Background(), auth.CurrentUserKey, user)
}

func RegisterMutation(resolver *Resolver, input *model.RegisterInput) (*model.AuthResponse, error) {
	if input == nil {
		input = &model.RegisterInput{
			Username:        "deviator",
			Email:           "danielveenstra@protonmail.com",
			Password:        "ScoobyDoo69",
			ConfirmPassword: "ScoobyDoo69",
			Firstname:       "Dan",
			Lastname:        "Veenstra",
		}
	}

	return resolver.Mutation().Register(context.Background(), *input)
}

func LoginMutation(resolver *Resolver, input model.LoginInput) (*model.AuthResponse, error) {
	return resolver.Mutation().Login(context.Background(), input)
}

func TestRegister(t *testing.T) {
	resolver, _ := NewTestHarness("register")
	_, err := RegisterMutation(resolver, nil)
	if err != nil {
		t.Error("failed to register. err: ", err)
	}
}

func TestRegisterAndLogin(t *testing.T) {
	resolver, _ := NewTestHarness("register_and_login")
	email := "danielveenstra@protonmail.com"
	password := "ScoobyDoo69"
	registerInput := &model.RegisterInput{
		Username:        "deviator",
		Email:           email,
		Password:        password,
		ConfirmPassword: password,
		Firstname:       "Dan",
		Lastname:        "Veenstra",
	}
	loginInput := model.LoginInput{
		Email:    email,
		Password: password,
	}
	_, err := RegisterMutation(resolver, registerInput)
	if err != nil {
		t.Error("couldn't register. err: ", err)
	}
	_, err = resolver.Mutation().Login(context.Background(), loginInput)
	if err != nil {
		t.Error("couldn't log in. err: ", err)
	}
}

func TestAddSrcRSSFeed(t *testing.T) {
	resolver, _ := NewTestHarness("add_src_rss_feed")
	authResponse, err := RegisterMutation(resolver, nil)
	if err != nil {
		t.Error("could not register. err: ", err)
	}

	ctx := NewMockAuthMiddlewareContext(resolver.DB, authResponse.User.ID)
	input := "https://bankless.substack.com/feed"
	srcRSSFeed, err := resolver.Mutation().AddSrcRSSFeed(ctx, input)
	if err != nil {
		t.Fatal("could not add srcRSSFeed. err: ", err)
	}
	if srcRSSFeed.ID == 0 {
		t.Errorf("missing ID")
	}
	if srcRSSFeed.Title == "" {
		t.Errorf("missing title")
	}
}

// func TestAddUserFeed(t *testing.T) {
// 	_, resolver, _ := NewTestHarness("add_source")
// }
