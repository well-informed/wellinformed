package model

import (
	"time"
)

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

func (SrcRSSFeed) IsFeed() {}

type User struct {
	ID               int64     `json:"id"`
	Firstname        string    `json:"firstname" db:"first_name"`
	Lastname         string    `json:"lastname" db:"last_name"`
	Username         string    `json:"username" db:"user_name"`
	Email            string    `json:"email" db:"email"`
	Password         string    `json:"password" db:"password"`
	ActiveUserFeedID int64     `json:"activeUserFeed" db:"active_user_feed"`
	CreatedAt        time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt        time.Time `json:"updatedAt" db:"updated_at"`
}

type UserFeed struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"userID" db:"user_id"`
	EngineID  int64  `json:"engine" db:"engine_id"`
	Title     string `json:"title"`
	Name      string `json:"name"`
	CreatedAt string `json:"CreatedAt" db:"created_at"`
	UpdatedAt string `json:"UpdatedAt" db:"updated_at"`
}

func (UserFeed) IsFeed() {}

type UserSubscription struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"userID" db:"user_id"`
	SrcRSSFeedID int64     `json:"srcRSSFeed" db:"source_id"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
}

type FeedSubscription struct {
	ID         int64      `json:"id"`
	UserFeedID int64      `json:"userFeed" db:"feed_id"`
	SourceID   int64      `json:"source" db:"source_id"`
	SourceType SourceType `json:"sourceType" db:"source_type"`
	CreatedAt  time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt  time.Time  `json:"updatedAt" db:"updated_at"`
}

type Engine struct {
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
	SourceType  SourceType `json:"sourceType" db:"source_type"`
}

type Interaction struct {
	ID            int64     `json:"id"`
	UserID        int64     `json:"user" db:"user_id"`
	ContentItemID int64     `json:"contentItem" db:"content_item_id"`
	ReadState     ReadState `json:"readState" db:"read_state"`
	Completed     bool      `json:"completed" db:"completed"`
	SavedForLater bool      `json:"savedForLater" db:"saved_for_later"`
	PercentRead   *float64  `json:"percentRead" db:"percent_read"`
	Rating        *int      `json:"rating" db:"rating"`
	CreatedAt     time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt     time.Time `json:"updatedAt" db:"updated_at"`
}

type UserRelationship struct {
	ID        int64     `json:"id"`
	Follower  int64     `json:"follower" db:"follower_id"`
	Followee  int64     `json:"followee" db:"followee_id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
