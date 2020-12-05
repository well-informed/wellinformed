package wellinformed

import (
	"context"
	"time"

	"github.com/well-informed/wellinformed/graph/model"
)

type Persistor interface {

	// SrcRSSFeed
	InsertSrcRSSFeed(model.SrcRSSFeed) (*model.SrcRSSFeed, error)
	// GetSrcRSSFeed(model.SrcRSSFeedInput) (*model.SrcRSSFeed, error)
	GetSrcRSSFeedByID(id int64) (*model.SrcRSSFeed, error)
	GetSrcRSSFeedByLink(link string) (*model.SrcRSSFeed, error)
	GetSrcRSSFeedByFeedLink(feedLink string) (*model.SrcRSSFeed, error)
	ListSrcRSSFeeds() ([]*model.SrcRSSFeed, error)
	ListSrcRSSFeedsByUser(*model.User) ([]*model.SrcRSSFeed, error)

	// UserSubscription
	InsertUserSubscription(user model.User, src model.SrcRSSFeed) (*model.UserSubscription, error)
	GetUserSubscription(userID int64, srcID int64) (*model.UserSubscription, error)
	DeleteUserSubscription(userID int64, srcID int64) (int, error)
	ListUserSubscriptions(userID int64) ([]*model.UserSubscription, error)

	// ContentItem
	GetContentItem(id int64) (*model.ContentItem, error)
	InsertContentItem(model.ContentItem) (*model.ContentItem, error)
	ListContentItemsBySource(*model.SrcRSSFeed) ([]*model.ContentItem, error)
	ServeContentItems([]*model.SrcRSSFeed, *time.Time, *time.Time) ([]*model.ContentItem, error)

	// User
	GetUserByEmail(email string) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	GetUserByID(id int64) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)

	// UserFeed
	CreateUserFeed(userFeed *model.UserFeed) (*model.UserFeed, error)
	GetUserFeedByID(id int64) (*model.UserFeed, error)
	ListUserFeedsByUser(userID int64) ([]*model.UserFeed, error)

	// Engine
	SaveEngine(engine *model.Engine) (*model.Engine, error)
	ListEnginesByUser(userID int64) ([]*model.Engine, error)
	GetEngineByID(id int64) (*model.Engine, error)
	GetEngineByName(userID int64, name string) (*model.Engine, error)

	// Interaction
	SaveInteraction(userID int64, interaction *model.InteractionInput) (*model.ContentItem, error)
	ListUserInteractions(userID int64, readState *model.ReadState) ([]*model.Interaction, error)
	GetInteractionByContentID(userID int64, contentItemID int64) (*model.Interaction, error)
	GetUserByInteraction(interactionID int64) (*model.User, error)
	GetContentItemByInteraction(contentItemID int64) (*model.ContentItem, error)

	// FeedSubscription
	CreateFeedSubscription(feedID int64, sourceID int64, sourceType model.SourceType) (*model.FeedSubscription, error)
	ListFeedSubscriptionsByFeedID(feedID int64) ([]*model.FeedSubscription, error)

	//UserRelationship
	SaveUserRelationship(followerID int64, followeeID int64) (*model.UserRelationship, error)
	DeleteUserRelationship(followerID int64, followeeID int64) error
	ListUserRelationshipsByFollowerID(followerID int64) ([]*model.UserRelationship, error)
	ListUserRelationshipsByFolloweeID(followeeID int64) ([]*model.UserRelationship, error)
}

type RSS interface {
	FetchSrcFeed(feedLink string, ctx context.Context) (model.SrcRSSFeed, []*model.ContentItem, error)
}

type Subscriber interface {
	SubscribeToRSSFeed(ctx context.Context, feedLink string) (*model.SrcRSSFeed, error)
	AddUserSubscription(user *model.User, srcRSSFeed *model.SrcRSSFeed) (*model.UserSubscription, error)
}

type FeedService interface {
	ServeContent(ctx context.Context, feed *model.UserFeed) ([]*model.ContentItem, error)
}
