package internal

import (
	"context"

	"github.com/it-chep/medblogers_base/internal/config"

	moduleadmin "github.com/it-chep/medblogers_base/internal/modules/admin"
	moduledoctors "github.com/it-chep/medblogers_base/internal/modules/doctors"
)

func (a *App) initPostgres(ctx context.Context) *App {
	host, buckets := config.GetPostgresHost(ctx, a.scratch)

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
