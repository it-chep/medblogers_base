package internal

import (
	"context"
	"fmt"
	doctorsV1 "medblogers_base/internal/app/api/doctors/v1"
	seoV1 "medblogers_base/internal/app/api/seo/v1"
	httpV1 "medblogers_base/internal/app/router/v1"
	pkgHttp "medblogers_base/internal/pkg/http"
	"net"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"medblogers_base/internal/config"
	descDoctorsV1 "medblogers_base/internal/pb/medblogers_base/api/doctors/v1"
	descSeoV1 "medblogers_base/internal/pb/medblogers_base/api/seo/v1"

	pkgConfig "medblogers_base/internal/pkg/config"
	"medblogers_base/internal/pkg/postgres"
	"net/http"

	"medblogers_base/internal/modules/admin"
	"medblogers_base/internal/modules/doctors"
)

type modules struct {
	admin   *admin.Module
	doctors *doctors.Module
}

type router struct {
	routerV1 *httpV1.Router
}

type controllers struct {
	doctorsController *doctorsV1.Implementation
	seoController     *seoV1.Implementation
}

type Server struct {
	grpcServer *grpc.Server
}

type App struct {
	mux      *runtime.ServeMux
	postgres postgres.PoolWrapper

	httpConns map[string]pkgHttp.Executor

	modules modules

	// http сервер
	controllers controllers
	router      router
	server      *Server

	// конфиги
	config        *config.Config
	mutableConfig pkgConfig.Config

	// периодические задачи
	//tasks
	//worker_pool
}

// New создает новое приложение
func New(ctx context.Context) *App {
	a := &App{}

	a.initConfig(ctx).
		initPostgres(ctx).
		initMutableConfig(ctx).
		initHttpConns(ctx).
		initModules(ctx).
		initRouters(ctx).
		initControllers(ctx).
		initServer(ctx)

	return a
}

// Run запускает приложение
func (a *App) Run(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("application recovered from panic")
		}
	}()
	listen, err := net.Listen("tcp", a.config.Server.GrpcAddress)
	if err != nil {
		fmt.Printf("[APP] Не удалось создать listen: %e", err)
		return
	}

	httpProxyOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = descDoctorsV1.RegisterDoctorServiceHandlerFromEndpoint(ctx, a.mux, a.config.Server.GrpcAddress, httpProxyOpts)
	if err != nil {
		fmt.Printf("[APP] Не удалось зарегистрироваь gprc хэндлер: %e", err)
		return
	}
	err = descSeoV1.RegisterSeoHandlerFromEndpoint(ctx, a.mux, a.config.Server.GrpcAddress, httpProxyOpts)
	if err != nil {
		fmt.Printf("[APP] Не удалось зарегистрироваь gprc хэндлер: %e", err)
		return
	}

	go func() {
		if err := a.server.grpcServer.Serve(listen); err != nil {
			fmt.Printf("[APP][GPRC] Не удалось запустить приложение: %v", err)
		}
	}()

	fmt.Printf("[APP] Запуск приложения, подключение HTTP: %s, GRPC: %s \n", a.config.Server.Address, a.config.Server.GrpcAddress)
	//a.workerPool.Run(ctx)

	if err := http.ListenAndServe(a.config.Server.Address, a.router.routerV1.Router); err != nil {
		fmt.Printf("[APP] Не удалось запустить приложение: %e", err)
		return
	}
}
