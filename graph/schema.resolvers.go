package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"artvortex-api/graph/model"
	pg "artvortex-api/internal/postgres"
	"context"
	"fmt"
	"time"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	var timestamp = time.Now().Format(time.RFC3339)
	pg.DB.Omit("ID", "CreditBalance", "DisplayName", "WalletAddress").Create(&model.User{
		UserName:  input.UserName,
		Email:     input.Email,
		UpdatedAt: &timestamp,
		CreatedAt: &timestamp,
	})
	var user model.User
	res := pg.DB.Last(&user)

	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	//define users
	var users []*model.User
	//fetch all users
	pg.DB.Find(&users)

	return users, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
