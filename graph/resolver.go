package graph

import (
	"github.com/Raven57/belajargraphql/graph/postgre"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	CommentsRepo postgre.CommentsRepo
	UsersRepo   postgre.UsersRepo
	PremiumtypesRepo postgre.PremiumtypesRepo

}
