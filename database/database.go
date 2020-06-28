package database

import (
	"database/sql"

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
	createMainFeedTable(db)
	createUsersTable(db)
	createUserHistoryTable(db)
}

func createSrcRSSFeedsTable(db *sql.DB) {
	stmt := `
	CREATE TABLE IF NOT EXISTS src_rss_feeds
	(	title varchar NOT NULL,
		description varchar,
		link varchar NOT NULL,
		feed_link varchar NOT NULL ,
		updated timestamp with time zone,
		published timestamp with time zone,
		language varchar,
		generator varchar,
		PRIMARY KEY (link)
	)`
	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal("error creating rss_feeds table. err: ", err)
	}
}

func (db DB) InsertSrcRSSFeed(feed model.SrcRSSFeed) error {
	stmt, err := db.Prepare(`INSERT INTO src_rss_feeds
	( title,
		description,
		link,
		feed_link,
		updated,
		published,
		language,
		generator)
		values($1,$2,$3,$4,$5,$6,$7,$8)
		`)
	if err != nil {
		log.Error("failed to prepare rss_feeds insert: ", err)
		return err
	}
	//Todo: Write this correctly in regards to model
	// _, err = stmt.Exec(
	// 	feed.Title,
	// 	feed.Description,
	// 	feed.Link,
	// 	feed.FeedLink,
	// 	feed.UpdatedParsed,
	// 	feed.PublishedParsed,
	// 	feed.Language,
	// 	feed.Generator)
	// if err != nil {
	// 	log.Errorf("failed to exec rss_feeds insert: ", err)
	// 	return err
	// }
	return nil
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
	( userID varchar NOT NULL PRIMARY KEY,
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
