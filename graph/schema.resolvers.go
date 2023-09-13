package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"

	"github.com/subrotokumar/stellerlink-backend/database"
	"github.com/subrotokumar/stellerlink-backend/model"
)

// Hello is the resolver for the hello field.
func (r *mutationResolver) Hello(ctx context.Context) (string, error) {
	return "Hello World", nil
}

// AddCharacter is the resolver for the addCharacter field.
func (r *mutationResolver) AddCharacter(ctx context.Context, input *model.CharacterInput) (*model.Character, error) {
	return db.AddCharacter(input), nil
}

// Character is the resolver for the Character field.
func (r *queryResolver) Character(ctx context.Context, id int) (*model.Character, error) {
	return db.GetCharacter(id), nil
}

// Characters is the resolver for the Characters field.
func (r *queryResolver) Characters(ctx context.Context) ([]*model.Character, error) {
	return db.GetCharacters(), nil
}

// Relic is the resolver for the relic field.
func (r *queryResolver) Relic(ctx context.Context, id int) (*model.Relic, error) {
	return db.GetRelic(id), nil
}

// Relics is the resolver for the relics field.
func (r *queryResolver) Relics(ctx context.Context) ([]*model.Relic, error) {
	return db.GetRelics(), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
var db *database.DB = database.Connect()
