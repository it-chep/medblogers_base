package internal

import (
	"context"
	"fmt"
	adminV1 "medblogers_base/internal/app/api/admin/v1"
	authV1 "medblogers_base/internal/app/api/auth"
	doctorsV1 "medblogers_base/internal/app/api/doctors/v1"
	freelancersV1 "medblogers_base/internal/app/api/freelancers/v1"
	seoV1 "medblogers_base/internal/app/api/seo/v1"
	moduleAuth "medblogers_base/internal/modules/auth"

	"google.golang.org/grpc/credentials/insecure"

	"medblogers_base/internal/app/interceptor"
	httpV1 "medblogers_base/internal/app/router/v1"
	"medblogers_base/internal/config"
	moduleadmin "medblogers_base/internal/modules/admin"
	moduledoctors "medblogers_base/internal/modules/doctors"
	moduleFreelancers "medblogers_base/internal/modules/freelancers"

	descAdminV1 "medblogers_base/internal/pb/medblogers_base/api/admin/v1"
	descAuthV1 "medblogers_base/internal/pb/medblogers_base/api/auth/v1"
	descDoctorsV1 "medblogers_base/internal/pb/medblogers_base/api/doctors/v1"
	descFreelancersV1 "medblogers_base/internal/pb/medblogers_base/api/freelancers/v1"
	descSeoV1 "medblogers_base/internal/pb/medblogers_base/api/seo/v1"
	pkgConfig "medblogers_base/internal/pkg/config"
	pkgHttp "medblogers_base/internal/pkg/http"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
	"net/http"
	"time"

	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"

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

func (a *App) initMutableConfig(_ context.Context) *App {
	a.mutableConfig = pkgConfig.New(a.postgres)
	return a
}

func (a *App) initHttpConns(_ context.Context) *App {
	a.httpConns = map[string]pkgHttp.Executor{
		config.Subscribers: &http.Client{
			Timeout: time.Second * 3,
		},
		config.Salebot: &http.Client{
			Timeout: time.Second * 3,
		},
	}
	return a
}

func (a *App) initModules(_ context.Context) *App {
	a.modules = modules{
		admin:       moduleadmin.New(a.postgres),
		doctors:     moduledoctors.New(a.httpConns, a.config, a.postgres),
		freelancers: moduleFreelancers.New(a.httpConns, a.config, a.postgres),
		auth:        moduleAuth.New(a.postgres),
	}

	return a
}

func (a *App) initCache(_ context.Context) *App {
	// todo сделать кэш
	return a
}

func (a *App) initControllers(_ context.Context) *App {
	a.controllers.doctorsController = doctorsV1.NewDoctorsService(a.modules.doctors, a.mutableConfig)
	a.controllers.seoController = seoV1.NewSeoService(a.modules.doctors, a.modules.freelancers)
	a.controllers.freelancersController = freelancersV1.NewFreelancersService(a.modules.freelancers)
	a.controllers.authController = authV1.NewAuthService(a.modules.auth, a.config)
	a.controllers.adminController = adminV1.NewAdminService(a.modules.admin, a.modules.auth, a.config)

	return a
}

func (a *App) initRouters(_ context.Context) *App {
	a.mux = runtime.NewServeMux()
	a.router.routerV1 = httpV1.NewRouter(a.config, a.mutableConfig)
	a.router.routerV1.Router.Mount("/", a.mux)
	return a
}

func (a *App) initServer(_ context.Context) *App {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcrecovery.UnaryServerInterceptor(),
			interceptor.AuthInterceptor,
			interceptor.ConfigInterceptor(a.mutableConfig),
			interceptor.EmailInterceptor(a.config),
			interceptor.LoggerInterceptor(logger.New()),
			interceptor.RateLimitInterceptor,
			interceptor.ResponseTimeInterceptor,
		),
	)
	descDoctorsV1.RegisterDoctorServiceServer(grpcServer, a.controllers.doctorsController)
	descSeoV1.RegisterSeoServer(grpcServer, a.controllers.seoController)
	descFreelancersV1.RegisterFreelancerServiceServer(grpcServer, a.controllers.freelancersController)
	descAuthV1.RegisterAuthServiceServer(grpcServer, a.controllers.authController)
	descAdminV1.RegisterAdminServiceServer(grpcServer, a.controllers.adminController)
	reflection.Register(grpcServer)

	a.server = &Server{
		grpcServer: grpcServer,
	}

	return a
}

func (a *App) initGRPCServiceHandlers(ctx context.Context) *App {
	httpProxyOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := descDoctorsV1.RegisterDoctorServiceHandlerFromEndpoint(ctx, a.mux, a.config.Server.GrpcAddress, httpProxyOpts)
	if err != nil {
		panic(fmt.Sprintf("[APP] Не удалось зарегистрироваь gprc хэндлер: %e", err))
	}

	err = descSeoV1.RegisterSeoHandlerFromEndpoint(ctx, a.mux, a.config.Server.GrpcAddress, httpProxyOpts)
	if err != nil {
		panic(fmt.Sprintf("[APP] Не удалось зарегистрироваь gprc хэндлер: %e", err))
	}

	err = descFreelancersV1.RegisterFreelancerServiceHandlerFromEndpoint(ctx, a.mux, a.config.Server.GrpcAddress, httpProxyOpts)
	if err != nil {
		panic(fmt.Sprintf("[APP] Не удалось зарегистрироваь gprc хэндлер: %e", err))
	}

	err = descAuthV1.RegisterAuthServiceHandlerFromEndpoint(ctx, a.mux, a.config.Server.GrpcAddress, httpProxyOpts)
	if err != nil {
		panic(fmt.Sprintf("[APP] Не удалось зарегистрироваь gprc хэндлер: %e", err))
	}

	err = descAdminV1.RegisterAdminServiceHandlerFromEndpoint(ctx, a.mux, a.config.Server.GrpcAddress, httpProxyOpts)
	if err != nil {
		panic(fmt.Sprintf("[APP] Не удалось зарегистрироваь gprc хэндлер: %e", err))
	}

	return a
}
