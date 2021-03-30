package ent

import (
	"context"
	"entgo.io/ent"
	"github.com/goxiaoy/go-saas/common"
	"github.com/goxiaoy/go-saas/data"
)

// Multi tenancy policy: auto set before mutated or auto apply filter to entity
type MultiTenancyPolicy struct {
}

func (m MultiTenancyPolicy) EvalMutation(ctx context.Context, mutation ent.Mutation) error {
	common.FromCurrentTenant(ctx)
	at := data.FromAutoSetTenantId(ctx)
	if at {
		//TODO build normalizer
	}
	return nil
}

func (m MultiTenancyPolicy) EvalQuery(ctx context.Context, query ent.Query) error {
	common.FromCurrentTenant(ctx)
	df := data.FromMultiTenancyDataFilter(ctx)
	if df {
		//TODO build filter
	}
	return nil
}
