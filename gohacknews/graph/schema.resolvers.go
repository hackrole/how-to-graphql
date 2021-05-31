package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/hackrole/learn-graphql/gohacknews/graph/generated"
	"github.com/hackrole/learn-graphql/gohacknews/graph/model"
)

// Links
func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var links []*model.Link
	dummyLink := model.Link{
		Title:   "our dummy link",
		Address: "https://address.org",
		User:    &model.User{Name: "admin"},
	}
	links = append(links, &dummyLink)
	return links, nil
}

// CreateLink
func (r *mutationResolver) CreateLink(ctx context.Context, input *model.NewLink) (*model.Link, error) {
	var link model.Link
	var user model.User
	link.Address = input.Address
	link.Title = input.Title
	user.Name = "test"
	link.User = &user
	return &link, nil
}

// CreateUser
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	return "hello", nil
}

// Login
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	return "", nil
}

// refreshToken
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	return "", nil
}

// Query
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

// Mutation
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
