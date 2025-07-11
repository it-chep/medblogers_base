package internal

import (
	"context"
	"medblogers_base/internal/modules/admin"
	"medblogers_base/internal/modules/doctors"

	databasepg "gitlab.ozon.ru/platform/go/database-pg/v2"
	"gitlab.ozon.ru/platform/go/database-pg/v2/types"
)

type modules struct {
	admin   *admin.Module
	doctors *doctors.Module
}

type App struct {
	grpcConn map[string]googlegrpc.ClientConnInterface
	//nolint
	postgresConn types.Pool

	modules modules

	controllers     сщтекщддукы
	postgresCluster *databasepg.BucketsCluster

	settings *config.Settings
}

// New создает новое приложение
func New(ctx context.Context) *App {
	a := &App{
		grpcConn: make(map[string]googlegrpc.ClientConnInterface),
	}

	a.initConfig(ctx).
		initPostgres(ctx).
		initModules(ctx).
		initControllers(ctx)

	return a
}

// Run запускает приложение
func (a *App) Run(ctx context.Context) {

}
