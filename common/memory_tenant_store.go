package common

import (
	"context"
)

type MemoryTenantStore struct {
	TenantConfig []TenantConfig
}

func NewMemoryTenantStore(t []TenantConfig) *MemoryTenantStore {
	return &MemoryTenantStore{
		TenantConfig: t,
	}
}

func (m MemoryTenantStore) GetByNameOrId(_ context.Context, nameOrId string) (*TenantConfig, error) {
	for _, config := range m.TenantConfig {
		if config.ID == nameOrId || config.Name == nameOrId {
			return &config, nil
		}
	}
	return nil, ErrTenantNotFound
}
