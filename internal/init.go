package internal

import (
	"context"
	"fmt"
	adminV1 "medblogers_base/internal/app/api/admin/blog/v1"
	mmV1 "medblogers_base/internal/app/api/admin/mm/v1"
	authV1 "medblogers_base/internal/app/api/auth"
	blogsV1 "medblogers_base/internal/app/api/blogs/v1"
	doctorsV1 "medblogers_base/internal/app/api/doctors/v1"
	freelancersV1 "medblogers_base/internal/app/api/freelancers/v1"
	seoV1 "medblogers_base/internal/app/api/seo/v1"

	"github.com/go-chi/chi/v5"
	base_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/not-for-prod/clay/server"
	"github.com/not-for-prod/clay/transport"
	"google.golang.org/grpc/metadata"

	"medblogers_base/internal/app/middleware"
	moduleAuth "medblogers_base/internal/modules/auth"
	mmDesc "medblogers_base/internal/pb/medblogers_base/api/admin/mm/v1"
	adminDesc "medblogers_base/internal/pb/medblogers_base/api/admin/v1"
	authDesc "medblogers_base/internal/pb/medblogers_base/api/auth/v1"
	blogsDesc "medblogers_base/internal/pb/medblogers_base/api/blogs/v1"
	doctorsDesc "medblogers_base/internal/pb/medblogers_base/api/doctors/v1"
	freelancersDesc "medblogers_base/internal/pb/medblogers_base/api/freelancers/v1"
	seoDesc "medblogers_base/internal/pb/medblogers_base/api/seo/v1"
	"medblogers_base/internal/pkg/token"

	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/config"
	moduleadmin "medblogers_base/internal/modules/admin"
	moduleBlogs "medblogers_base/internal/modules/blogs"
	moduledoctors "medblogers_base/internal/modules/doctors"
	moduleFreelancers "medblogers_base/internal/modules/freelancers"
	moduleSeo "medblogers_base/internal/modules/seo"

	pkgConfig "medblogers_base/internal/pkg/config"
	pkgHttp "medblogers_base/internal/pkg/http"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
	"net/http"
	"time"

	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
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
		admin:       moduleadmin.New(a.httpConns, a.config, a.postgres),
		doctors:     moduledoctors.New(a.httpConns, a.config, a.postgres),
		freelancers: moduleFreelancers.New(a.httpConns, a.config, a.postgres),
		auth:        moduleAuth.New(a.postgres),
		blogs:       moduleBlogs.NewModule(a.postgres, a.config),
		seo:         moduleSeo.New(a.postgres),
	}

	return a
}

func (a *App) initCache(_ context.Context) *App {
	// todo сделать кэш
	return a
}

func (a *App) initControllers(_ context.Context) *App {
	a.controllers = []transport.ServiceDesc{
		doctorsDesc.NewDoctorServiceServiceDesc(doctorsV1.NewDoctorsService(a.modules.doctors, a.mutableConfig)),
		freelancersDesc.NewFreelancerServiceServiceDesc(freelancersV1.NewFreelancersService(a.modules.freelancers)),
		authDesc.NewAuthServiceServiceDesc(authV1.NewAuthService(a.modules.auth, a.config)),
		blogsDesc.NewBlogServiceServiceDesc(blogsV1.NewService(a.modules.blogs)),
		adminDesc.NewAdminServiceServiceDesc(adminV1.NewAdminService(a.modules.admin, a.modules.auth, a.config)),
		seoDesc.NewSeoServiceDesc(seoV1.NewSeoService(a.modules.doctors, a.modules.freelancers, a.modules.seo)),
		mmDesc.NewMMAdminServiceServiceDesc(mmV1.NewMMService(a.modules.admin, a.modules.auth)),
	}

	return a
}

func (a *App) initServer(_ context.Context) *App {
	a.muxChi = chi.NewMux()
	a.clayServer = server.NewServer(
		7002,
		server.WithHTTPMux(a.muxChi),
		server.WithHTTPPort(8080),
		server.WithGRPCOpts(
			grpc.ChainUnaryInterceptor(
				grpcrecovery.UnaryServerInterceptor(),
				interceptor.EmailInterceptor(a.config),
				interceptor.AuthInterceptor,
				interceptor.ConfigInterceptor(a.mutableConfig),
				interceptor.LoggerInterceptor(logger.New()),
				interceptor.RateLimitInterceptor,
				interceptor.ResponseTimeInterceptor,
			),
		),
		server.WithHTTPMiddlewares(
			base_middleware.Recoverer,
			middleware.CORSMiddleware(a.config),
			middleware.ConfigMiddleware(a.mutableConfig),
			middleware.EmailMiddleware(a.config),
			middleware.LoggerMiddleware(logger.New()),
			middleware.RateLimitMiddleware,
			middleware.ResponseTimeMiddleware,
		),
		server.WithRuntimeServeMuxOpts(
			runtime.WithOutgoingHeaderMatcher(
				func(s string) (string, bool) {
					if s == "set-cookie" {
						return "set-cookie", true
					}
					return runtime.MetadataHeaderPrefix + s, false
				},
			),
			runtime.WithMetadata(
				func(ctx context.Context, req *http.Request) metadata.MD {
					tokenCookie, _ := req.Cookie(token.RefreshCookie)
					if tokenCookie == nil {
						return metadata.Pairs()
					}
					return metadata.Pairs(token.RefreshCookie, tokenCookie.Value)
				},
			),
		),
	)

	return a
}
