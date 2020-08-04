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

type ContentItemInteractionsInput struct {
	UserID *int64 `json:"userID"`
}

type DeleteResponse struct {
	Ok bool `json:"ok"`
}

type GetUserInput struct {
	UserID   *int64  `json:"userID"`
	Email    *string `json:"email"`
	Username *string `json:"username"`
}

type InteractionInput struct {
	ContentItemID int64     `json:"contentItemID"`
	ReadState     ReadState `json:"readState"`
	PercentRead   *int      `json:"percentRead"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PreferenceSetInput struct {
	Name string `json:"name"`
	// true sets the entered preference set as active, false never has any effect.
	// A prefSet can only become inactive if another prefSet is set to active
	Activate  bool       `json:"activate"`
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

type UserInteractionsInput struct {
	ReadState *ReadState `json:"readState"`
}

type ReadState string

const (
	ReadStateCompleted     ReadState = "completed"
	ReadStateSavedForLater ReadState = "savedForLater"
	ReadStatePartiallyRead ReadState = "partiallyRead"
	ReadStateUnread        ReadState = "unread"
)

var AllReadState = []ReadState{
	ReadStateCompleted,
	ReadStateSavedForLater,
	ReadStatePartiallyRead,
	ReadStateUnread,
}

func (e ReadState) IsValid() bool {
	switch e {
	case ReadStateCompleted, ReadStateSavedForLater, ReadStatePartiallyRead, ReadStateUnread:
		return true
	}
	return false
}

func (e ReadState) String() string {
	return string(e)
}

func (e *ReadState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ReadState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ReadState", str)
	}
	return nil
}

func (e ReadState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
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
