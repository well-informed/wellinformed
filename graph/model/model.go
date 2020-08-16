package model

import "time"

type Pageable interface {
	GetID() string
}

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

func (src *SrcRSSFeed) GetID() string {
	return string(src.ID)
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

type UserSubscription struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"userID" db:"user_id"`
	SrcRSSFeedID int64     `json:"srcRSSFeed" db:"source_id"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
}

func (u *UserSubscription) GetID() string {
	return string(u.ID)
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
	SourceType  string     `json:"sourceType" db:"source_type"`
}

func (c *ContentItem) GetID() string {
	return string(c.ID)
}

type Interaction struct {
	ID            int64     `json:"id"`
	UserID        int64     `json:"user" db:"user_id"`
	ContentItemID int64     `json:"contentItem" db:"content_item_id"`
	ReadState     ReadState `json:"readState" db:"read_state"`
	PercentRead   *float64  `json:"percentRead" db:"percent_read"`
	CreatedAt     time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt     time.Time `json:"updatedAt" db:"updated_at"`
}

func (i *Interaction) GetID() string {
	return string(i.ID)
}
