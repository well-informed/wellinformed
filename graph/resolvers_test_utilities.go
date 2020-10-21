package graph

import (
	"context"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"
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
	//Postgres can't accept capitalized database names, so lowercase them here for convenience
	dbName = strings.ToLower(dbName)
	conf := wellinformed.Config{
		ServerPort: "8081",
		DBHost:     "localhost",
		DBName:     dbName, //dbname cannot be capitalized
		DBUser:     "postgres",
		DBPassword: "password",
		LogLevel:   log.InfoLevel,
	}
	log.SetLevel(conf.LogLevel)
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

func NewMockAuthenticatedContext(db wellinformed.Persistor, id int64) (user *model.User, ctx context.Context) {
	//Mimic middleware authentication by setting user to key in context.
	user, err := db.GetUserByID(id)
	if err != nil {
		log.Error("could not fetch registered user for mock authentication middleware. err: ", err)
	}
	return user, context.WithValue(context.Background(), auth.CurrentUserKey, user)
}

//Sets up a new test harness with a unique Database and a registered and authenticated user
func NewAuthedUserTestEnv(testName string) (resolver *Resolver, user *model.User, ctx context.Context) {
	resolver, _ = NewTestHarness(testName)
	authResponse, err := RegisterMutation(resolver, nil)
	if err != nil {
		log.Fatal("could not register. err: ", err)
	}
	user, ctx = NewMockAuthenticatedContext(resolver.DB, authResponse.User.ID)
	return resolver, user, ctx
}

//Standard ways to call mutations to use as test building blocks

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

func AddStockSrcRSSFeeds(resolver *Resolver, ctx context.Context) (srcRSSFeeds []*model.SrcRSSFeed, errs []error) {
	testSrcRSSFeeds := []string{
		"https://bankless.substack.com/feed",
		"https://www.buzzfeed.com/index.xml",
		"https://waitbutwhy.com/feed",
		"https://www.overcomingbias.com/feed",
		"https://www.lesswrong.com/feed.xml?view=curated-rss&karmaThreshold=2",
	}
	for _, v := range testSrcRSSFeeds {
		srcRSSFeed, err := resolver.Mutation().AddSrcRSSFeed(ctx, v, 1)
		srcRSSFeeds = append(srcRSSFeeds, srcRSSFeed)
		errs = append(errs, err)
	}
	return srcRSSFeeds, errs
}
