package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"
	"my-go-gql-sample/database"
	"my-go-gql-sample/graph/generated"
	"my-go-gql-sample/graph/model"
	"my-go-gql-sample/util"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (string, error) {
	log.Printf("[mutationResolver.CreateTodo] input: %#v", input)
	id := util.CreateUniqueID()
	err := database.NewTodoDao(r.DB).InsertOne(&database.Todo{
		ID:     id,
		Text:   input.Text,
		Done:   false,
		UserID: input.UserID,
	})
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	log.Printf("[mutationResolver.CreateUser] input: %#v", input)
	id := util.CreateUniqueID()
	err := database.NewUserDao(r.DB).InsertOne(&database.User{
		ID:   id,
		Name: input.Name,
	})
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	log.Println("[queryResolver.Todos]")
	todos, err := database.NewTodoDao(r.DB).FindAll()
	if err != nil {
		return nil, err
	}
	var results []*model.Todo
	for _, todo := range todos {
		results = append(results, &model.Todo{
			ID:   todo.ID,
			Text: todo.Text,
			Done: todo.Done,
		})
	}
	return results, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	log.Printf("[queryResolver.Todo] id: %s", id)
	todo, err := database.NewTodoDao(r.DB).FindOne(id)
	if err != nil {
		return nil, err
	}
	if todo == nil {
		return nil, errors.New("not found")
	}
	return &model.Todo{
		ID:   todo.ID,
		Text: todo.Text,
		Done: todo.Done,
	}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	log.Println("[queryResolver.Users]")
	users, err := database.NewUserDao(r.DB).FindAll()
	if err != nil {
		return nil, err
	}
	var results []*model.User
	for _, user := range users {
		results = append(results, &model.User{
			ID:   user.ID,
			Name: user.Name,
		})
	}
	return results, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	log.Printf("[queryResolver.User] id: %s", id)
	user, err := database.NewUserDao(r.DB).FindOne(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("not found")
	}
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func (r *queryResolver) Schedules(ctx context.Context) ([]*model.Schedule, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Todos(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	log.Println("[userResolver.Todos]")
	todos, err := database.NewTodoDao(r.DB).FindByUserID(obj.ID)
	if err != nil {
		return nil, err
	}
	var results []*model.Todo
	for _, todo := range todos {
		results = append(results, &model.Todo{
			ID:   todo.ID,
			Text: todo.Text,
			Done: todo.Done,
		})
	}
	return results, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
