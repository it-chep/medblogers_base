package internal

import (
	"context"
	v1 "github.com/it-chep/medblogers_base/internal/app/api/doctors/v1"
	"github.com/it-chep/medblogers_base/internal/config"
	"github.com/it-chep/medblogers_base/internal/pkg/postgres"
	"go.uber.org/zap"
	"net/http"

	"github.com/it-chep/medblogers_base/internal/modules/admin"
	"github.com/it-chep/medblogers_base/internal/modules/doctors"
)

type modules struct {
	admin   *admin.Module
	doctors *doctors.Module
}

type controllers struct {
	restController *v1.Service
}

type App struct {
	logger   *zap.Logger
	postgres postgres.PoolWrapper

	modules modules

	// http сервер
	controllers controllers
	server      *http.Server

	config *config.Config

	// периодические задачи
	//tasks
	//worker_pool
}

// New создает новое приложение
func New(ctx context.Context) *App {
	a := &App{}

	a.initConfig(ctx).
		initLogger(ctx).
		initPostgres(ctx).
		initModules(ctx).
		initControllers(ctx).
		initServer(ctx)

	return a
}

// Run запускает приложение
func (a *App) Run(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			a.logger.Error("application recovered from panic")
		}
	}()

	a.logger.Info("[APP] Запуск приложения")
	//a.workerPool.Run(ctx)

	if err := a.server.ListenAndServe(); err != nil {
		a.logger.Fatal("[APP] Не удалось запустить приложение: %s", zap.Error(err))
	}
}
