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

type User struct {
	ID                  int64     `json:"id"`
	Firstname           string    `json:"firstname"`
	Lastname            string    `json:"lastname"`
	Username            string    `json:"username"`
	Email               string    `json:"email"`
	Password            string    `json:"password"`
	ActivePreferenceSet int64     `json:"activePreferenceSet"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
}
