// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type AuthResponse struct {
	AuthToken *AuthToken `json:"authToken"`
	User      *User      `json:"user"`
}

type AuthToken struct {
	AccessToken string    `json:"accessToken"`
	ExpiredAt   time.Time `json:"expiredAt"`
}

type ContentItem struct {
	ID          int64      `json:"id"`
	SourceID    int64      `json:"sourceID"`
	SourceTitle string     `json:"sourceTitle"`
	SourceLink  string     `json:"sourceLink"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Content     string     `json:"content"`
	Link        string     `json:"link"`
	Updated     *time.Time `json:"updated"`
	Published   *time.Time `json:"published"`
	Author      *string    `json:"author"`
	GUID        *string    `json:"guid"`
	ImageTitle  *string    `json:"imageTitle"`
	ImageURL    *string    `json:"imageURL"`
	SourceType  string     `json:"sourceType"`
}

type DeleteResponse struct {
	Ok bool `json:"ok"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PreferenceSet struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Sort      SortType   `json:"sort"`
	StartDate *time.Time `json:"startDate"`
	EndDate   *time.Time `json:"endDate"`
}

type RegisterInput struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
}

type SrcRSSFeedInput struct {
	ID       *int64  `json:"id"`
	Link     *string `json:"link"`
	FeedLink *string `json:"feedLink"`
}

type UserFeed struct {
	UserID       int64          `json:"userID"`
	Name         string         `json:"name"`
	ContentItems []*ContentItem `json:"contentItems"`
}

type UserSubscription struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"userID"`
	SrcRSSFeed int64     `json:"srcRSSFeed"`
	CreatedAt  time.Time `json:"createdAt"`
}

type SortType string

const (
	SortTypeChronological SortType = "chronological"
	SortTypeSourceName    SortType = "sourceName"
)

var AllSortType = []SortType{
	SortTypeChronological,
	SortTypeSourceName,
}

func (e SortType) IsValid() bool {
	switch e {
	case SortTypeChronological, SortTypeSourceName:
		return true
	}
	return false
}

func (e SortType) String() string {
	return string(e)
}

func (e *SortType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid sortType", str)
	}
	return nil
}

func (e SortType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
