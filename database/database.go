package database

import (
	"database/sql"
	"errors"
	"strings"

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
	createContentItemsTable(db)
	createUsersTable(db)
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

//TODO: complete implementation. Might be a good first candidate for an ORM.
func (db DB) InsertUserSubscription(user model.User, src model.SrcRSSFeed) (*model.UserSubscription, error) {
	log.Panic("not implemented")
	return nil, nil
	// stmt, err := db.Prepare(`INSERT INTO user_subscriptions`)
}

func (db DB) getUserByField(selection string, whereClause string, args ...interface{}) (model.User, error) {
	var user model.User

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

func (db DB) GetUserByEmail(value string) (model.User, error) {
	return db.getUserByField("*", "email = $1", value)
}

func (db DB) GetUserByUsername(value string) (model.User, error) {
	return db.getUserByField("*", "user_name = $1", value)
}

func (db DB) GetUserById(value string) (model.User, error) {
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
	 source_title varchar NOT NULL,
	 source_link varchar NOT NULL,
	 title varchar NOT NULL,
	 description varchar,
	 content varchar,
	 link varchar NOT NULL,
	 updated timestamp with time zone,
	 published timestamp with time zone,
	 author varchar,
	 guid varchar NOT NULL,
	 image_title varchar,
	 image_url varchar,
	 UNIQUE (source_id, link)
 )`
	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal("error creating content_items table. err: ", err)
	}
}

func (db DB) InsertContentItem(contentItem model.ContentItem) (*model.ContentItem, error) {
	log.Debugf("about to insert item with source_id: %v, link: %v", contentItem.SourceID, contentItem.Link)
	stmt, err := db.Prepare(`INSERT INTO content_items
	( source_id,
		source_title,
		source_link,
		title,
		description,
		content,
		link,
		updated,
		published,
		author,
		guid,
		image_title,
		image_url)
	values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
	RETURNING id`)
	if err != nil {
		log.Error("failed to prepare content_items insert", err)
		return nil, err
	}

	var id int64
	err = stmt.QueryRow(
		contentItem.SourceID,
		contentItem.SourceTitle,
		contentItem.SourceLink,
		contentItem.Title,
		contentItem.Description,
		contentItem.Content,
		contentItem.Link,
		contentItem.Updated,
		contentItem.Published,
		contentItem.Author,
		contentItem.GUID,
		contentItem.ImageTitle,
		contentItem.ImageURL,
	).Scan(&id)
	if err != nil {
		log.Errorf("failed to insert row to content_items. err: ", err)
		return nil, err
	}
	contentItem.ID = id
	return &contentItem, nil
}

func (db DB) ListContentItemsBySource(src *model.SrcRSSFeed) ([]*model.ContentItem, error) {
	log.Debug("received query with src feed object: ", src)
	stmt := `SELECT * FROM content_items WHERE source_id = $1`
	rows, err := db.Query(stmt, src.ID)
	if err != nil {
		log.Error("Error selecting content items by source from db. err: ", err)
		return nil, err
	}
	defer rows.Close()
	contentItems := make([]*model.ContentItem, 0)
	for rows.Next() {
		var contentItem model.ContentItem
		err := rows.Scan(
			&contentItem.ID,
			&contentItem.SourceID,
			&contentItem.SourceTitle,
			&contentItem.SourceLink,
			&contentItem.Title,
			&contentItem.Description,
			&contentItem.Content,
			&contentItem.Link,
			&contentItem.Updated,
			&contentItem.Published,
			&contentItem.Author,
			&contentItem.GUID,
			&contentItem.ImageTitle,
			&contentItem.ImageURL,
		)
		if err != nil {
			log.Error("error with scan. err: ", err)
			return nil, err
		}
		log.Debugf("selected contentItem, ID: %v, title: %v", contentItem.ID, contentItem.Title)
		contentItems = append(contentItems, &contentItem)
	}
	if err := rows.Err(); err != nil {
		log.Error("error while retrieving content items by source. err: ", err)
		return nil, err
	}
	return contentItems, nil
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
