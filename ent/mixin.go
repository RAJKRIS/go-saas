package ent

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

const TenantIdField = "tenant_id"

// mixin for multi-tenancy entities
type MultiTenantMixin struct {
	mixin.Schema
}

func (MultiTenantMixin) Indexes() []ent.Index {
	return []ent.Index{
		//add index
		index.Fields(TenantIdField),
	}
}
func (MultiTenantMixin) Fields() []ent.Field {
	return []ent.Field{
		//set to str
		field.String(TenantIdField),
	}
}

func (MultiTenantMixin) Policy() ent.Policy {
	return MultiTenancyPrivacy{}
}
