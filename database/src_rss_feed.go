package database

import (
	"database/sql"
	b64 "encoding/base64"
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

func buildPage(first int, after *string, edges []*model.SrcRSSFeedEdge) (*model.SrcRSSFeedsConnection, error) {
	log.Debugf("how many edges before prune? %d", len(edges))
	if after != nil {
		for i := 0; i < len(edges); i++ {
			if *after == edges[i].Cursor {
				if i+1 == len(edges) {
					return nil, errors.New("cursor not found in list")
				} else if i+first+1 > len(edges) {
					edges = edges[i+1:]
				} else {
					edges = edges[i+1 : i+first+1]
				}
				break
			}
		}
	} else if first < len(edges) {
		edges = edges[:first]
	}
	log.Debugf("how many edges after prune? %d", len(edges))
	info := &model.SrcRSSFeedsPageInfo{
		HasPreviousPage: len(edges) > 0 && after != nil,
		HasNextPage:     len(edges) > first,
		StartCursor:     edges[0].Cursor,
		EndCursor:       edges[len(edges)-1].Cursor,
	}
	return &model.SrcRSSFeedsConnection{
		Edges:    edges,
		PageInfo: info,
	}, nil
}

func feedsToEdges(feeds []*model.SrcRSSFeed) []*model.SrcRSSFeedEdge {
	edges := make([]*model.SrcRSSFeedEdge, 0)
	for _, feed := range feeds {
		edges = append(edges, &model.SrcRSSFeedEdge{
			Node:   feed,
			Cursor: b64.StdEncoding.EncodeToString([]byte(feed.Updated.String())),
		})
	}
	return edges
}

//Need to write this manually now...in order to select only the fields from the first table
//Maybe there's a way to only select some fields?
const feedsByUserStmt = `SELECT src_rss_feeds.*
FROM src_rss_feeds
INNER JOIN user_subscriptions
ON src_rss_feeds.id = user_subscriptions.source_id
WHERE user_subscriptions.user_id = $1 
ORDER BY src_rss_feeds.id`

func (db DB) PageSrcRSSFeedsByUser(user *model.User, input *model.SrcRSSFeedsConnectionInput) (*model.SrcRSSFeedsConnection, error) {
	feeds, err := db.listSrcRSSFeedsByQuery(feedsByUserStmt, user.ID)
	edges := feedsToEdges(feeds)
	if err != nil {
		log.Error("error selecting base list of src_rss_feeds. err: ", err)
	}
	return buildPage(input.First, input.After, edges)
}

func (db DB) ListSrcRSSFeedsByUser(user *model.User) ([]*model.SrcRSSFeed, error) {
	return db.listSrcRSSFeedsByQuery(feedsByUserStmt, user.ID)
}

const allFeedsStmt = `SELECT * FROM src_rss_feeds ORDER BY id`

func (db DB) PageSrcRSSFeeds(input *model.SrcRSSFeedsConnectionInput) (*model.SrcRSSFeedsConnection, error) {
	feeds, err := db.listSrcRSSFeedsByQuery(allFeedsStmt)
	edges := feedsToEdges(feeds)
	if err != nil {
		log.Error("error selecting base list of src_rss_feeds. err: ", err)
	}
	return buildPage(input.First, input.After, edges)
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
