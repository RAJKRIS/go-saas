package seed

import (
	"context"
	"github.com/goxiaoy/go-saas/common"
)

type Seeder interface {
	Seed(ctx context.Context)
}

type DefaultSeeder struct {
	extra map[string]interface{}
	opt   *SeedOption
}

func NewDefaultSeeder(opt *SeedOption, extra map[string]interface{}) *DefaultSeeder {
	return &DefaultSeeder{
		opt:   opt,
		extra: extra,
	}
}

func (d *DefaultSeeder) Seed(ctx context.Context) {
	for _, tenant := range d.opt.TenantIds {
		// change to next tenant
		newCtx := common.NewCurrentTenant(ctx, tenant, "")
		sCtx := NewSeedContext(tenant, d.extra)
		//create seeder
		for _, contributor := range d.opt.Contributors {
			contributor.Seed(newCtx, sCtx)
		}
	}
}