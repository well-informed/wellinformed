package database

import (
	"database/sql"
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/graph/model"
)

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
		log.Error("failed to insert row to src_rss_feeds. err: ", err)
		return nil, err
	}
	feed.ID = id
	log.Info("got id back: ", id)
	return &feed, nil
}

func (db DB) GetSrcRSSFeed(input model.SrcRSSFeedInput) (*model.SrcRSSFeed, error) {
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

func feedsToNodes(feeds []*model.SrcRSSFeed) []*model.Node {
	nodes := make([]*model.Node, 0)
	for _, feed := range feeds {
		nodes = append(nodes, &model.Node{
			Value: feed,
			ID:    feed.ID,
		})
	}
	return nodes
}

//Need to write this manually now...in order to select only the fields from the first table
//Maybe there's a way to only select some fields?
const feedsByUserStmt = `SELECT src_rss_feeds.*
FROM src_rss_feeds
INNER JOIN user_subscriptions
ON src_rss_feeds.id = user_subscriptions.source_id
WHERE user_subscriptions.user_id = $1
ORDER BY src_rss_feeds.id`

func (db DB) PageSrcRSSFeedsByUser(user *model.User, input *model.ConnectionInput) (*model.Connection, error) {
	feeds, err := db.listSrcRSSFeedsByQuery(feedsByUserStmt, user.ID)
	edges := nodesToEdges(feedsToNodes(feeds))
	if err != nil {
		log.Error("error selecting base list of src_rss_feeds. err: ", err)
	}
	return buildPage(input.First, input.After, feeds)
}

func (db DB) ListSrcRSSFeedsByUser(user *model.User) ([]*model.SrcRSSFeed, error) {
	return db.listSrcRSSFeedsByQuery(feedsByUserStmt, user.ID)
}

const allFeedsStmt = `SELECT * FROM src_rss_feeds ORDER BY id`

func (db DB) PageSrcRSSFeeds(input *model.ConnectionInput) (*model.Connection, error) {
	feeds, err := db.listSrcRSSFeedsByQuery(allFeedsStmt)
	if err != nil {
		log.Error("error selecting base list of src_rss_feeds. err: ", err)
	}
	var node model.Pageable
	node = feeds[0]
	// nodes := make([]*model.Pageable, len(feeds))
	return buildPage(input.First, input.After, []model.Pageable{node})
}

func (db DB) ListSrcRSSFeeds() ([]*model.SrcRSSFeed, error) {
	return db.listSrcRSSFeedsByQuery(allFeedsStmt)
}

func (db DB) listSrcRSSFeedsByQuery(stmt string, args ...interface{}) ([]*model.SrcRSSFeed, error) {
	rows, err := db.Query(stmt, args...)
	defer rows.Close()
	if err != nil {
		log.Error("error selecting all srcRSSFeeds. err: ", err)
		return nil, err
	}
	feeds := make([]*model.SrcRSSFeed, 0)
	for rows.Next() {
		var feed model.SrcRSSFeed
		err := rows.Scan(
			&feed.ID,
			&feed.Title,
			&feed.Description,
			&feed.Link,
			&feed.FeedLink,
			&feed.Updated,
			&feed.LastFetchedAt,
			&feed.Language,
			&feed.Generator,
		)
		if err != nil {
			log.Error("error scanning srcFeed row: err: ", err)
		}
		feeds = append(feeds, &feed)
	}
	if err := rows.Err(); err != nil {
		log.Error("error listing srcRSSFeeds. err: ", err)
		return nil, err
	}
	return feeds, nil
}
