package internal

import (
	"context"
	"fmt"
	v1 "github.com/it-chep/medblogers_base/internal/app/api/doctors/v1"
	"github.com/it-chep/medblogers_base/internal/config"
	"github.com/it-chep/medblogers_base/internal/pkg/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
	"time"

	moduleadmin "github.com/it-chep/medblogers_base/internal/modules/admin"
	moduledoctors "github.com/it-chep/medblogers_base/internal/modules/doctors"
)

func (a *App) initLogger(_ context.Context) *App {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	a.logger = zap.New(core)
	return a
}

func (a *App) initPostgres(ctx context.Context) *App {

	DSN := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?pool_max_conns=%d",
		a.config.Storage.User, a.config.Storage.Password, a.config.Storage.Host,
		a.config.Storage.Port, a.config.Storage.Database, a.config.Storage.MaxConnects,
	)

	poolConfig, err := pgxpool.ParseConfig(DSN)
	if err != nil {
		panic("[APP] Не удалось распарсить конфиг POSTGRES")
	}
	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		panic("[APP] Не удалось создать пул соединений POSTGRES")
	}

	a.postgres = postgres.NewPoolWrapper(pool)
	// todo gracefull
	//	a.postgresConn.Close()
	//

	return a
}

func (a *App) initConfig(_ context.Context) *App {
	a.config = config.NewConfig()
	return a
}

func (a *App) initModules(_ context.Context) *App {
	a.modules = modules{
		admin:   moduleadmin.New(),
		doctors: moduledoctors.New(a.postgres),
	}

	return a
}

func (a *App) initCache(_ context.Context) *App {
	// todo сделать кэш
	return a
}

func (a *App) initControllers(_ context.Context) *App {
	a.controllers.restController = v1.NewService(a.modules.doctors)
	return a
}

func (a *App) initServer(_ context.Context) *App {

	a.server = &http.Server{
		Addr:         a.config.Server.Address,
		Handler:      a.controllers.restController,
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 10 * time.Second,
	}

	return a
}
