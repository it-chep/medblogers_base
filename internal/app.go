package internal

import (
	"context"
	"fmt"
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/auth"
	"medblogers_base/internal/modules/blogs"
	"medblogers_base/internal/modules/freelancers"
	"medblogers_base/internal/modules/seo"
	pkgConfig "medblogers_base/internal/pkg/config"
	pkgHttp "medblogers_base/internal/pkg/http"
	"medblogers_base/internal/pkg/postgres"
	"medblogers_base/internal/pkg/worker_pool"
	"runtime/debug"
	"time"

	"medblogers_base/internal/modules/admin"
	"medblogers_base/internal/modules/doctors"

	"github.com/go-chi/chi/v5"
	"github.com/not-for-prod/clay/server"
	"github.com/not-for-prod/clay/transport"
)

type modules struct {
	auth        *auth.Module
	admin       *admin.Module
	doctors     *doctors.Module
	freelancers *freelancers.Module
	blogs       *blogs.Module
	seo         *seo.Module
}

type App struct {
	muxChi   *chi.Mux
	postgres postgres.PoolWrapper

	clayServer *server.Server
	httpConns  map[string]pkgHttp.Executor

	modules     modules
	controllers []transport.ServiceDesc

	// конфиги
	config        *config.Config
	mutableConfig pkgConfig.Config

	// периодические задачи
	//tasks
	workerPool worker_pool.WorkerPool
}

// New создает новое приложение
func New(ctx context.Context) *App {
	a := &App{}

	a.initConfig(ctx).
		initPostgres(ctx).
		initMutableConfig(ctx).
		initServer(ctx).
		initHttpConns(ctx).
		initModules(ctx).
		initWorkers(ctx).
		initControllers(ctx)

	return a
}

// Run запускает приложение
func (a *App) Run(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("PANIC RECOVERED: %v\n", r)
			debug.PrintStack()
		}
	}()

	go a.workerPool.Run(ctx)

	fmt.Printf("[APP][GPRC] Приложение запустилось HTTP - http://localhost:8080, GRPC - http://localhost:7002 , Время старта: %s", time.Now().Format(time.DateTime))
	if err := a.clayServer.Run(a.controllers...); err != nil {
		fmt.Printf("[APP][GPRC] Не удалось запустить приложение: %v", err)
	}
}
