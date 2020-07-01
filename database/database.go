package database

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
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
	// createUsersTable(db)
	// createUserHistoryTable(db)
}

func createSrcRSSFeedsTable(db *sql.DB) {
	stmt := `
	CREATE TABLE IF NOT EXISTS src_rss_feeds
	(	id SERIAL PRIMARY KEY,
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

func (db DB) InsertSrcRSSFeed(feed model.SrcRSSFeed) (*model.SrcRSSFeed, error) {
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
		return nil, err
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
		return nil, err
	}
	feed.ID = id
	log.Info("got id back: ", id)
	return &feed, nil
}

func (db DB) SelectSrcRSSFeed(input model.SrcRSSFeedInput) (*model.SrcRSSFeed, error) {
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
		return nil, errors.New("no key for select found")
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
	if err == sql.ErrNoRows {
		return nil, nil
	}
	log.Debugf("Selected feed with whereClause %v with key %v: %v", whereClause, arg, feed)
	return &feed, err
}

func createContentItemsTable(db *sql.DB) {
	stmt := `
	CREATE TABLE IF NOT EXISTS content_items
 ( id BIGSERIAL PRIMARY KEY,
	 source_id int NOT NULL REFERENCES src_rss_feeds(id),
	 source_title NOT NULL,
	 source_link NOT NULL,
	 title varchar NOT NULL,
	 description varchar,
	 content varchar,
	 link varchar NOT NULL,
	 updated timestamp with time zone,
	 published timestamp with time zone,
	 author varchar,
	 guid varchar NOT NULL,
	 image_title varchar,
	 image_url varchar
 )`
	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal("error creating content_items table. err: ", err)
	}
}

func (db DB) InsertContentItem(model.ContentItem) (*model.ContentItem, error) {
	// stmt, err := db.Prepare(`INSERT INTO main_feed
	// (title,
	// description,
	// htmlContent,
	// link,
	// updated,
	// published,
	// author,
	// guid,
	// parent_feed)
	// values($1,$2,$3,$4,$5,$6,$7,$8,$9)`)
	// if err != nil {
	// 	log.Error("failed to prepare rss_articles insert", err)
	// 	return err
	// }

	// _, err = stmt.Exec(
	// 	article.Title,
	// 	article.Description,
	// 	article.Content,
	// 	article.Link,
	// 	article.UpdatedParsed,
	// 	article.PublishedParsed,
	// 	article.Author.Name,
	// 	article.GUID,
	// 	feedLink)
	// if err != nil {
	// 	log.Errorf("failed to exec insert of content %+v. err: %v", article, err)
	// 	return err
	// }
	// return nil
	log.Panic("not implemented")
	return &model.ContentItem{}, nil
}

func (db DB) ListContentItems() ([]model.ContentItem, error) {
	log.Panic("not implemented")
	return []model.ContentItem{}, nil
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
