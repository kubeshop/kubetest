package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.27

import (
	"context"
	"fmt"

	"github.com/kubeshop/testkube/internal/graphql/graph/model"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/rand"
)

// Executors is the resolver for the executors field.
func (r *queryResolver) Executors(ctx context.Context) ([]*model.Executor, error) {
	panic(fmt.Errorf("not implemented: Executors - executors"))
}

// Executor is the resolver for the executor field.
func (r *subscriptionResolver) Executor(ctx context.Context) (<-chan *model.ExecutorDetails, error) {
	ch := make(chan *model.ExecutorDetails)

	// TODO You can (and probably should) handle your channels in a central place outside of `schema.resolvers.go`.
	// For this example we'll simply use a Goroutine with a simple loop.
	go func() {
		r.Log.Infof("%+v\n", "subscribed to events.executor.>")

		r.Bus.SubscribeTopic("events.executor.>", rand.String(30), func(e testkube.Event) error {
			r.Log.Infof("%s %s %s\n", e.Type_, *e.Resource, e.ResourceId)

			exec, err := r.Client.Get(e.ResourceId)
			if err != nil {
				return err
			}

			// TODO valid data mapper between types
			// It's a little bit tricky because we have to map between different types based
			// on even more ugly pointers than spec (slices have pointers too :/ )
			ch <- &model.ExecutorDetails{Name: &exec.Name, Executor: &model.Executor{
				Image: &exec.Spec.Image,
			}}
			return nil
		})
	}()

	return ch, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
