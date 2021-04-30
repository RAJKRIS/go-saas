package ent

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/privacy"
	"fmt"
	"github.com/goxiaoy/go-saas/common"
	"github.com/goxiaoy/go-saas/data"
	"github.com/pkg/errors"
)

// Multi tenancy policy: auto set before mutated or auto apply filter to entity
type MultiTenancyPrivacy struct {
}

func (m MultiTenancyPrivacy) EvalMutation(ctx context.Context, mutation ent.Mutation) error {
	ct := common.FromCurrentTenant(ctx)
	at := data.FromAutoSetTenantId(ctx)
	if at {
		err := mutation.SetField(TenantIdField, ct)
		if err != nil {
			return errors.Wrap(err, "Fail to set tenant id")
		}
	}
	return nil
}

func (m MultiTenancyPrivacy) EvalQuery(ctx context.Context, query ent.Query) error {
	ct := common.FromCurrentTenant(ctx)
	df := data.FromMultiTenancyDataFilter(ctx)
	if df {
		// TODO(Goxiaoy) https://github.com/ent/ent/issues/833
		panic("UnImplemented")
		type WhereQuery interface {
			Where(func(s *sql.Selector))
		}
		type Entity interface {
			Query() WhereQuery
		}

		tq, ok := query.(Entity)
		if !ok {
			return fmt.Errorf("unexpected query type %T"+": %w", query, privacy.Deny)
		}
		fq := func(s *sql.Selector) {
			if ct.Id == "" {
				//host side
				s.Where(sql.IsNull(TenantIdField))
			} else {
				s.Where(sql.EQ(s.C(TenantIdField), ct))
			}
		}
		tq.Query().Where(fq)
	}
	return nil
}
