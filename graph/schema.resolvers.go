package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/Raven57/belajargraphql/graph/generated"
	"github.com/Raven57/belajargraphql/graph/model"
)

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	//var restrictions []*model.User
	//
	//err := r.DB.Model(&restrictions).Select()
	//fmt.Println(err)
	//if err != nil {
	//
	//	return nil, errors.New("Failed to query")
	//}

	return r.UsersRepo.GetAllUsers()
}

func (r *queryResolver) Premiumtypes(ctx context.Context) ([]*model.Premiumtype, error) {

	return r.PremiumtypesRepo.GetAllPremiums()
}

func (r *queryResolver) Comments(ctx context.Context) ([]*model.Comment, error) {
	return r.CommentsRepo.GetAllComments()
}

func (r *userResolver) Premiumtype(ctx context.Context, obj *model.User) (*model.Premiumtype, error) {
	//prem := new(model.Premiumtype)
	//
	//for _, p := range premiumtypes {
	//	if p.ID == obj.PremiumID {
	//		prem = p
	//		break
	//	}
	//}
	//
	//if prem == nil {
	//	return nil, nil
	//}
	//
	//return prem, nil

	//var restrictions *model.Premiumtype
	//
	//err := r.DB.Model(&restrictions).Where("id=?", obj.PremiumID).First()
	//fmt.Println(err)
	//if err != nil {
	//
	//	return nil, errors.New("Failed to query")
	//}

	return r.PremiumtypesRepo.GetPremiumByID(obj.PremiumID)
}

func (r *userResolver) Comments(ctx context.Context, obj *model.User) ([]*model.Comment, error) {
	//var com []*model.Comment
	//
	//err := r.DB.Model(&com).Where("id=?", obj.ID).Select()
	//if err != nil {
	//	return nil, nil
	//}
	//
	//if com == nil {
	//	return nil, nil
	//}
	//return com, nil
	return r.CommentsRepo.CommentsByUserID(obj.ID)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
