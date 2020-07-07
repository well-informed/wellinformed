package database

import (
	"database/sql"
	"errors"
	"strings"

	_ "github.com/lib/pq"
	"github.com/mmcdole/gofeed"
	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/graph/model"
)

type DB struct {
	*sql.DB //embeds the sql db methods on the DB struct
}

/*NewDB Creates a new handle on the database
and creates necessary tables if they do not already exist*/
func NewDB() DB {
	connStr := "user=postgres password=password dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("could not connect to database. err: ", err)
	}
	createTables(db)

	return DB{db}
}

//TODO: Write Insert, Delete, Update, and Read statements for all the entities in here.
//Use the graph models as input

/*Creates all necessary tables, either returns successfully,
or exits the program with call to log.Fatal()*/
func createTables(db *sql.DB) {
	createSrcRSSFeedsTable(db)
	// createMainFeedTable(db)
	createUsersTable(db)
	// createUserHistoryTable(db)
}

func createSrcRSSFeedsTable(db *sql.DB) {
	stmt := `
	CREATE TABLE IF NOT EXISTS src_rss_feeds
	(	id BIGSERIAL PRIMARY KEY,
		title varchar NOT NULL,
		description varchar,
		link varchar UNIQUE NOT NULL,
		feed_link varchar UNIQUE NOT NULL ,
		updated timestamp with time zone,
		last_fetched_at timestamp with time zone,
		language varchar,
		generator varchar
	)`
	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal("error creating rss_feeds table. err: ", err)
	}
}

func (db DB) InsertSrcRSSFeed(feed model.SrcRSSFeed) (model.SrcRSSFeed, error) {
	stmt, err := db.Prepare(`INSERT INTO src_rss_feeds
	( title,
		description,
		link,
		feed_link,
		updated,
		last_fetched_at,
		language,
		generator)
		values($1,$2,$3,$4,$5,$6,$7,$8)
		RETURNING id
		`)
	if err != nil {
		log.Error("failed to prepare src_rss_feeds insert: ", err)
		return feed, err
	}

	var id int64
	err = stmt.QueryRow(
		feed.Title,
		feed.Description,
		feed.Link,
		feed.FeedLink,
		feed.Updated,
		feed.LastFetchedAt,
		feed.Language,
		feed.Generator,
	).Scan(&id)
	if err != nil {
		log.Errorf("failed to insert row to src_rss_feeds. err: ", err)
		return feed, err
	}
	feed.ID = id
	log.Info("got id back: ", id)
	return feed, nil
}

func (db DB) getUserByField(selection string, whereClause string, args ...interface{}) (*model.User, error) {
	var user *model.User

	s := []string{"SELECT", selection, "FROM users WHERE", whereClause}
	stmt := strings.Join(s, " ")

	err := db.QueryRow(stmt, args...).Scan(
		&user.ID,
		&user.Firstname,
		&user.Lastname,
		&user.Username,
		&user.Email,
		&user.Password,
	)

	return user, err
}

func (db DB) GetUserByEmail(value string) (*model.User, error) {
	return db.getUserByField("*", "email = $1", value)
}

func (db DB) GetUserByUsername(value string) (*model.User, error) {
	return db.getUserByField("*", "user_name = $1", value)
}

func (db DB) GetUserById(value string) (*model.User, error) {
	return db.getUserByField("*", "id = $1", value)
}

func (db DB) CreateUser(user model.User) (model.User, error) {
	stmt, err := db.Prepare(`INSERT INTO users
	( email,
		first_name,
		last_name,
		user_name,
		password)
		values($1,$2,$3,$4,$5)
		RETURNING id
		`)
	if err != nil {
		log.Error("failed to prepare user insert: ", err)
		return user, err
	}

	var ID int64
	err = stmt.QueryRow(
		user.Email,
		user.Firstname,
		user.Lastname,
		user.Username,
		user.Password,
	).Scan(&ID)
	if err != nil {
		log.Errorf("failed to insert row to create user. err: ", err)
		return user, err
	}
	user.ID = ID
	log.Info("got id back: ", ID)
	return user, nil
}

