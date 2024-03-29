// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Feed interface {
	IsFeed()
}

type AddSourceInput struct {
	// the ID of the feed you want to subscribe to
	SourceFeedID int64 `json:"sourceFeedID"`
	// The source's __typename value
	SourceType SourceType `json:"sourceType"`
	// the ID of the feed that is subscribing to the source. defaults to the active feed
	TargetFeedID *int64 `json:"targetFeedID"`
}

type AddUserFeedInput struct {
	Name         string `json:"name"`
	EngineID     *int64 `json:"engineID"`
	ClonedFeedID *int64 `json:"clonedFeedID"`
}

type AuthResponse struct {
	AuthToken *AuthToken `json:"authToken"`
	User      *User      `json:"user"`
}

type AuthToken struct {
	AccessToken string    `json:"accessToken"`
	ExpiredAt   time.Time `json:"expiredAt"`
}

type ContentItemConnection struct {
	Edges    []*ContentItemEdge   `json:"edges"`
	PageInfo *ContentItemPageInfo `json:"pageInfo"`
}

type ContentItemConnectionInput struct {
	First int     `json:"first"`
	After *string `json:"after"`
}

type ContentItemEdge struct {
	Node   *ContentItem `json:"node"`
	Cursor string       `json:"cursor"`
}

type ContentItemInteractionsInput struct {
	UserID *int64 `json:"userID"`
}

type ContentItemPageInfo struct {
	HasPreviousPage bool   `json:"hasPreviousPage"`
	HasNextPage     bool   `json:"hasNextPage"`
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
}

type DeleteResponse struct {
	Ok bool `json:"ok"`
}

type EngineInput struct {
	Name      string     `json:"name"`
	Sort      SortType   `json:"sort"`
	StartDate *time.Time `json:"startDate"`
	EndDate   *time.Time `json:"endDate"`
}

type GetUserInput struct {
	UserID   *int64  `json:"userID"`
	Email    *string `json:"email"`
	Username *string `json:"username"`
}

type InteractionConnection struct {
	Edges    []*InteractionEdge   `json:"edges"`
	PageInfo *InteractionPageInfo `json:"pageInfo"`
}

type InteractionConnectionInput struct {
	First int     `json:"first"`
	After *string `json:"after"`
}

type InteractionEdge struct {
	Node   *Interaction `json:"node"`
	Cursor string       `json:"cursor"`
}

type InteractionInput struct {
	ContentItemID int64     `json:"contentItemID"`
	ReadState     ReadState `json:"readState"`
	Completed     *bool     `json:"completed"`
	SavedForLater *bool     `json:"savedForLater"`
	PercentRead   *float64  `json:"percentRead"`
	Rating        *int      `json:"rating"`
}

type InteractionPageInfo struct {
	HasPreviousPage bool   `json:"hasPreviousPage"`
	HasNextPage     bool   `json:"hasNextPage"`
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
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

type SrcRSSFeedConnection struct {
	Edges    []*SrcRSSFeedEdge   `json:"edges"`
	PageInfo *SrcRSSFeedPageInfo `json:"pageInfo"`
}

type SrcRSSFeedConnectionInput struct {
	First int     `json:"first"`
	After *string `json:"after"`
}

type SrcRSSFeedEdge struct {
	Node   *SrcRSSFeed `json:"node"`
	Cursor string      `json:"cursor"`
}

type SrcRSSFeedInput struct {
	ID       *int64  `json:"id"`
	Link     *string `json:"link"`
	FeedLink *string `json:"feedLink"`
}

type SrcRSSFeedPageInfo struct {
	HasPreviousPage bool   `json:"hasPreviousPage"`
	HasNextPage     bool   `json:"hasNextPage"`
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
}

type UserInteractionsInput struct {
	ReadState *ReadState `json:"readState"`
}

type UserSubscriptionConnection struct {
	Edges    []*UserSubscriptionEdge   `json:"edges"`
	PageInfo *UserSubscriptionPageInfo `json:"pageInfo"`
}

type UserSubscriptionConnectionInput struct {
	First int     `json:"first"`
	After *string `json:"after"`
}

type UserSubscriptionEdge struct {
	Node   *UserSubscription `json:"node"`
	Cursor string            `json:"cursor"`
}

type UserSubscriptionPageInfo struct {
	HasPreviousPage bool   `json:"hasPreviousPage"`
	HasNextPage     bool   `json:"hasNextPage"`
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
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

type SourceType string

const (
	SourceTypeUnknown    SourceType = "Unknown"
	SourceTypeSrcRSSFeed SourceType = "SrcRSSFeed"
	SourceTypeUserFeed   SourceType = "UserFeed"
)

var AllSourceType = []SourceType{
	SourceTypeUnknown,
	SourceTypeSrcRSSFeed,
	SourceTypeUserFeed,
}

func (e SourceType) IsValid() bool {
	switch e {
	case SourceTypeUnknown, SourceTypeSrcRSSFeed, SourceTypeUserFeed:
		return true
	}
	return false
}

func (e SourceType) String() string {
	return string(e)
}

func (e *SourceType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SourceType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SourceType", str)
	}
	return nil
}

func (e SourceType) MarshalGQL(w io.Writer) {
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
