package graph

import "github.com/well-informed/wellinformed/graph/model"

//Takes a user object and switches their active user feed to the provided feed ID
func (r *mutationResolver) switchActiveUserFeed(user *model.User, feedID int64) error {
	user.ActiveUserFeedID = feedID
	_, err := r.DB.UpdateUser(*user)
	return err
}
