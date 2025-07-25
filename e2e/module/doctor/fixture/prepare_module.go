package fixture

import (
	"context"
	"medblogers_base/e2e/module/doctor/app"
	"medblogers_base/internal/pkg/postgres"
)

var SetupModule = func(ctx context.Context, pool *postgres.PoolWrapper) *app.TestableModule {
	return &app.TestableModule{
		PoolWrapper: pool,
	}
}
