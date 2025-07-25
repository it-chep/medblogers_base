package internal

import (
	"context"
	"fmt"
	doctorsV1 "medblogers_base/internal/app/api/doctors/v1"
	"medblogers_base/internal/app/interceptor"
	httpV1 "medblogers_base/internal/app/router/v1"
	"medblogers_base/internal/config"
	moduleadmin "medblogers_base/internal/modules/admin"
	moduledoctors "medblogers_base/internal/modules/doctors"
	desc "medblogers_base/internal/pb/medblogers_base/api/doctors/v1"
	pkgConfig "medblogers_base/internal/pkg/config"
	"medblogers_base/internal/pkg/postgres"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

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

func (a *App) initMutableConfig(ctx context.Context) *App {
	a.mutableConfig = pkgConfig.New(a.postgres)
	return a
}

func (a *App) initModules(_ context.Context) *App {
	a.modules = modules{
		admin:   moduleadmin.New(),
		doctors: moduledoctors.New(a.config, a.postgres),
	}

	return a
}

func (a *App) initCache(_ context.Context) *App {
	// todo сделать кэш
	return a
}

func (a *App) initControllers(_ context.Context) *App {
	a.controllers.doctorsController = doctorsV1.NewDoctorsService(a.modules.doctors, a.mutableConfig)
	return a
}

func (a *App) initRouters(_ context.Context) *App {
	a.mux = runtime.NewServeMux()
	a.router.routerV1 = httpV1.NewRouter(a.mutableConfig)
	a.router.routerV1.Router.Mount("/", a.mux)
	return a
}

func (a *App) initServer(_ context.Context) *App {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.AuthInterceptor,
			interceptor.ConfigInterceptor(a.mutableConfig),
			interceptor.LoggerInterceptor,
			interceptor.RateLimitInterceptor,
			interceptor.ResponseTimeInterceptor,
		),
	)
	desc.RegisterDoctorServiceServer(grpcServer, a.controllers.doctorsController)
	reflection.Register(grpcServer)

	a.server = &Server{
		grpcServer: grpcServer,
	}

	return a
}
