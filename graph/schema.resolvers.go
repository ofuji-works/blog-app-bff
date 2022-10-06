package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ofuji-works/blog-app-bff/graph/generated"
	"github.com/ofuji-works/blog-app-bff/graph/model"
)

// CreateWebVitalsMetric is the resolver for the createWebVitalsMetric field.
func (r *mutationResolver) CreateWebVitalsMetric(ctx context.Context, data *model.WebVitalsMetricCreateInput) (*model.WebVitalsMetric, error) {
	web_vitals_metric := &model.WebVitalsMetric{
		Name:  data.Name,
		Value: data.Value,
	}
	return web_vitals_metric, nil
}

// WebVitalsMetric is the resolver for the web_vitals_metric field.
func (r *queryResolver) WebVitalsMetric(ctx context.Context) ([]*model.WebVitalsMetric, error) {
	panic(fmt.Errorf("not implemented: WebVitalsMetric - web_vitals_metric"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
