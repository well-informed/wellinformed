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
	ID                      int64     `json:"id"`
	Firstname               string    `json:"firstname"`
	Lastname                string    `json:"lastname"`
	Username                string    `json:"username"`
	Email                   string    `json:"email"`
	Password                string    `json:"password"`
	ActivePreferenceSetName string    `json:"activePreferenceSet"`
	CreatedAt               time.Time `json:"createdAt"`
	UpdatedAt               time.Time `json:"updatedAt"`
}

type PreferenceSet struct {
	ID        int64      `json:"id"`
	UserID    int64      `json:"user" db:"user_id"`
	Name      string     `json:"name"`
	Sort      SortType   `json:"sort"`
	StartDate *time.Time `json:"startDate" db:"start_date"`
	EndDate   *time.Time `json:"endDate" db:"end_date"`
}

type ContentItem struct {
	ID          int64      `json:"id"`
	SourceID    int64      `json:"sourceID" db:"source_id"`
	SourceTitle string     `json:"sourceTitle" db:"source_title"`
	SourceLink  string     `json:"sourceLink" db:"source_link"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Content     string     `json:"content"`
	Link        string     `json:"link"`
	Updated     *time.Time `json:"updated"`
	Published   *time.Time `json:"published"`
	Author      *string    `json:"author"`
	GUID        *string    `json:"guid"`
	ImageTitle  *string    `json:"imageTitle" db:"image_title"`
	ImageURL    *string    `json:"imageURL" db:"image_url"`
	SourceType  string     `json:"sourceType" db:"source_type`
}
