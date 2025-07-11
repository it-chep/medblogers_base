package internal

import (
	"context"

	"github.com/it-chep/medblogers_base/internal/config"

	moduleadmin "github.com/it-chep/medblogers_base/internal/modules/admin"
	moduledoctors "github.com/it-chep/medblogers_base/internal/modules/doctors"

	databasepg "gitlab.ozon.ru/platform/go/database-pg/v2"
	"gitlab.ozon.ru/platform/go/database-pg/v2/roles"
	"gitlab.ozon.ru/platform/go/database-pg/v2/shard"
	"gitlab.ozon.ru/platform/tracer-go/logger"
)

func (a *App) initPostgres(ctx context.Context) *App {
	host, buckets := config.GetPostgresHost(ctx, a.scratch)

	cluster, err := databasepg.NewBucketsCluster(
		ctx,
		host,
		postgres.GetShardKeyToBucketFn(shard.Bucket(buckets)),
		databasepg.WithRoleMappers(roles.RoleMapperSyncFallback, roles.RoleMapperAsyncFallback),
	)

	if err != nil {
		logger.Fatalf(ctx, "[APP][POSTGRES] не удалось создать кластер базы данных: %s", err)
	}

	a.postgresCluster = cluster
	a.postgresConn = databasepg.WrapBucketsCluster(cluster)

	// todo gracefull
	//	a.postgresConn.Close()
	//

	return a
}

func (a *App) initConfig(context.Context) *App {
	a.settings = config.NewSettings()
	return a
}

func (a *App) initModules(_ context.Context) *App {
	a.modules = modules{
		admin:   moduleadmin.New(),
		doctors: moduledoctors.New(),
	}

	return a
}
