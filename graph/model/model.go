package model

import "time"

type SrcRSSFeed struct {
	ID            int64     `json:"id"`
	Title         string    `json:"title"`
	Description   *string   `json:"description"`
	Link          string    `json:"link"`
	FeedLink      string    `json:"feedLink"`
	Updated       time.Time `json:"updated"`
	LastFetchedAt time.Time `json:"lastFetchedAt"`
	Language      *string   `json:"language"`
	Generator     *string   `json:"generator"`
}
