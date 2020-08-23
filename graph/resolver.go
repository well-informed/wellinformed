package graph

import (
	"github.com/well-informed/wellinformed"
	"github.com/well-informed/wellinformed/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB          wellinformed.Persistor
	RSS         wellinformed.RSS
	Sub         wellinformed.Subscriber
	FeedService wellinformed.FeedService
	UserService *user.UserService
}
