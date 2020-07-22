package graph

import "github.com/well-informed/wellinformed"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB          wellinformed.Persistor
	RSS         wellinformed.RSS
	Sub         wellinformed.Subscriber
	Feed        wellinformed.FeedService
	UserService wellinformed.UserService
}
