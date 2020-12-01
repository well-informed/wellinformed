package graph

import (
	"errors"

	"github.com/well-informed/wellinformed/graph/model"
)

func (r *mutationResolver) saveInteraction(user *model.User, input *model.InteractionInput) (*model.ContentItem, error) {
	//Ensure optional arguments are set to defaults before insert into DB
	var f = false
	if input.Completed == nil {
		input.Completed = &f
	}
	if input.SavedForLater == nil {
		input.SavedForLater = &f
	}

	//If rating included, check that it's between 0 and 10
	if input.Rating != nil {
		if *input.Rating < 0 || *input.Rating > 10 {
			return nil, errors.New("invalid rating: must be between 0 and 10.")
		}
	}
	return r.DB.SaveInteraction(user.ID, input)
}
