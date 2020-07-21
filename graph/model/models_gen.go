// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
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

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