func (db DB) SelectSrcRSSFeed(input model.SrcRSSFeedInput) (model.SrcRSSFeed, error) {
	var feed model.SrcRSSFeed
	var whereClause string
	var arg interface{}

	if input.ID != nil {
		whereClause = `WHERE id = $1`
		arg = *input.ID
	} else if input.Link != nil {
		whereClause = `WHERE link = $1`
		arg = *input.Link
	} else if input.FeedLink != nil {
		whereClause = `WHERE feed_link = $1`
		arg = *input.FeedLink
	} else {
		return feed, errors.New("no key for select found")
	}
	stmt := `SELECT * FROM src_rss_feeds ` + whereClause
	err := db.QueryRow(stmt, arg).Scan(
		&feed.ID,
		&feed.Title,
		&feed.Description,
		&feed.Link,
		&feed.FeedLink,
		&feed.Updated,
		&feed.LastFetchedAt,
		&feed.Language,
		&feed.Generator)
	log.Debugf("Selected feed with whereClause %v with key %v: %v", whereClause, arg, feed)
	return feed, err
}

func createMainFeedTable(db *sql.DB) {
	stmt := `
	CREATE TABLE IF NOT EXISTS main_feed
 ( title varchar,
	 description varchar,
	 content varchar,
	 link varchar,
	 updated timestamp with time zone,
	 published timestamp with time zone,
	 author varchar,
	 parent_feed varchar NOT NULL,
	 guid varchar NOT NULL,
	 FOREIGN KEY (parent_feed) REFERENCES src_rss_feeds(link),
	 PRIMARY KEY (parent_feed,guid)
 )`
	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal("error creating rss_articles table. err: ", err)
	}
}

func insertItem(db *sql.DB, feedLink string, article *gofeed.Item) error {
	stmt, err := db.Prepare(`INSERT INTO main_feed
	(title,
	description,
	htmlContent,
	link,
	updated,
	published,
	author,
	guid,
	parent_feed)
	values($1,$2,$3,$4,$5,$6,$7,$8,$9)`)
	if err != nil {
		log.Error("failed to prepare rss_articles insert", err)
		return err
	}

	_, err = stmt.Exec(
		article.Title,
		article.Description,
		article.Content,
		article.Link,
		article.UpdatedParsed,
		article.PublishedParsed,
		article.Author.Name,
		article.GUID,
		feedLink)
	if err != nil {
		log.Errorf("failed to exec insert of content %+v. err: %v", article, err)
		return err
	}
	return nil
}

func createUsersTable(db *sql.DB) {
	stmt := `
	CREATE TABLE IF NOT EXISTS users
	( id BIGSERIAL PRIMARY KEY,
		email varchar,
		first_name varchar,
		last_name varchar,
		user_name varchar,
		password varchar)`

	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal("error creating users table. err: ", err)
	}
}

func createUserHistoryTable(db *sql.DB) {
	stmt := `
	CREATE TABLE IF NOT EXISTS user_history
	( userID varchar,
		parent_feed varchar,
		guid varchar,
		trustworthiness smallint,
		insightfulness smallint,
		entertainment smallint,
		importance smallint,
		overall smallint,
		notes text,
		FOREIGN KEY (userID) REFERENCES users(userID),
		FOREIGN KEY (parent_feed, guid) REFERENCES main_feed(parent_feed, guid),
		PRIMARY KEY (userID, parent_feed, guid)
		)`
	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal("error creating history table. err: ", err)
	}
}

func createUserPrefSetTable(db *sql.DB) {
	stmt := `
	CREATE TABLE IF NOT EXISTS preference_sets
	( userID varchar,
		pref_set_name varchar,
		FOREIGN KEY (userID) REFERENCES users(userID),
		PRIMARY KEY (userID, pref_set_name)
	)`
	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal("error creating preference_sets table. err: ", err)
	}
}

func createUserSourcesTable(db *sql.DB) {
	stmt := `
	CREATE TABLE IF NOT EXISTS user_sources
	( userID varchar,
		pref_set_name varchar,
		source varchar,
		FOREIGN KEY (userID, pref_set_name) REFERENCES preference_sets(userID, pref_set_name),
		FOREIGN KEY (source) REFERENCES src_rss_feeds(link)
		PRIMARY KEY (userID, pref_set_name, source)
	)`
	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal("error creating user_sources table. err: ", err)
	}
}
