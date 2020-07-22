package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/Raven57/belajargraphql/graph/generated"
	"github.com/Raven57/belajargraphql/graph/model"
)

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return users, nil
}

func (r *queryResolver) Premiumtypes(ctx context.Context) ([]*model.Premiumtype, error) {
	return premiumtypes, nil
}

func (r *userResolver) Premiumtype(ctx context.Context, obj *model.User) (*model.Premiumtype, error) {
	prem := new(model.Premiumtype)

	for _, p := range premiumtypes {
		if p.ID == obj.PremiumID {
			prem = p
			break
		}
	}

	if prem == nil {
		return nil, errors.New("premium type tak ade")
	}

	return prem, nil
}

func (r *userResolver) Comments(ctx context.Context, obj *model.User) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var premiumtypes = []*model.Premiumtype{
	{
		ID:   "1",
		Text: "Premium",
	},
	{
		ID:   "2",
		Text: "Normal",
	},
}
var users = []*model.User{
	{
		ID:        "1",
		Name:      "Abc",
		PremiumID: "1",
	},
	{
		ID:        "2",
		Name:      "Nomor dua",
		PremiumID: "2",
	},
	{
		ID:        "3",
		Name:      "Nomor tiga",
		PremiumID: "2",
	},
}
