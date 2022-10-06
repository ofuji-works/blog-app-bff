package graph

import "github.com/ofuji-works/blog-app-bff/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	web_vitals_metric []*model.WebVitalsMetric
}
