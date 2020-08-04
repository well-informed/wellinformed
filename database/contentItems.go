package database

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/graph/model"
)

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
	ON CONFLICT DO NOTHING
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
	if err != nil && err != sql.ErrNoRows {
		log.Error("failed to insert row to content_items. err: ", err)
		return nil, err
	}
	//May return a zero ID if duplicate entry already exists
	contentItem.ID = id
	return &contentItem, nil
}

func (db DB) GetContentItem(id int64) (*model.ContentItem, error) {
	stmt := `SELECT * FROM content_items WHERE id = $1`

	var contentItem model.ContentItem
	err := db.QueryRow(stmt, id).Scan(
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
		log.Errorf("failed to select content_item. err: err")
		return nil, err
	}
	return &contentItem, nil
}

func (db DB) ListContentItemsBySource(src *model.SrcRSSFeed) ([]*model.ContentItem, error) {
	log.Debug("received query with src feed object: ", src)
	stmt := `SELECT * FROM content_items WHERE source_id = $1`
	rows, err := db.Query(stmt, src.ID)
	defer rows.Close()
	if err != nil {
		log.Error("Error selecting content items by source from db. err: ", err)
		return nil, err
	}
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

func (db DB) ServeContentItems(srcList []*model.SrcRSSFeed, sortType model.SortType, start_dt *time.Time, end_dt *time.Time) ([]*model.ContentItem, error) {
	contentItems := make([]*model.ContentItem, 0)
	var srcIDs []int64
	for _, src := range srcList {
		srcIDs = append(srcIDs, src.ID)
	}
	var order string
	if sortType == model.SortTypeChronological {
		order = "ORDER BY published DESC"
	} else if sortType == model.SortTypeSourceName {
		order = "ORDER BY source_title DESC"
	}
	//Set up selection of source values
	query, args, err := sqlx.In("SELECT * FROM content_items WHERE source_id IN (?) ", srcIDs)
	if err != nil {
		log.Error("couldn't construct sqlx.In query. err: ", err)
		return nil, err
	}

	//Add start and end datetime if available
	var startWhere string
	var endWhere string
	if start_dt != nil {
		startWhere = `AND published >= ? `
		query = query + startWhere
		args = append(args, start_dt)
	}
	if end_dt != nil {
		endWhere = `AND published <= ? `
		query = query + endWhere
		args = append(args, end_dt)
	}

	//Add sorting
	query = query + order

	//rebind with postgres style parameters
	query = db.Rebind(query)

	err = db.Select(&contentItems, query, args...)
	if err != nil {
		log.Error("could not execute ServeContentItems query: ", query)
		log.Error("err: ", err)
		return nil, err
	}
	return contentItems, nil

}

func (db DB) GetContentItemByInteraction(interactionId int64) (*model.ContentItem, error) {
	stmt := `
	SELECT c.* FROM content_items c 
	INNER JOIN interactions i on c.id = i.content_item_id 
	WHERE i.id = $1
	LIMIT 1
	`

	var contentItem model.ContentItem
	err := db.QueryRow(stmt, interactionId).Scan(
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
		log.Errorf("failed to select content_item. err: err")
		return nil, err
	}
	return &contentItem, nil
}
