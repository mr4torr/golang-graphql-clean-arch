package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mr4torr/go-lang-graphql/entity"
	"github.com/mr4torr/go-lang-graphql/entrypoint/graph/generated"
)

func (r *todoResolver) Active(ctx context.Context, obj *entity.Todo) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
